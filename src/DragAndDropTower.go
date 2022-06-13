package main

import (
	"gfx"
)

type DragAndDropTower struct {
	location  Location
	towerSlot TowerSlot
}

func CreateDragAndDropTower(slot TowerSlot) DragAndDropTower {
	return DragAndDropTower{
		location:  CreateLocation(0, 0),
		towerSlot: slot,
	}
}

func (t *DragAndDropTower) Update() {
	t.location.x = MouseX - towerWidth/2
	t.location.y = MouseY - towerWidth/2
}

func (t *DragAndDropTower) Render(world *World) {
	if t.CanPlaceTower(world) {
		gfx.Stiftfarbe(100, 255, 100)
	} else {
		gfx.Stiftfarbe(255, 100, 100)
	}
	gfx.Vollrechteck(t.location.x-10, t.location.y-10, towerWidth+20, towerWidth+20)
	gfx.Stiftfarbe(180, 180, 180)
	gfx.Kreis(t.location.x+towerWidth/2, t.location.y+towerWidth/2, uint16(GetTowerRange(t.towerSlot.level)))
	t.towerSlot.texture.Render(t.location.x, t.location.y)
}

func (t *DragAndDropTower) CanPlaceTower(world *World) bool {
	if world.coins < t.towerSlot.coasts {
		return false
	}
	defenseRegions := world.level.GetRegionsOfType(RTYPE_DEFENSE)
	for x := t.location.x; x < t.location.x+towerWidth; x++ {
		for y := t.location.y; y < t.location.y+towerWidth; y++ {
			var InDefenseRegion bool = false
			for i := range defenseRegions {
				defenseRegion := defenseRegions[i]
				if defenseRegion.Region.ContainsPosition(x, y) {
					InDefenseRegion = true
					break
				}
			}
			if !InDefenseRegion {
				return false
			}
		}
	}
	return true
}
