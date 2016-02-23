package obeliskagi

import "fmt"

type getOptionCommand struct {
	fileName     string
	escapeDigits string
	timeout      int
}

func (cmd *getOptionCommand) compile() string {
	return fmt.Sprintf("GET OPTION %s %s %d", escapeAndQuote(cmd.fileName), escapeAndQuote(cmd.escapeDigits), cmd.timeout)
}

// GetOption will get the options.
func GetOption(fileName string, escapeDigits string, timeout int) Command {
	return &getOptionCommand{
		fileName:     fileName,
		escapeDigits: escapeDigits,
		timeout:      timeout,
	}
}
