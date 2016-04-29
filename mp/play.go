package mp

import (
	"fmt"
)

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3", "mp3":
		p = &MP3Player{}
	case "WAV", "wav":
		p = &WAVPlayer{}
	default:
		fmt.Println("UNsupported music type --> ", mtype)
		return
	}

	p.Play(source)
}
