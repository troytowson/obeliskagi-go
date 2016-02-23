package obeliskagi

import "fmt"

type getVariableCommand struct {
	variableName string
}

func (cmd *getVariableCommand) compile() string {
	return fmt.Sprintf("GET VARIABLE %s", escapeAndQuote(cmd.variableName))
}

// GetVariable will get the variable
func GetVariable(variableName string) Command {
	return &getVariableCommand{
		variableName: variableName,
	}
}
