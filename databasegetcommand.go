package obeliskagi

import "fmt"

type databaseGetCommand struct {
	family string
	key    string
}

// DatabaseGet command will get value from the database with the specific key and family.
func DatabaseGet(family string, key string) Command {
	return &databaseGetCommand{
		family: family,
		key:    key,
	}
}

func (cmd *databaseGetCommand) compile() string {
	return fmt.Sprintf("DATABASE GET %s %s", escapeAndQuote(cmd.family), escapeAndQuote(cmd.key))
}
