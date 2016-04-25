package obeliskagi

import "fmt"

type sayPhoneticCommand struct {
	string       string
	escapeDigits string
}

func (cmd *sayPhoneticCommand) compile() string {
	return fmt.Sprintf("SAY PHONETIC %s %s", escapeAndQuote(cmd.string), escapeAndQuote(cmd.escapeDigits))
}

// SayPhonetic is a command for saying a number.
func SayPhonetic(string, escapeDigits string) Command {
	return &sayPhoneticCommand{
		string:       string,
		escapeDigits: escapeDigits,
	}
}
