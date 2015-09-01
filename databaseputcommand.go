package obeliskagi

import "fmt"

type databasePutCommand struct {
	family string
	key    string
	value  string
}

func (cmd *databasePutCommand) compile() string {
	return fmt.Sprintf("DATABASE PUT %s %s %s", escapeAndQuote(cmd.family), escapeAndQuote(cmd.key), escapeAndQuote(cmd.value))
}

// DatabasePut command will add or update an entry in the Asterisk database for a given family, key, and value.
func DatabasePut(family string, key string, value string) Command {
	return &databasePutCommand{
		family: family,
		key:    key,
		value:  value,
	}
}
