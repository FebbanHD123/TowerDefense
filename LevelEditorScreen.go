package main

import (
	"gfx"
)

//Vor.: Es wird ein Level mit mindestens einem
//Hintergrundbild übergeben

//Eff.: Erstellen von Leveln

type LevelEditorScreen struct {
	level     Level
	makierung Makierung
}

func NewLevelEditor(backGroundImage Image) LevelEditorScreen {
	return LevelEditorScreen{
		level: Level{
			BackGroundImage: backGroundImage,
		},
		makierung: NeueMakierung(),
	}
}

func (s *LevelEditorScreen) update(deltaTime int64) {

}

func (s *LevelEditorScreen) render() {
	s.level.Render(0, 0)
	if s.makierung.IsFinished() {
		rect := s.makierung.ToRect()
		gfx.Vollrechteck(rect.X, rect.Y, rect.Width, rect.Height)
	}
}

func (s *LevelEditorScreen) init() {

}

func (s *LevelEditorScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	s.makierung.AddPos(mouseX, mouseY)
}

func (s *LevelEditorScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *LevelEditorScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}
