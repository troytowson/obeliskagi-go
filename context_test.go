package obeliskagi

import (
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

func TestNewContext(t *testing.T) {
	requestStub := "request"
	channelStub := "channel"
	languageStub := "language"
	typeStub := "type"
	uniqueIDStub := "uniqueID"
	versionStub := "version"
	callerIDStub := "callerID"
	callerIDNameStub := "callerIDName"
	callingPresStub := "callingPres"
	callingANI2Stub := "callingANI2"
	callingTONStub := "callingTON"
	callingTNSStub := "callingTNS"
	dNIDStub := "DNID"
	rDNISStub := "RDNIS"
	contextStub := "context"
	extensionStub := "extension"
	priorityStub := "priority"
	enhancedStub := "enhanced"
	accountCodeStub := "accountCode"
	threadIDStub := "threadID"
	scriptStub := "script"

	lines := []string{
		agiRquest + ":" + requestStub,
		agiChannel + ":" + channelStub,
		agiLanguage + ":" + languageStub,
		agiType + ":" + typeStub,
		agiUniqueID + ":" + uniqueIDStub,
		agiVersion + ":" + versionStub,
		agiCallerID + ":" + callerIDStub,
		agiCallerIDName + ":" + callerIDNameStub,
		agiCallingPres + ":" + callingPresStub,
		agiCallingANI2 + ":" + callingANI2Stub,
		agiCallingTON + ":" + callingTONStub,
		agiCallingTNS + ":" + callingTNSStub,
		agiDNID + ":" + dNIDStub,
		agiRDNIS + ":" + rDNISStub,
		agiContext + ":" + contextStub,
		agiExtension + ":" + extensionStub,
		agiPriority + ":" + priorityStub,
		agiEnhanced + ":" + enhancedStub,
		agiAccountCode + ":" + accountCodeStub,
		agiThreadID + ":" + threadIDStub,
		agiNetworkScript + ":" + scriptStub,
	}

	ctx := newContext(lines)

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
	requestStub := "request"
	channelStub := "channel"
	languageStub := "language"
	typeStub := "type"
	uniqueIDStub := "uniqueID"
	versionStub := "version"
	callerIDStub := "callerID"
	callerIDNameStub := "callerIDName"
	callingPresStub := "callingPres"
	callingANI2Stub := "callingANI2"
	callingTONStub := "callingTON"
	callingTNSStub := "callingTNS"
	dNIDStub := "DNID"
	rDNISStub := "RDNIS"
	contextStub := "context"
	extensionStub := "extension"
	priorityStub := "priority"
	enhancedStub := "enhanced"
	accountCodeStub := "accountCode"
	threadIDStub := "threadID"
	scriptStub := "script"

	lines := []string{
		agiRquest + ":" + requestStub,
		agiChannel + ":" + channelStub,
		agiLanguage + ":" + languageStub,
		agiType + ":" + typeStub,
		agiUniqueID + ":" + uniqueIDStub,
		agiVersion + ":" + versionStub,
		agiCallerID + ":" + callerIDStub,
		agiCallerIDName + ":" + callerIDNameStub,
		agiCallingPres + ":" + callingPresStub,
		agiCallingANI2 + ":" + callingANI2Stub,
		agiCallingTON + ":" + callingTONStub,
		agiCallingTNS + ":" + callingTNSStub,
		agiDNID + ":" + dNIDStub,
		agiRDNIS + ":" + rDNISStub,
		agiContext + ":" + contextStub,
		agiExtension + ":" + extensionStub,
		agiPriority + ":" + priorityStub,
		agiEnhanced + ":" + enhancedStub,
		agiAccountCode + ":" + accountCodeStub,
		agiThreadID + ":" + threadIDStub,
		agiNetworkScript + ":" + scriptStub,
	}

	for i := 0; i < b.N; i++ {
		newContext(lines)
	}
}
