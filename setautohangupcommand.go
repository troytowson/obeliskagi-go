package obeliskagi

import "fmt"

type setAutoHangupCommand struct {
	time int
}

func (cmd *setAutoHangupCommand) compile() string {
	return fmt.Sprintf("SET AUTOHANGUP %d", cmd.time)
}

// SetAutoHangup is a command for sending text.
func SetAutoHangup(time int) Command {
	return &setAutoHangupCommand{
		time: time,
	}
}
