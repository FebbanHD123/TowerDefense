package main

type MainMenuScreen struct {
	backgroundImage Image
	buttons         []ButtonWidget
	nameTextBox     TextBoxWidget
}

func (s *MainMenuScreen) init() {
	s.backgroundImage = CreateImageName("Background.bmp")
	s.buttons = []ButtonWidget{
		CreateButtonWidget("Start Game", width/2-50, height/2, 200, 50, func() {
			ingameScreen := CreateIngameScreen(s.nameTextBox.text, Levels[0])
			SetScreen(&ingameScreen)
		}),
		CreateButtonWidget("Map-Editor", width/2-50, height/2+70, 200, 50, func() {
			//levelEditorScreen := CreateLevelEditor(Levels[0], s)
			levelEditorScreen := CreateNewLevelEditor(s)
			SetScreen(&levelEditorScreen)
		}),
		CreateButtonWidget("Top-Liste", width/2-50, height/2+140, 200, 50, func() {
			topList := CreateTopListScreen(s)
			SetScreen(&topList)
		}),
	}
	s.nameTextBox = CreateTextBox(width/2-50, height/2+250, 200, 50)
	s.nameTextBox.SetText("Gast")
}

func (s *MainMenuScreen) render() {
	s.backgroundImage.Render(0, 0)
	for i := range s.buttons {
		s.buttons[i].Render(MouseX, MouseY)
	}
	RenderText("Nutzername:", width/2-50, height/2+250-FontHeight-5)
	s.nameTextBox.Render()
}

func (s *MainMenuScreen) update(deltaTime int64) {

}

func (s *MainMenuScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	for i := range s.buttons {
		s.buttons[i].MousePress(taste, mouseX, mouseY)
	}
	s.nameTextBox.mousePress(taste, mouseX, mouseY)
}

func (s *MainMenuScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *MainMenuScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *MainMenuScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {
	s.nameTextBox.keyPressed(taste, gedrueckt, tiefe)
}
