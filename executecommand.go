package obeliskagi

import "fmt"

type executeCommand struct {
	application string
	options     []string
}

func (cmd *executeCommand) compile() string {
	return fmt.Sprintf("EXEC %s %s", cmd.application, escapeAndQuoteArray(cmd.options))
}

// Execute command will execute an action.
func Execute(application string, options ...string) Command {
	return &executeCommand{
		application: application,
		options:     options,
	}
}
