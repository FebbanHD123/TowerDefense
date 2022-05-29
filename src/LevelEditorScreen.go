package main

import (
	"gfx"
)

//Vor.: Es wird ein Level mit mindestens einem
//Hintergrundbild übergeben

//Eff.: Erstellen von Leveln

type LevelEditorScreen struct {
	level                                                                                      Level
	makierung                                                                                  Makierung
	nameTextBox, backgroundImageTextBox                                                        TextBoxWidget
	createButton, cancelButton, setPathButton, setDefenseButton, setGoalButton, setSpawnButton ButtonWidget
	previousScreen                                                                             Screen
	showRegions                                                                                bool
	toggleShowRegionsButton                                                                    ButtonWidget
}

func CreateNewLevelEditor(previousScreen Screen) LevelEditorScreen {
	//Vor.: previous screen wird übergeben und ist ein vailder screen
	//Eff.: Ein Neuer Editor wird zurückgegeben
	return LevelEditorScreen{
		makierung:      CreateMakierung(),
		previousScreen: previousScreen,
	}
}

func CreateLevelEditor(level Level, previousScreen Screen) LevelEditorScreen {
	//Vor.: previous screen wird übergeben und ist ein vailder screen
	//Eff.: Eiin editor mit den übergebenen level wird zurückgegeben
	return LevelEditorScreen{
		level:          level,
		makierung:      CreateMakierung(),
		previousScreen: previousScreen,
	}
}

func (s *LevelEditorScreen) init() {
	s.showRegions = true
	s.createButton = CreateButtonWidget("Speichern", width-220, height-70, 200, 50, func() {
		s.level.DeleteOldFileIfExists()
		s.level.Name = s.nameTextBox.text
		s.level.BackGroundImage = CreateImageName(s.backgroundImageTextBox.text)
		s.level.Save()
		SetScreen(s.previousScreen)
	})
	s.createButton.setActivated(false)

	s.cancelButton = CreateButtonWidget("Abbrechen", 20, height-70, 200, 50, func() {
		SetScreen(s.previousScreen)
	})
	if s.previousScreen == nil {
		s.cancelButton.setActivated(false)
	}

	//y = 60-110
	s.nameTextBox = CreateTextBox(20, 60+FontHeight, 200, 50)
	s.nameTextBox.SetText(s.level.Name)

	s.backgroundImageTextBox = CreateTextBox(20, 500+FontHeight, 200, 50)
	s.backgroundImageTextBox.SetText(s.level.BackGroundImage.Id.Path)

	//y = 150 - 200
	s.setPathButton = CreateButtonWidget("Path-Region", 20, 150, 200, 50, func() {
		levelRegion := CreateLevelRegion(RTYPE_PATH, s.makierung.ToRect())
		s.level.AddRegion(levelRegion)
		s.makierung.Reset()
	})

	//y = 230 - 280
	s.setGoalButton = CreateButtonWidget("Goal-Region", 20, 230, 200, 50, func() {
		levelRegion := CreateLevelRegion(RTYPE_GOAL, s.makierung.ToRect())
		s.level.AddRegion(levelRegion)
		s.makierung.Reset()
	})
	//y = 310 - 360
	s.setDefenseButton = CreateButtonWidget("Defense-Region", 20, 310, 200, 50, func() {
		levelRegion := CreateLevelRegion(RTYPE_DEFENSE, s.makierung.ToRect())
		s.level.AddRegion(levelRegion)
		s.makierung.Reset()
	})
	//y = 390 - 440
	s.setSpawnButton = CreateButtonWidget("Spawn-Region", 20, 390, 200, 50, func() {
		levelRegion := CreateLevelRegion(RTYPE_SPAWN, s.makierung.ToRect())
		s.level.AddRegion(levelRegion)
		s.makierung.Reset()
	})
	s.toggleShowRegionsButton = CreateButtonWidget("Loading", 800+240+20, 50, 200, 50, func() {
		s.showRegions = !s.showRegions
	})
}

func (s *LevelEditorScreen) update(deltaTime int64) {
	s.nameTextBox.Update()
	s.backgroundImageTextBox.Update()
	s.createButton.setActivated(s.IsComplete())
	s.setPathButton.setActivated(s.makierung.IsFinished())
	s.setDefenseButton.setActivated(s.makierung.IsFinished())
	s.setGoalButton.setActivated(s.makierung.IsFinished())
	s.setSpawnButton.setActivated(s.makierung.IsFinished())

	if !s.backgroundImageTextBox.isEmpty() {
		s.level.BackGroundImage = CreateImageName(s.backgroundImageTextBox.text)
	}

	if !s.showRegions {
		s.toggleShowRegionsButton.SetTitle("Regionen Anzeigen")
	} else {
		s.toggleShowRegionsButton.SetTitle("Regionen Ausblenden")
	}
}

func (s *LevelEditorScreen) render() {
	gfx.Stiftfarbe(50, 50, 50)
	gfx.Cls()

	gfx.Stiftfarbe(255, 255, 255)
	RenderText("Name:", 20, 60-FontHeight+5)
	RenderText("Background-Name:", 20, 500-FontHeight+5)

	s.nameTextBox.Render()
	s.createButton.Render(MouseX, MouseY)
	s.setPathButton.Render(MouseX, MouseY)
	s.setDefenseButton.Render(MouseX, MouseY)
	s.setGoalButton.Render(MouseX, MouseY)
	s.setSpawnButton.Render(MouseX, MouseY)
	s.cancelButton.Render(MouseX, MouseY)
	s.backgroundImageTextBox.Render()
	s.toggleShowRegionsButton.Render(MouseX, MouseY)

	gfx.Stiftfarbe(150, 0, 0)
	RenderCenteredText("No BackgroundImage selected!", width/2, height/2)
	gfx.Stiftfarbe(255, 255, 255)
	s.level.Render(240, 0)
	if s.showRegions {
		for i := range s.level.GetRegions() {
			switch s.level.GetRegions()[i].Type {
			case RTYPE_PATH:
				gfx.Stiftfarbe(0, 0, 100)
			case RTYPE_DEFENSE:
				gfx.Stiftfarbe(100, 0, 0)
			case RTYPE_GOAL:
				gfx.Stiftfarbe(0, 100, 0)
			case RTYPE_SPAWN:
				gfx.Stiftfarbe(100, 100, 0)
			default:
				gfx.Stiftfarbe(0, 0, 0)
			}
			rect := s.level.GetRegions()[i].Region
			gfx.Vollrechteck(rect.X, rect.Y, rect.Width, rect.Height)
		}
	}
	gfx.Stiftfarbe(255, 255, 255)
	if s.makierung.IsFinished() {
		rect := s.makierung.ToRect()
		gfx.Vollrechteck(rect.X, rect.Y, rect.Width, rect.Height)
	}
}

func (s *LevelEditorScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	if CreateRect(240, 0, 800, height).ContainsPosition(MouseX, MouseY) {
		s.makierung.AddPos(mouseX, mouseY)
	}
	s.nameTextBox.mousePress(taste, mouseX, mouseY)
	s.createButton.MousePress(taste, mouseX, mouseY)
	s.setPathButton.MousePress(taste, mouseX, mouseY)
	s.setGoalButton.MousePress(taste, mouseX, mouseY)
	s.setDefenseButton.MousePress(taste, mouseX, mouseY)
	s.setSpawnButton.MousePress(taste, mouseX, mouseY)
	s.cancelButton.MousePress(taste, mouseX, mouseY)
	s.backgroundImageTextBox.mousePress(taste, mouseX, mouseY)
	s.toggleShowRegionsButton.MousePress(taste, mouseX, mouseY)
}

func (s *LevelEditorScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *LevelEditorScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *LevelEditorScreen) keyPressed(taste uint16, gedrueck bool, tiefe uint16) {
	s.nameTextBox.keyPressed(taste, gedrueck, tiefe)
	s.backgroundImageTextBox.keyPressed(taste, gedrueck, tiefe)
}

func (s *LevelEditorScreen) IsComplete() bool {
	if s.nameTextBox.isEmpty() || s.backgroundImageTextBox.isEmpty() || !s.level.HasAllRequiredRegions() {
		return false
	}
	return true
}
