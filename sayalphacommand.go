package obeliskagi

import "fmt"

type sayAlphaCommand struct {
	number       string
	escapeDigits string
}

func (cmd *sayAlphaCommand) compile() string {
	return fmt.Sprintf("SAY ALPHA %s %s", escapeAndQuote(cmd.number), escapeAndQuote(cmd.escapeDigits))
}

// SayAlpha is a command for saying a given string
func SayAlpha(number, escapeDigits string) Command {
	return &sayAlphaCommand{
		number:       number,
		escapeDigits: escapeDigits,
	}
}
