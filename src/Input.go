package main

import (
	"gfx"
)

//Ziel ist es die Maus- und Tastatureingaben des Spielers
//zu lesen und diese weiter zu verarbieten.
//Das ganze soll Asynchrone/Paralell zum Main Thread des Spiels geschehen,
//da der Prozess während des Lesens angehalten wird und es so zu rucklern
//kommen würde

var (
	MouseY uint16
	MouseX uint16
)

func StartInputListen(clickConsumer func(taste uint8, status int8), tastaturConsumer func(taste uint16, gedrueckt uint8, tiefe uint16)) {
	go listenMouse(clickConsumer)
	go listenKeyBoard(tastaturConsumer)
}

func listenMouse(clickConsumer func(taste uint8, status int8)) {
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		MouseX = mausX
		MouseY = mausY
		if taste == 1 || taste == 3 {
			clickConsumer(taste, status)
		}
	}
}

func listenKeyBoard(tastaturConsumer func(taste uint16, gedrueckt uint8, tiefe uint16)) {
	for {
		taste, gedrueckt, tiefe := gfx.TastaturLesen1()
		tastaturConsumer(taste, gedrueckt, tiefe)
	}
}

type InputConsumer interface {
	MausClick(taste uint8, status int8)
}
