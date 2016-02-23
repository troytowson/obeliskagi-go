package obeliskagi

import "fmt"

type recordFileCommand struct {
	fileName      string
	format        string
	escapeDigits  string
	timeout       int
	offsetSamples int
	enableBeep    bool
	maxSilence    int
}

func (cmd *recordFileCommand) compile() string {
	enableBeepString := ""
	if cmd.enableBeep {
		enableBeepString = "BEEP"
	}

	return fmt.Sprintf("RECORD FILE %s %s %s %d %b %s S=%d",
		escapeAndQuote(cmd.fileName),
		escapeAndQuote(cmd.format),
		escapeAndQuote(cmd.escapeDigits),
		cmd.timeout,
		cmd.offsetSamples,
		enableBeepString,
		cmd.maxSilence)
}

// RecordFile is a ...
func RecordFile(fileName string, format string, escapeDigits string, timeout int, offsetSamples int, enableBeep bool, maxSilence int) Command {
	return &recordFileCommand{
		fileName:      fileName,
		format:        format,
		escapeDigits:  escapeDigits,
		timeout:       timeout,
		offsetSamples: offsetSamples,
		enableBeep:    enableBeep,
		maxSilence:    maxSilence,
	}
}
