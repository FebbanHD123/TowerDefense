package main

import (
	"fmt"
	"gfx"
)

//Vor.: Es wird ein Level mit mindestens einem
//Hintergrundbild übergeben

//Eff.: Erstellen von Leveln

type LevelEditorScreen struct {
	level       Level
	makierung   Makierung
	buttons     []ButtonWidget
	nameTextBox TextBox
}

func NewLevelEditor(backGroundImage Image) LevelEditorScreen {
	return LevelEditorScreen{
		level: Level{
			BackGroundImage: backGroundImage,
		},
		makierung: CreateMakierung(),
	}
}

func (s *LevelEditorScreen) update(deltaTime int64) {
	s.nameTextBox.Update()
}

func (s *LevelEditorScreen) render() {
	gfx.Stiftfarbe(50, 50, 50)
	gfx.Cls()

	gfx.Stiftfarbe(255, 255, 255)
	RenderText("Map-Name:", 20, 60-5-FontHeight)
	s.nameTextBox.Render()
	for _, button := range s.buttons {
		button.Render(MouseX, MouseY)
	}

	s.level.Render(240, 0)
	for i := range s.level.GetRegions() {
		switch s.level.GetRegions()[i].Type {
		case RTYPE_PATH:
			gfx.Stiftfarbe(0, 0, 100)
		case RTYPE_DEFENSE:
			gfx.Stiftfarbe(100, 0, 0)
		case RTYPE_GOAL:
			gfx.Stiftfarbe(0, 100, 0)
		default:
			gfx.Stiftfarbe(0, 0, 0)
		}
		rect := s.level.GetRegions()[i].Region
		gfx.Vollrechteck(rect.X, rect.Y, rect.Width, rect.Height)
	}
	gfx.Stiftfarbe(255, 255, 255)
	if s.makierung.IsFinished() {
		rect := s.makierung.ToRect()
		gfx.Vollrechteck(rect.X, rect.Y, rect.Width, rect.Height)
	}
}

func (s *LevelEditorScreen) init() {
	createButton := CreateButtonWidget("Erstellen", width-220, height-70, 200, 50, func() {
		fmt.Println("Erstellen!")
	})
	createButton.setActivated(false)
	s.buttons = []ButtonWidget{
		createButton,
	}
	s.nameTextBox = CreateTextBox(20, 60, 200, 50)
}

func (s *LevelEditorScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	if CreateRect(240, 0, 800, height).ContainsPosition(MouseX, MouseY) {
		s.makierung.AddPos(mouseX, mouseY)
	}
	s.nameTextBox.mousePress(taste, mouseX, mouseY)
}

func (s *LevelEditorScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *LevelEditorScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *LevelEditorScreen) keyPressed(taste uint16, gedrueck bool, tiefe uint16) {
	s.nameTextBox.keyPressed(taste, gedrueck, tiefe)
	if gedrueck && s.makierung.IsFinished() {
		var regionType int
		switch taste {
		case uint16("1"[0]):
			//weg
			regionType = RTYPE_PATH
		case uint16("2"[0]):
			//defense
			regionType = RTYPE_DEFENSE
			fmt.Println("defense")
		case uint16("3"[0]):
			//goal
			regionType = RTYPE_GOAL
			fmt.Println("goal")

		default:
			fmt.Println("1 = weg, 2 = defense, 3 = ziel")
			return
		}
		levelRegion := CreateLevelRegion(regionType, s.makierung.ToRect())
		s.level.AddRegion(levelRegion)
		s.makierung.Reset()
	}
}
