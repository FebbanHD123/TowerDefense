package main

type Screen interface {
	update(deltaTime int64)
	render()
	init()
	mousePress(taste uint8, mouseX, mouseY uint16)
	mouseRelease(taste uint8, mouseX, mouseY uint16)
	mouseMove(mouseX, mouseY uint16, deltaX, deltaY int)
	keyPressed(taste uint16, gedrueckt bool, tiefe uint16)
}
