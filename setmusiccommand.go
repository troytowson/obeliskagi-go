package obeliskagi

import "fmt"

type setMusicCommand struct {
	musicStatus bool
}

func (cmd *setMusicCommand) compile() string {
	var status string
	if cmd.musicStatus {
		status = "ON"
	} else {
		status = "OFF"
	}

	return fmt.Sprintf("SET MUSIC %s", status)
}

const (
	// MusicOn indicates the music status is ON.
	MusicOn = true
	// MusicOff indicates the music status is OFF.
	MusicOff = !MusicOn
)

// SetMusic is a command for setting the extension.
func SetMusic(musicStatus bool) Command {
	return &setMusicCommand{
		musicStatus: musicStatus,
	}
}
