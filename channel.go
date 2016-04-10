package obeliskagi

import (
	"bufio"
	"io"
	"net"
)

// Channel is the means of communication between the asterisk server and you.
type Channel struct {
	reader  *bufio.Reader
	writer  *bufio.Writer
	Context *Context
}

func newChannel(conn net.Conn) *Channel {
	return &Channel{
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
}

func openChannel(conn net.Conn, logger StandardLogger, scriptFunc ObeliskScriptFunc) {
	defer conn.Close()

	channel := newChannel(conn)
	channel.fetchContext()

	scriptFunc(*channel)

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

func (chnl *Channel) fetchContext() {
	lines := readLines(chnl.reader)
	chnl.Context = newContext(lines)
}

func (chnl *Channel) readReply() *Reply {
	var ln string
	lns := readLines(chnl.reader)
	for _, l := range lns {
		ln += l
	}
	return newReply(ln)
}

// SendCommand will send the command on the current channel
func (chnl *Channel) SendCommand(cmd Command) *Reply {
	chnl.writer.WriteString(cmd.compile())
	return chnl.readReply()
}
