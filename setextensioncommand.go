package obeliskagi

import "fmt"

type setExtensionCommand struct {
	extension string
}

func (cmd *setExtensionCommand) compile() string {
	return fmt.Sprintf("SET EXTENSION %s", cmd.extension)
}

// SetExtension is a command for setting the extension.
func SetExtension(extension string) Command {
	return &setExtensionCommand{
		extension: extension,
	}
}
