package obeliskagi

import "fmt"

type receiveTextCommand struct {
	timeout int
}

func (cmd *receiveTextCommand) compile() string {
	return fmt.Sprintf("RECEIVE TEXT %s", cmd.timeout)
}

// ReceiveText will recieve text.
func ReceiveText(timeout int) Command {
	return &receiveTextCommand{
		timeout: timeout,
	}
}
