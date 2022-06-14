package main

type GameOverScreen struct {
	buttons []ButtonWidget
}

func CreateGameOverScreen() GameOverScreen {
	return GameOverScreen{}
}

func (s *GameOverScreen) render() {
	//Eff.: Rendert den Screen
	RenderCenteredText("Du bist durchgefallen!", width/2, height/3)
	RenderCenteredText("Die Schule wurde zerstörrt!", width/2, height/3+FontHeight+5)
	for i := range s.buttons {
		s.buttons[i].Render(MouseX, MouseY)
	}
}

func (s *GameOverScreen) init() {
	//Eff. Initialisiert den screen (Buttons)
	s.buttons = []ButtonWidget{
		CreateButtonWidget("Hauptmenü", width/2-100, height/2+50, 200, 50, func() {
			screen := MainMenuScreen{}
			SetScreen(&screen)
		}),
		CreateButtonWidget("Erneut versuchen", width/2-100, height/2+200, 200, 50, func() {
			ingameScreen := CreateIngameScreen(Levels[0])
			SetScreen(&ingameScreen)
		}),
	}
}

func (s *GameOverScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	for i := range s.buttons {
		s.buttons[i].MousePress(taste, mouseX, mouseY)
	}
}

func (s *GameOverScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {
	//Nicht benutzt aber nötig
}

func (s *GameOverScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {
	//Nicht benutzt aber nötig
}

func (s *GameOverScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {
	//Nicht benutzt aber nötig
}

func (s *GameOverScreen) update(deltaTime int64) {
	//Nicht benutzt aber nötig
}
