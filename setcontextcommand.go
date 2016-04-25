package obeliskagi

import "fmt"

type setContextCommand struct {
	context string
}

func (cmd *setContextCommand) compile() string {
	return fmt.Sprintf("SET CONTEXT %s", cmd.context)
}

// SetContext is a command for setting the context.
func SetContext(context string) Command {
	return &setContextCommand{
		context: context,
	}
}
