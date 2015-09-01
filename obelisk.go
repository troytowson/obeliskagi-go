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
	scriptFunc ObeliskFunc
}

// New intialises a new instance of the Obelisk FastAGI
func New(scriptFunc ObeliskFunc) *Obelisk {
	return &Obelisk{
		scriptFunc: scriptFunc,
	}
}

// ObeliskFunc is the function type for the api.
type ObeliskFunc func(channel Channel)

// Start is used to start listening to the Asterisk server.
func (agi *Obelisk) Start() error {
	if agi.isRunning() {
		return fmt.Errorf("Unable to start as it has already been started.")
	}

	fn := func() {
		agi.setRunning(true)

		l, err := net.Listen("tcp", "")

		if err != nil {
			return
		}

		defer l.Close()

		for agi.isRunning() {

			if conn, err := l.Accept(); err != nil {
				go openChannel(conn, agi.scriptFunc)
			}
		}
	}
	go fn()

	return nil
}

func openChannel(conn net.Conn, scriptFunc ObeliskFunc) {

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
