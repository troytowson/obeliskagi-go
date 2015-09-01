package obeliskagi

import (
	"bufio"
	"net"
)

// Context is the details of the channel.
type Context struct {
	channelName string
}

// Channel is the means of communication between the asterisk server and you.
type Channel struct {
	reader  *bufio.Reader
	writer  *bufio.Writer
	conn    net.Conn
	context Context
}

func (chnl *Channel) readContext() {

}

// SendCommand will send the command on the current channel
func (chnl *Channel) SendCommand(cmd Command) {
	chnl.writer.WriteString(cmd.compile())
	_, err := chnl.reader.ReadString('\n')

	if err != nil {

	}

}
