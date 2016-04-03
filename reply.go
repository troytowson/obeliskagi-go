package obeliskagi

import (
	"regexp"
	"strconv"
)

const (
	// Trying represents a trying result.
	Trying = 100
	// Success represents a success result.
	Success = 200
	// InvalidOrUnknownCommand represents a invalid or unknown command result.
	InvalidOrUnknownCommand = 510
	// DeadChannel represents a dead channel result.
	DeadChannel = 511
	// InvalidCommandSyntax represents an invalid commmand syntax result
	InvalidCommandSyntax = 520
)

var replyPattern *regexp.Regexp

func init() {
	replyPattern, _ = regexp.Compile(`^(\d{3})\sresult=([^\s$]*)(?:\s\(([^)]*)\)){0,1}(?:\sendpos=(\d*)){0,1}$`)
}

// Reply is the response from asterisk once a command is sent.
type Reply struct {
	Line   string
	Result int
	Status int
	Data   string
}

// initialises a new instance of the Reply
func newReply(line string) *Reply {
	r := &Reply{
		Line: line,
	}

	res := replyPattern.FindStringSubmatch(line)

	if len(res) > 0 {
		r.Status, _ = strconv.Atoi(res[1])

		if len(res) > 1 {
			r.Result, _ = strconv.Atoi(res[2])

			if len(res) > 2 {
				r.Data = res[3]
			}
		}
	}

	return r
}
