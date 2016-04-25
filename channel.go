package obeliskagi

import (
	"bufio"
	"io"
	"io/ioutil"
	"net"
	"strings"
)

// Channel is the means of communication between the asterisk server and you.
type Channel struct {
	reader  io.Reader
	writer  io.Writer
	Context *Context
}

func newChannel(conn net.Conn) *Channel {
	return &Channel{
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

func readData(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (chnl *Channel) fetchContext() error {
	d, err := readData(chnl.reader)
	if err != nil {
		return err
	}
	chnl.Context = newContext(strings.Split(d, "\n"))
	return nil
}

func (chnl *Channel) readReply() (*Reply, error) {
	d, err := readData(chnl.reader)
	if err != nil {
		return nil, err
	}
	return newReply(d), nil
}

// SendCommand will send the command on the current channel
func (chnl *Channel) SendCommand(cmd Command) (*Reply, error) {
	io.WriteString(chnl.writer, cmd.compile())
	return chnl.readReply()
}
