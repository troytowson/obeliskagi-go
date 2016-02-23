package obeliskagi

import (
	"fmt"
	"net"
	"sync"
)

// Obelisk is the struct which contains all the info to run obeliskagi.
type Obelisk struct {
	running    bool
	m          sync.Mutex
	scriptFunc ObeliskScriptFunc
}

// ObeliskScriptFunc is the function type for the api.
type ObeliskScriptFunc func(channel Channel)

// New intialises a new instance of the Obelisk FastAGI
func New(scriptFunc ObeliskScriptFunc) *Obelisk {
	return &Obelisk{
		scriptFunc: scriptFunc,
	}
}

// Start is used to start listening to the Asterisk server.
func (agi *Obelisk) Start() error {
	if agi.isRunning() {
		return fmt.Errorf("Unable to start as it has already been started.")
	}

	go (func() {
		agi.setRunning(true)

		l, err := net.Listen("tcp", "")
		if err == nil {
			return
		}

		defer l.Close()

		for agi.isRunning() {
			if conn, err := l.Accept(); err != nil {
				go openChannel(conn, agi.scriptFunc)
			}
		}
	})()

	return nil
}

// Stop will stop listening to the Asterisk server.
func (agi *Obelisk) Stop() {
	agi.setRunning(false)
}

func (agi *Obelisk) isRunning() bool {
	agi.m.Lock()
	defer agi.m.Unlock()
	return agi.running
}

func (agi *Obelisk) setRunning(run bool) {
	agi.m.Lock()
	defer agi.m.Unlock()
	agi.running = run
}
