package obeliskagi

import (
	"net"
	"sync"
)

// Obelisk is the struct which contains all the info to run obeliskagi.
type Obelisk struct {
	running bool
	m       sync.Mutex
	config  Configuration
}

// New intialises a new instance of the Obelisk FastAGI
func New(cfg Configuration) *Obelisk {
	return &Obelisk{
		config: cfg,
	}
}

// Start is used to start listening to the Asterisk server.
func (agi *Obelisk) Start() error {
	l, err := net.Listen("tcp", agi.config.Address)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go func(conn net.Conn, scriptFunc ObeliskScriptFunc) {
			defer conn.Close()
			channel := newChannel(conn)
			channel.fetchContext()
			scriptFunc(*channel)
		}(conn, agi.config.ScriptFunc)
	}
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
