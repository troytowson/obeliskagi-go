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
	lines := readLines(chnl.reader)
	chnl.Ctx = newContext(lines)
}

// SendCommand will send the command on the current channel
func (chnl *Channel) SendCommand(cmd Command) *Reply {
	chnl.writer.WriteString(cmd.compile())
	var ln string
	lns := readLines(chnl.reader)
	for _, l := range lns {
		ln += l
	}
	return newReply(ln)
}

func readLines(r *bufio.Reader) []string {
	var lines []string
	for {
		ln, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		lines = append(lines, ln)
	}
	return lines
}
