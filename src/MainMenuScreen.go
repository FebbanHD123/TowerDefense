package main

import (
	"time"
)

type MainMenuScreen struct {
	backgroundImage Image
	x, y            uint16
	startTime       time.Time
	buttons         []ButtonWidget
}

func (s *MainMenuScreen) init() {
	s.backgroundImage = CreateImageName("Background.bmp")
	s.startTime = time.Now()
	s.buttons = []ButtonWidget{
		CreateButtonWidget("Start Game", width/2-50, height/2, 200, 50, func() {
			ingameScreen := CreateIngameScreen(Levels[0])
			SetScreen(&ingameScreen)
		}),
		CreateButtonWidget("Map-Editor", width/2-50, height/2+70, 200, 50, func() {
			levelEditorScreen := CreateLevelEditor(Levels[0], s)
			//levelEditorScreen := CreateNewLevelEditor(s)
			SetScreen(&levelEditorScreen)
		}),
	}
}

func (s *MainMenuScreen) render() {
	s.backgroundImage.Render(0, 0)
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

func (s *MainMenuScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {
}
