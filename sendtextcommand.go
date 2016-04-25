package obeliskagi

import "fmt"

type sendTextCommand struct {
	text string
}

func (cmd *sendTextCommand) compile() string {
	return fmt.Sprintf("SEND TEXT %s", escapeAndQuote(cmd.text))
}

// SendText is a command for sending text.
func SendText(text string) Command {
	return &sendTextCommand{
		text: text,
	}
}
