package obeliskagi

import "fmt"

type receiveCharCommand struct {
	timeout int
}

func (cmd *receiveCharCommand) compile() string {
	return fmt.Sprintf("RECEIVE CHAR %d", cmd.timeout)
}

// ReceiveChar will recieve a character
func ReceiveChar(timeout int) Command {
	return &receiveCharCommand{
		timeout: timeout,
	}
}
