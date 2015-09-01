package obeliskagi

import "fmt"

type getFullVariableCommand struct {
	variableName string
	channelName  string
}

func (cmd *getFullVariableCommand) compile() string {
	return fmt.Sprintf("GET FULL VARIABLE %s %s", escapeAndQuote(cmd.variableName), escapeAndQuote(cmd.channelName))
}

// GetFullVariable gets the full variable for the channel.
func GetFullVariable(variableName string, channelName string) Command {
	return &getFullVariableCommand{
		variableName: variableName,
		channelName:  channelName,
	}
}
