package main

import (
	"gfx"
)

type DragAndDropTower struct {
	location  Location
	towerSlot TowerSlot
}

func CreateDragAndDropTower(slot TowerSlot) DragAndDropTower {
	//Eff.: Gibt ein Objekt der Klasse DragAndDrop mit den übergebenen Werten zurück
	return DragAndDropTower{
		location:  CreateLocation(0, 0),
		towerSlot: slot,
	}
}

func (t *DragAndDropTower) Update() {
	//Eff.: Setzt die Position auf die der Maus
	t.location.x = MouseX - towerSize/2
	t.location.y = MouseY - towerSize/2
}

func (t *DragAndDropTower) Render(world World) {
	//Eff.: Rendert den Tower und seine Range
	if t.CanPlaceTower(world) {
		gfx.Stiftfarbe(100, 255, 100)
	} else {
		gfx.Stiftfarbe(255, 100, 100)
	}
	gfx.Vollrechteck(t.location.x-10, t.location.y-10, towerSize+20, towerSize+20)
	RenderTowerRange(t.location, t.towerSlot.level)
	t.towerSlot.texture.Render(t.location.x, t.location.y)
}

func (t *DragAndDropTower) CanPlaceTower(world World) bool {
	//Eff.: Gibt zurück, ob der Tower an der momentanen Position gesetzt werden kann.
	//		Also, wenn er in einer Defense-Region ist und keinen anderen Tower trifft
	if world.coins < t.towerSlot.coasts {
		return false
	}
	defenseRegions := world.level.GetRegionsOfType(RTYPE_DEFENSE)
	for x := t.location.x; x < t.location.x+towerSize; x++ {
		for y := t.location.y; y < t.location.y+towerSize; y++ {
			var CanPlace bool = false
			for i := range defenseRegions {
				defenseRegion := defenseRegions[i]
				if defenseRegion.Region.ContainsPosition(x, y) {
					CanPlace = true
					break
				}
			}
			if CanPlace {
				for i := range world.towers {
					tower := world.towers[i]
					if tower.GetHitBox().ContainsPosition(x, y) {
						CanPlace = false
						break
					}
				}
			}
			if !CanPlace {
				return false
			}
		}
	}
	return true
}
