package main

import (
	"gfx"
	"strconv"
	"time"
)

const (
	targetFps   = 60
	optimalTime = 1000000000 / targetFps
	debugMode   = true
	width       = 1280
	height      = 720
)

var currentFps = 0
var lastFpsTime int64 = 0
var fpsTimer = 0
var currentDeltaTime int64

var currentScreen Screen
var lastIterationTime time.Time

var (
	lastMouseX uint16
	lastMouseY uint16
)

func initWindow() {
	//Vor.: -
	//Eff.: Initialisiert alles rund um gfx

	//Erstelle fenster
	gfx.Fenster(width, height)

	//Setzte Font
	InitFont()

	screen := MainMenuScreen{}
	SetScreen(&screen)
	lastIterationTime = time.Now()

	StartInputListen(MouseClickConsumer, TastaturConsumer)

}

func loop() {
	//Vor.: -
	//Eff.: Haupt GameLoop

	//berechne delta, um bei egal wie vielen fps das Spiel so zu updaten, als wären es 60 Frames per second (targetFps)
	now := time.Now()
	updateLength := now.Sub(lastIterationTime)
	lastIterationTime = now
	delta := int64(float64(updateLength.Nanoseconds()) / float64(optimalTime))
	currentDeltaTime = delta

	//fps time hochzählen, um diesen Wert sp#ter abfragen zu können
	lastFpsTime += updateLength.Nanoseconds()
	//fps timer hochzählen
	fpsTimer++

	//wenn das letzte fps update eine Sekunde oder länger her ist,
	//wird currentFps auf fpsTimer gesetzt und der fpsTimer wieder auf 0 gesetzt.
	if lastFpsTime > 1000000000 {
		lastFpsTime = 0
		currentFps = fpsTimer
		fpsTimer = 0
	}

	//Update aus um dual buffering zu ermöglichen
	//Erklärung: Alle dinge die gerendert werden, sollen nicht direkt nach dem rendern angezeigt werde,
	//sondern erst, wenn der Bildschirm fertig gerendert wurde.
	//Also ein Bildschirm, den man sieht und im Hintergrund wird der nächste gerendert
	//und anschließend angezeigt.
	gfx.UpdateAus()

	//Stiftfarbe auf Schwarz setzten, damit der geleerte bildschirm schwarz ist
	gfx.Stiftfarbe(0, 0, 0)
	//Alten Bildschirm leeren/clearen
	gfx.Cls()

	//Stiftfarbe auf standart setzten
	gfx.Stiftfarbe(255, 255, 255)

	//Wenn sich die position zum letzten update verändert hat,
	//wird im screen die methode mouseMove ausgeführt
	if MouseX != lastMouseX || MouseY != lastMouseY {
		currentScreen.mouseMove(MouseX, MouseY, int(MouseX)-int(lastMouseX), int(MouseY)-int(lastMouseY))
	}
	//Setzte aktuelle maus position
	lastMouseX = MouseX
	lastMouseY = MouseY

	//Update momentanen screen:
	//Positionen verändern etc.
	currentScreen.update(delta)

	//Rendert momentanen Screen:
	//Geänderte gegenstände... rendern
	currentScreen.render()

	//render debug zeugs, wenn der debugmode aktiviert ist
	if debugMode {
		renderDebug()
	}

	//gerenderten screen sichtbar machen
	gfx.UpdateAn()

	//Jedes update soll 10 Millisekunden dauern.
	//Optimaltime sind 10ms in ns
	//Wartet Mindestens 10ns bis es den loop wiederholt
	time.Sleep(lastIterationTime.Sub(time.Now()) + optimalTime)
}

func renderDebug() {
	//Eff.: Rendert alles für den Entwickler

	fpsDebug := "fps: " + strconv.Itoa(currentFps)
	gfx.Stiftfarbe(255, 255, 255)
	gfx.SchreibeFont(5, 5, fpsDebug)
}

func SetScreen(screen Screen) {
	//Vor.: Ein screen wird übergeben
	//Eff.: Der übergebene screen wird gesetzt und später gerendert
	currentScreen = screen
	screen.init()
}

func MouseClickConsumer(taste uint8, status int8) {
	//Info: Wird ausgeführt, wenn ein Mausklick eingeht
	switch status {
	case 1:
		currentScreen.mousePress(taste, MouseX, MouseY)
	case -1:
		currentScreen.mouseRelease(taste, MouseX, MouseY)
	}
}

func TastaturConsumer(taste uint16, gedrueckt uint8, tiefe uint16) {
	//Wird ausgeführt, wenn die Tastatur gedrückt wird
	currentScreen.keyPressed(taste, gedrueckt == 1, tiefe)
}
