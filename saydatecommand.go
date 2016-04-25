package obeliskagi

import (
	"fmt"
	"time"
)

type sayDateCommand struct {
	date         *time.Time
	escapeDigits string
}

func (cmd *sayDateCommand) compile() string {
	return fmt.Sprintf("SAY DATE %d %s", cmd.date.Unix(), escapeAndQuote(cmd.escapeDigits))
}

// SayDate is a command for saying a date.
func SayDate(date *time.Time, escapeDigits string) Command {
	return &sayDateCommand{
		date:         date,
		escapeDigits: escapeDigits,
	}
}
