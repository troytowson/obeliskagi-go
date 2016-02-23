package obeliskagi

import (
	"bufio"
	"io"
	"net"
)

// Channel is the means of communication between the asterisk server and you.
type Channel struct {
	reader *bufio.Reader
	writer *bufio.Writer
	Ctx    *Context
}

func openChannel(conn net.Conn, scriptFunc ObeliskScriptFunc) {
	channel := &Channel{
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}

	channel.readContext()
	scriptFunc(*channel)
}

func (chnl *Channel) readContext() {
	var lines []string
	for {
		ln, _, err := chnl.reader.ReadLine()

		if err == io.EOF {
			break
		}

		lines = append(lines, string(ln))
	}

	chnl.Ctx = newContext(lines)
}

// SendCommand will send the command on the current channel
func (chnl *Channel) SendCommand(cmd Command) {
	/*chnl.writer.WriteString(cmd.compile())
	if _, err := chnl.reader.ReadString('\n'); err != nil {

	}*/
}
