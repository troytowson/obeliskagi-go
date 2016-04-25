package obeliskagi

import "fmt"

type sayNumberCommand struct {
	number       string
	escapeDigits string
}

func (cmd *sayNumberCommand) compile() string {
	return fmt.Sprintf("SAY DIGITS %s %s", escapeAndQuote(cmd.number), escapeAndQuote(cmd.escapeDigits))
}

// SayNumber is a command for saying a number.
func SayNumber(number, escapeDigits string) Command {
	return &sayNumberCommand{
		number:       number,
		escapeDigits: escapeDigits,
	}
}
