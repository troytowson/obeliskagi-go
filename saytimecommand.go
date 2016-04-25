package obeliskagi

import (
	"fmt"
	"time"
)

type sayTimeCommand struct {
	date         *time.Time
	escapeDigits string
}

func (cmd *sayTimeCommand) compile() string {
	return fmt.Sprintf("SAY TIME %d %s", cmd.date.Unix(), escapeAndQuote(cmd.escapeDigits))
}

// SayTime is a command or saying a time.
func SayTime(date *time.Time, escapeDigits string) Command {
	return &sayDateTimeCommand{
		date:         date,
		escapeDigits: escapeDigits,
	}
}
