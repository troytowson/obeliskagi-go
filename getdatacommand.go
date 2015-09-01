package obelisk

import "fmt"

type getDataCommand struct {
	file      string
	timeout   int
	maxDigits int
}

func (cmd *getDataCommand) compile() string {

	return fmt.Sprintf("GET DATA %s %d %d", escapeAndQuote(cmd.file), cmd.timeout, cmd.maxDigits)
}

// GetData command gets data.
func GetData(file string, timeout int, maxDigits int) Command {
	return &getDataCommand{
		file:      file,
		timeout:   timeout,
		maxDigits: maxDigits,
	}
}
