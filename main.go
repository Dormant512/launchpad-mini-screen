package main

import (
	"github.com/MarinX/keylogger"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/portmididrv" // autoregisters driver
	"log"
)

func main() {
	fontMap, _ := parseFontFile("./letters")

	defer midi.CloseDriver()

	out, err := midi.FindOutPort("Launchpad Mini MIDI 1")
	if err != nil {
		log.Fatal("can't find Launchpad Mini")
		return
	}

	keyboard := keylogger.FindKeyboardDevice()
	if len(keyboard) == 0 {
		log.Fatal("can't find Keyboard")
		return
	}
	k, err := keylogger.New(keyboard)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer k.Close()

	send, _ := midi.SendTo(out)
	defer func() {
		for i := uint8(0); i < 255; i++ {
			send(midi.NoteOff(0, i))
		}
	}()

	events := k.Read()

	for e := range events {
		switch e.Type {
		case keylogger.EvKey:
			if e.KeyPress() {
				err = lightLetter(fontMap, send, e.KeyString())
				if err != nil {
					log.Fatal(err)
				}
			}
			if e.KeyRelease() {
				err = lightLetter(fontMap, send, "CLEAR")
				if err != nil {
					log.Fatal(err)
				}
			}
			break
		}
	}
}
