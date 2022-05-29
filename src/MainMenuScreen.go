package main

import (
	"time"
)

var image Image = CreateImageName("Background.bmp")

type MainMenuScreen struct {
	x, y      uint16
	startTime time.Time
	buttons   []ButtonWidget
	textbox   TextBoxWidget
}

func (s *MainMenuScreen) init() {
	s.startTime = time.Now()
	s.buttons = []ButtonWidget{
		CreateButtonWidget("Map-Editor", width/2-50, height/2-10, 100, 40, func() {
			levelEditorScreen := CreateLevelEditor(Levels[0], s)
			//levelEditorScreen := CreateNewLevelEditor(s)
			SetScreen(&levelEditorScreen)
		}),
	}
	s.textbox = CreateTextBox(width/2-100, height/2+100, 200, 50)
}

func (s *MainMenuScreen) render() {

	image.Render(0, 0)
	s.textbox.Render()
	for i := range s.buttons {
		s.buttons[i].Render(MouseX, MouseY)
	}
}

func (s *MainMenuScreen) update(deltaTime int64) {
	s.x += 1 * uint16(deltaTime) * 2
	s.textbox.Update()
}

func (s *MainMenuScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	for i := range s.buttons {
		s.buttons[i].MousePress(taste, mouseX, mouseY)
	}
	s.textbox.mousePress(taste, mouseX, mouseY)
}

func (s *MainMenuScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *MainMenuScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *MainMenuScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {
	s.textbox.keyPressed(taste, gedrueckt, tiefe)
}
