package obeliskagi

import (
	"bytes"
	"strings"
)

const (
	agiRquest          = "agi_request"
	agiChannel         = "agi_channel"
	agiLanguage        = "agi_language"
	agiType            = "agi_type"
	agiUniqueID        = "agi_uniqueid"
	agiVersion         = "agi_version"
	agiCallerID        = "agi_callerid"
	agiCallerIDName    = "agi_calleridname"
	agiCallingPres     = "agi_callingpres"
	agiCallingANI2     = "agi_callingani2"
	agiCallingTON      = "agi_callington"
	agiCallingTNS      = "agi_callingtns"
	agiDNID            = "agi_dnid"
	agiRDNIS           = "agi_rdnis"
	agiContext         = "agi_context"
	agiExtension       = "agi_extension"
	agiPriority        = "agi_priority"
	agiEnhanced        = "agi_enhanced"
	agiAccountCode     = "agi_accountcode"
	agiThreadID        = "agi_threadid"
	agiCustomARGPrefix = "agi_arg_"
	agiNetworkScript   = "agi_network_script"
)

// Context is the details of the channel.
type Context struct {
	RequestURL   string
	ChannelName  string
	Language     string
	Type         string
	UniqueID     string
	Version      string
	CallerID     string
	CallerIDName string
	CallingPres  string
	CallingANI2  string
	CallingTON   string
	CallingTNS   string
	DNID         string
	RDNIS        string
	Context      string
	Extension    string
	Priority     string
	Enhanced     string
	AccountCode  string
	ThreadID     string
	Script       string
}

func newContext(lines []string) *Context {
	ctx := &Context{}

	mapLinesToFields(lines, map[string]func(string){
		agiRquest:        func(v string) { ctx.RequestURL = v },
		agiChannel:       func(v string) { ctx.ChannelName = v },
		agiLanguage:      func(v string) { ctx.Language = v },
		agiType:          func(v string) { ctx.Type = v },
		agiUniqueID:      func(v string) { ctx.UniqueID = v },
		agiVersion:       func(v string) { ctx.Version = v },
		agiCallerID:      func(v string) { ctx.CallerID = v },
		agiCallerIDName:  func(v string) { ctx.CallerIDName = v },
		agiCallingPres:   func(v string) { ctx.CallingPres = v },
		agiCallingANI2:   func(v string) { ctx.CallingANI2 = v },
		agiCallingTON:    func(v string) { ctx.CallingTON = v },
		agiCallingTNS:    func(v string) { ctx.CallingTNS = v },
		agiDNID:          func(v string) { ctx.DNID = v },
		agiRDNIS:         func(v string) { ctx.RDNIS = v },
		agiContext:       func(v string) { ctx.Context = v },
		agiExtension:     func(v string) { ctx.Extension = v },
		agiPriority:      func(v string) { ctx.Priority = v },
		agiEnhanced:      func(v string) { ctx.Enhanced = v },
		agiAccountCode:   func(v string) { ctx.AccountCode = v },
		agiThreadID:      func(v string) { ctx.ThreadID = v },
		agiNetworkScript: func(v string) { ctx.Script = v },
	})

	return ctx
}

func mapLinesToFields(lines []string, actions map[string]func(string)) {
	for _, ln := range lines {
		var bf bytes.Buffer
		var k string

		for _, c := range ln {
			if c == ':' && k == "" {
				k = bf.String()
				bf.Reset()
				continue
			}
			bf.WriteRune(c)
		}

		v := bf.String()

		if a, ok := actions[k]; ok && strings.HasPrefix(k, "agi_") {
			a(strings.TrimSpace(v))
		}
	}
}
