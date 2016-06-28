package obeliskagi

// ScriptFunction is the function type for the api.
type ScriptFunction func(channel Channel)

// Configuration represents the configuration of obelisk.
type Configuration struct {
	Address string
	Scripts map[string]ScriptFunction
}
