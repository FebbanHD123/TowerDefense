package main

import (
	"gfx"
	"strconv"
)

const towerSlotWidth = enemyWidth + 65

type TowerSlot struct {
	location Location
	texture  Image
	level    int
	coasts   int
}

func CreateTowerSlot(x, y uint16, level int) TowerSlot {
	var image Image
	if level > len(towerTextures) {
		image = towerTextures[len(towerTextures)]
	} else {
		image = towerTextures[level]
	}
	var coasts int
	switch level {
	case 1:
		coasts = 20
	case 2:
		coasts = 35
	case 3:
		coasts = 75
	case 4:
		coasts = 100
	case 5:
		coasts = 250
	case 6:
		coasts = 200
	}

	return TowerSlot{
		location: CreateLocation(x, y),
		texture:  image,
		level:    level,
		coasts:   coasts,
	}
}

func (t *TowerSlot) Render() {
	gfx.Stiftfarbe(70, 70, 70)
	if t.IsMouseOver() {
		gfx.Vollrechteck(t.location.x, t.location.y, towerSlotWidth, towerSlotWidth)
	}
	gfx.Stiftfarbe(30, 30, 30)
	gfx.Rechteck(t.location.x, t.location.y, towerSlotWidth, towerSlotWidth)
	t.texture.Render(t.location.x+(towerSlotWidth-towerWidth)/2, t.location.y+(towerSlotWidth-towerWidth)/4)
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("  "+strconv.Itoa(t.coasts)+"€", t.location.x+towerSlotWidth/2, t.location.y+towerSlotWidth-FontHeight-5)
}

func (t *TowerSlot) IsMouseOver() bool {
	rect := CreateRect(t.location.x, t.location.y, towerSlotWidth, towerSlotWidth)
	return rect.ContainsPosition(MouseX, MouseY)
}
