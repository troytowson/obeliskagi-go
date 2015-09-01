package obeliskagi

import (
	"bytes"
	"fmt"
	"strings"
)

// Command is a contract for all commands.
type Command interface {
	compile() string
}

func escapeAndQuote(option string) string {

	if strings.TrimSpace(option) == "" {
		option = ""
	}

	option = strings.Replace(option, "\\\\", "\\\\\\\\", 0)
	option = strings.Replace(option, "\\\"", "\\\\\"", 0)
	option = strings.Replace(option, "\\\n", "", 0)
	return fmt.Sprintf("\"%s\"", option)
}

func escapeAndQuoteArray(options []string) string {

	if options == nil {
		return escapeAndQuote("")
	}
	buffer := bytes.Buffer{}
	for _, option := range options {
		if buffer.Len() > 0 {
			buffer.WriteString(",")
		}

		buffer.WriteString(option)
	}

	return escapeAndQuote(buffer.String())
}
