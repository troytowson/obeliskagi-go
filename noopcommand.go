package obeliskagi

import "fmt"

type noopCommand struct {
}

func (cmd *noopCommand) compile() string {
	return fmt.Sprintf("NOOP")
}

// Noop is a no operation
func Noop() Command {
	return &noopCommand{}
}
