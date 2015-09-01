package obeliskagi

type answerCommand struct {
}

func (cmd *answerCommand) compile() string {
	return "ANSWER"
}

// Answer command will answer the call.
func Answer() Command {
	return &answerCommand{}
}
