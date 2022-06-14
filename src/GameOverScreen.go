package main

type GameOverScreen struct{}

func CreateGameOverScreen() GameOverScreen {
	return GameOverScreen{}
}

func (s *GameOverScreen) update(deltaTime int64) {

}

func (s *GameOverScreen) render() {
	RenderCenteredText("Du bist durchgefallen!", width/2, height/3)
	RenderCenteredText("Die Schule wurde zerstörrt!", width/2, height/3+FontHeight+5)
}

func (s *GameOverScreen) init() {

}

func (s *GameOverScreen) mousePress(taste uint8, mouseX, mouseY uint16) {

}

func (s *GameOverScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *GameOverScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *GameOverScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {

}
