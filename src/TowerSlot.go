package main

import (
	"gfx"
	"strconv"
)

const towerSlotWidth = enemySize + 65

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

	return TowerSlot{
		location: CreateLocation(x, y),
		texture:  image,
		level:    level,
		coasts:   GetTowerCoasts(level),
	}
}

func (t *TowerSlot) Render() {
	gfx.Stiftfarbe(70, 70, 70)
	if t.IsMouseOver() {
		gfx.Vollrechteck(t.location.x, t.location.y, towerSlotWidth, towerSlotWidth)
	}
	gfx.Stiftfarbe(30, 30, 30)
	gfx.Rechteck(t.location.x, t.location.y, towerSlotWidth, towerSlotWidth)
	t.texture.Render(t.location.x+(towerSlotWidth-towerSize)/2, t.location.y+(towerSlotWidth-towerSize)/4)
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("  "+strconv.Itoa(t.coasts)+"€", t.location.x+towerSlotWidth/2, t.location.y+towerSlotWidth-FontHeight-5)
}

func (t *TowerSlot) IsMouseOver() bool {
	rect := CreateRect(t.location.x, t.location.y, towerSlotWidth, towerSlotWidth)
	return rect.ContainsPosition(MouseX, MouseY)
}
