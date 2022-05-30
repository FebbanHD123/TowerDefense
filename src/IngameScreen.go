package main

type IngameScreen struct {
	world World
}

func CreateIngameScreen(level Level) IngameScreen {
	return IngameScreen{
		world: CreateWorld(level),
	}
}

func (s *IngameScreen) init() {

}

func (s *IngameScreen) update(deltaTime int64) {
	s.world.Update(deltaTime)
}

func (s *IngameScreen) render() {
	s.world.Render(240, 0)
}

func (s *IngameScreen) mousePress(taste uint8, mouseX, mouseY uint16) {

}

func (s *IngameScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (s *IngameScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *IngameScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {

}
