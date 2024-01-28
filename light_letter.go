package main

import "gitlab.com/gomidi/midi/v2"

func lightLetter(fontMap map[string]grid, send func(msg midi.Message) error, key string) error {
	letterGrid, ok := fontMap[key]
	if !ok {
		return nil
	}
	for x, row := range letterGrid {
		for y, cell := range row {
			note := uint8(x*16 + y)
			if cell == '.' {
				err := send(midi.NoteOff(0, note))
				if err != nil {
					return err
				}
			} else {
				err := send(midi.NoteOn(0, note, R))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
