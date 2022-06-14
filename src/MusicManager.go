package main

import (
	"gfx"
	"time"
)

var musikID Identifier = CreateIdentifier("music/Inffertig3.wav")

const musikLänge = time.Minute * 4

func StarteMusik() {
	go func() {
		for {
			gfx.SpieleSound(musikID.getPath())
			time.Sleep(musikLänge)
		}
	}()
}
