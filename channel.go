package obeliskagi

import (
	"bufio"
	"io"
)

// Channel is the means of communication between the asterisk server and you.
type Channel struct {
	reader  io.Reader
	writer  io.Writer
	Context *Context
}

// SendCommand will send the command on the current channel
func (channel *Channel) SendCommand(cmd Command) (*Reply, error) {
	writer := bufio.NewWriter(channel.writer)
	writer.WriteString(cmd.compile() + "\n")
	writer.Flush()

	scanner := bufio.NewScanner(channel.reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return newReply(scanner.Text()), nil
}
