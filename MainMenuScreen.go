package main

import (
	"time"
)

var image Image = CreateImageName("Background.bmp")

type MainMenuScreen struct {
	x, y      uint16
	startTime time.Time
	buttons   []ButtonWidget
}

func (s *MainMenuScreen) init() {
	s.startTime = time.Now()
	s.buttons = []ButtonWidget{
		CreateButtonWidget("Map-Editor", width/2-50, height/2-10, 100, 40, func() {
			levelEditorScreen := NewLevelEditor(CreateImageName("Background.bmp"))
			SetScreen(&levelEditorScreen)
		}),
	}
}

func (s *MainMenuScreen) render() {

	image.Render(0, 0)
	for i := range s.buttons {
		s.buttons[i].Render(MouseX, MouseY)
	}
}

func (s *MainMenuScreen) update(deltaTime int64) {
	s.x += 1 * uint16(deltaTime) * 2
}

func (s *MainMenuScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	for i := range s.buttons {
		s.buttons[i].MousePress(taste, mouseX, mouseY)
	}
}

func (s *MainMenuScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *MainMenuScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}
