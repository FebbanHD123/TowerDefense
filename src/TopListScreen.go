package main

import (
	"gfx"
	"strconv"
)

type TopListScreen struct {
	loaded         bool
	error          bool
	topListe       []Player
	cancelButton   ButtonWidget
	previousScreen Screen
}

func CreateTopListScreen(previousScreen Screen) TopListScreen {
	return TopListScreen{
		loaded:         false,
		previousScreen: previousScreen,
	}
}

func (t *TopListScreen) update(deltaTime int64) {

}

func (t *TopListScreen) render() {
	gfx.Stiftfarbe(255, 255, 255)
	if t.error {
		RenderCenteredText("Error loading topliste...", width/2, height/2)
	}
	if !t.loaded {
		RenderCenteredText("Lade Top-Liste...", width/2, height/2)
		return
	}
	var y uint16 = 20
	for i := range t.topListe {
		player := t.topListe[i]
		RenderCenteredText(strconv.Itoa(i+1)+". "+player.Name+" HighScore: "+strconv.Itoa(player.HighScore), width/2, y)
		y += FontHeight + 20
	}
	t.cancelButton.Render(MouseX, MouseY)
}

func (t *TopListScreen) init() {
	go func() {
		topListe, err := GetTopListFromAPI()
		if err != nil {
			t.error = true
			return
		}
		t.loaded = true
		t.topListe = topListe
	}()
	t.cancelButton = CreateButtonWidget("Hauptmenü", 20, height-70, 200, 50, func() {
		SetScreen(t.previousScreen)
	})
}

func (t *TopListScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	t.cancelButton.MousePress(taste, mouseX, mouseY)
}

func (t *TopListScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {

}

func (t *TopListScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (t *TopListScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {

}
