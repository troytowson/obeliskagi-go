package obeliskagi

import "fmt"

type hangupCommand struct {
	channelName string
}

func (cmd *hangupCommand) compile() string {
	return fmt.Sprintf("HANGUP %s", escapeAndQuote(cmd.channelName))
}

// Hangup command will hangup the call.
func Hangup(channelName string) Command {
	return &hangupCommand{
		channelName: channelName,
	}
}
