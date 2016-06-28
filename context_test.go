package obeliskagi

import (
	"bytes"
	"testing"
)

func TestMapLinesToFields(t *testing.T) {
	lines := []string{"agi_test:test"}
	actions := map[string]func(string){"agi_test": func(val string) {
		if val != "test" {
			t.Fatal("E")
		}
	}}

	mapLinesToFields(lines, actions)
}

func BenchmarkMapLinesToFields(t *testing.B) {

	lines := []string{"agi_test:test"}
	actions := map[string]func(string){"agi_test": func(val string) {}}

	for n := 0; n < t.N; n++ {
		mapLinesToFields(lines, actions)
	}
}

var (
	requestStub      = "request"
	channelStub      = "channel"
	languageStub     = "language"
	typeStub         = "type"
	uniqueIDStub     = "uniqueID"
	versionStub      = "version"
	callerIDStub     = "callerID"
	callerIDNameStub = "callerIDName"
	callingPresStub  = "callingPres"
	callingANI2Stub  = "callingANI2"
	callingTONStub   = "callingTON"
	callingTNSStub   = "callingTNS"
	dNIDStub         = "DNID"
	rDNISStub        = "RDNIS"
	contextStub      = "context"
	extensionStub    = "extension"
	priorityStub     = "priority"
	enhancedStub     = "enhanced"
	accountCodeStub  = "accountCode"
	threadIDStub     = "threadID"
	scriptStub       = "script"
)

func createBuffer() *bytes.Buffer {
	b := &bytes.Buffer{}
	b.WriteString(agiRequest + ":" + requestStub + "\n")
	b.WriteString(agiChannel + ":" + channelStub + "\n")
	b.WriteString(agiLanguage + ":" + languageStub + "\n")
	b.WriteString(agiType + ":" + typeStub + "\n")
	b.WriteString(agiUniqueID + ":" + uniqueIDStub + "\n")
	b.WriteString(agiVersion + ":" + versionStub + "\n")
	b.WriteString(agiCallerID + ":" + callerIDStub + "\n")
	b.WriteString(agiCallerIDName + ":" + callerIDNameStub + "\n")
	b.WriteString(agiCallingPres + ":" + callingPresStub + "\n")
	b.WriteString(agiCallingANI2 + ":" + callingANI2Stub + "\n")
	b.WriteString(agiCallingTON + ":" + callingTONStub + "\n")
	b.WriteString(agiCallingTNS + ":" + callingTNSStub + "\n")
	b.WriteString(agiDNID + ":" + dNIDStub + "\n")
	b.WriteString(agiRDNIS + ":" + rDNISStub + "\n")
	b.WriteString(agiContext + ":" + contextStub + "\n")
	b.WriteString(agiExtension + ":" + extensionStub + "\n")
	b.WriteString(agiPriority + ":" + priorityStub + "\n")
	b.WriteString(agiEnhanced + ":" + enhancedStub + "\n")
	b.WriteString(agiAccountCode + ":" + accountCodeStub + "\n")
	b.WriteString(agiThreadID + ":" + threadIDStub + "\n")
	b.WriteString(agiNetworkScript + ":" + scriptStub + "\n")
	b.WriteString("\n")
	return b
}

func TestNewContext(t *testing.T) {
	ctx, _ := fetchContext(createBuffer())

	if ctx == nil {
		t.Fail()
	}

	if ctx.RequestURL != requestStub ||
		ctx.ChannelName != channelStub ||
		ctx.Language != languageStub ||
		ctx.Type != typeStub ||
		ctx.UniqueID != uniqueIDStub ||
		ctx.Version != versionStub ||
		ctx.CallerID != callerIDStub ||
		ctx.CallerIDName != callerIDNameStub ||
		ctx.CallingPres != callingPresStub ||
		ctx.CallingANI2 != callingANI2Stub ||
		ctx.CallingTON != callingTONStub ||
		ctx.CallingTNS != callingTNSStub ||
		ctx.DNID != dNIDStub ||
		ctx.RDNIS != rDNISStub ||
		ctx.Context != contextStub ||
		ctx.Extension != extensionStub ||
		ctx.Priority != priorityStub ||
		ctx.Enhanced != enhancedStub ||
		ctx.AccountCode != accountCodeStub ||
		ctx.ThreadID != threadIDStub ||
		ctx.Script != scriptStub {
		t.Fail()
	}
}

func BenchmarkNewContext(b *testing.B) {
	buf := createBuffer()

	for i := 0; i < b.N; i++ {
		fetchContext(buf)
	}
}
