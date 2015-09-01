package obeliskagi

import (
	"fmt"
	"strconv"
)

type waitCommand struct {
	duration int
}

func (cmd *waitCommand) compile() string {
	return fmt.Sprintf("EXEC WAIT %s", escapeAndQuote(strconv.Itoa(cmd.duration)))
}

// Wait is a command that will execute a wait in the server
func Wait(duration int) Command {
	return &waitCommand{
		duration: duration,
	}
}
