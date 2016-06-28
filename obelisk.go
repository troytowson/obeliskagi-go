package obeliskagi

import (
	"errors"
	"log"
	"net"
	"os"
)

// a logger for logging errors in the go routines
var errorLogger *log.Logger

// initialisation code.
func init() {
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Obelisk is the struct which contains all the info to run obeliskagi.
type Obelisk struct {
	config Configuration
}

// New intialises a new instance of the Obelisk FastAGI
func New(cfg Configuration) *Obelisk {
	return &Obelisk{
		config: cfg,
	}
}

// Listen is used to start listening to the Asterisk server.
func (agi *Obelisk) Listen() error {
	// Assert the address is not empty
	if agi.config.Address == "" {
		return errors.New("The address needs to be set.")
	}

	// Assert that there is at least one script added to the configuration.
	if len(agi.config.Scripts) > 0 {
		return errors.New("At least one script should be configured.")
	}

	// Start listening to the tcp socket.
	listener, err := net.Listen("tcp", agi.config.Address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			return err
		}

		go func(connection net.Conn, scripts map[string]ScriptFunction) {
			defer connection.Close()

			ctx, err := fetchContext(connection)
			if err != nil {
				errorLogger.Println(err)
				return
			}

			channel := &Channel{
				reader:  connection,
				writer:  connection,
				Context: ctx,
			}

			scriptFunction, found := scripts[channel.Context.Script]
			if !found {
				errorLogger.Println("The script could not be found.")
				return
			}
			scriptFunction(*channel)

		}(connection, agi.config.Scripts)
	}
}
