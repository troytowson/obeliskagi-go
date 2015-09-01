package obeliskagi

import "fmt"

type databaseDeleteTreeCommand struct {
	family  string
	keytree string
}

func (cmd *databaseDeleteTreeCommand) compile() string {
	return fmt.Sprintf("DATABASE DELTREE %s %s", escapeAndQuote(cmd.family), escapeAndQuote(cmd.keytree))
}

// DatabaseDeleteTree command will delte the database tree.
func DatabaseDeleteTree(family string, keytree string) Command {
	return &databaseDeleteTreeCommand{
		family:  family,
		keytree: keytree,
	}
}
