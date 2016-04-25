package obeliskagi

import (
	"fmt"
	"time"
)

type sayDateTimeCommand struct {
	date         *time.Time
	escapeDigits string
	format       string
	timezone     string
}

func (cmd *sayDateTimeCommand) compile() string {
	return fmt.Sprintf("SAY DATETIME %d %s %s %s",
		cmd.date.Unix(),
		escapeAndQuote(cmd.escapeDigits),
		escapeAndQuote(cmd.format),
		escapeAndQuote(cmd.timezone))
}

// SayDateTime is a command or saying a date and time.
func SayDateTime(date *time.Time, escapeDigits, format, timezone string) Command {
	return &sayDateTimeCommand{
		date:         date,
		escapeDigits: escapeDigits,
		format:       format,
		timezone:     timezone,
	}
}
