package obeliskagi

// Reply is the response from asterisk once a command is sent.
type Reply struct {
}

// initialises a new instance of the Reply
func newReply(reply string) *Reply {
	return &Reply{}
}

const (
	// Trying
	Trying = 100
	// Success
	Success = 200
	// InvalidOrUnknownCommand
	InvalidOrUnknownCommand = 510
	// DeadChannel
	DeadChannel = 511
	// InvalidCommandSyntax
	InvalidCommandSyntax = 520
)
