package obeliskagi

import "fmt"

type databaseDeleteCommand struct {
	family string
	key    string
}

func (cmd *databaseDeleteCommand) compile() string {
	return fmt.Sprintf("DATABASE DEL %s %s", escapeAndQuote(cmd.family), escapeAndQuote(cmd.key))
}

// DatabaseDelete command will delete the database.
func DatabaseDelete(family string, key string) Command {
	return &databaseDeleteCommand{
		family: family,
		key:    key,
	}
}
