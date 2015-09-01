package obeliskagi

func main() {

	myFunc := func(chnl Channel) {
		chnl.SendCommand(Execute("RINGING"))
		chnl.SendCommand(Wait(1))
		chnl.SendCommand(Answer())
		chnl.SendCommand(Wait(1))
		chnl.SendCommand(Hangup(chnl.context.channelName))
	}

	obelisk := New(myFunc)
	obelisk.Start()
}
