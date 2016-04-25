package obeliskagi

import "fmt"

type setCallerIDCommand struct {
	callerID string
}

func (cmd *setCallerIDCommand) compile() string {
	return fmt.Sprintf("SET CALLERID %s", cmd.callerID)
}

// SetCallerID is a command for setting the caller id.
func SetCallerID(callerID string) Command {
	return &setCallerIDCommand{
		callerID: callerID,
	}
}
