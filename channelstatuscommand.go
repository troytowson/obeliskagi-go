package obeliskagi

import "fmt"

type channelStatusCommand struct {
	channelName string
}

func (cmd *channelStatusCommand) compile() string {
	return fmt.Sprintf("CHANNEL STATUS %s", escapeAndQuote(cmd.channelName))
}

// ChannelStatus checks the status of the channel.
func ChannelStatus(cn string) Command {
	return &channelStatusCommand{
		channelName: cn,
	}
}
