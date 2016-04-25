package obeliskagi

import "fmt"

type sendImageCommand struct {
	image string
}

func (cmd *sendImageCommand) compile() string {
	return fmt.Sprintf("SEND IAMGE %s", escapeAndQuote(cmd.image))
}

// SendImage is a command for sending images.
func SendImage(image string) Command {
	return &sendImageCommand{
		image: image,
	}
}
