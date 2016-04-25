package obeliskagi

import "fmt"

type sayDigitsCommand struct {
	number       string
	escapeDigits string
}

func (cmd *sayDigitsCommand) compile() string {
	return fmt.Sprintf("SAY DIGITS %s %s", escapeAndQuote(cmd.number), escapeAndQuote(cmd.escapeDigits))
}

// SayDigits is a command for saying digits.
func SayDigits(number, escapeDigits string) Command {
	return &sayDigitsCommand{
		number:       number,
		escapeDigits: escapeDigits,
	}
}
