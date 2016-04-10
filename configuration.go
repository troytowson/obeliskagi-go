package obeliskagi

// ObeliskScriptFunc is the function type for the api.
type ObeliskScriptFunc func(channel Channel)

// Configuration represents the configuration of obelisk.
type Configuration struct {
	Address        string
	ScriptFunc     ObeliskScriptFunc
	LoggingEnabled bool
	Logger         StandardLogger
}
