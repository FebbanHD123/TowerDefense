package main

import (
	"gfx"
	"strconv"
	"time"
)

type IngameScreen struct {
	world       World
	towerSlots  []TowerSlot
	dragAndDrop *DragAndDropTower
}

func CreateIngameScreen(level Level) IngameScreen {
	screen := IngameScreen{
		world: CreateWorld(level),
		towerSlots: []TowerSlot{
			CreateTowerSlot(240/2-towerSlotWidth/2, 30+(20+towerSlotWidth)*0, 1),
			CreateTowerSlot(240/2-towerSlotWidth/2, 30+(20+towerSlotWidth)*1, 2),
			CreateTowerSlot(240/2-towerSlotWidth/2, 30+(20+towerSlotWidth)*2, 3),
			CreateTowerSlot(240/2-towerSlotWidth/2, 30+(20+towerSlotWidth)*3, 4),
			CreateTowerSlot(240/2-towerSlotWidth/2, 30+(20+towerSlotWidth)*4, 5),
			CreateTowerSlot(240/2-towerSlotWidth/2, 30+(20+towerSlotWidth)*5, 6),
		},
	}
	return screen
}

func (s *IngameScreen) init() {

}

func (s *IngameScreen) update(deltaTime int64) {
	s.world.Update(deltaTime)

	//wenn keine enemys mehr erstellt werden sollen, was bedeutet, dass alle
	//enemys des levels gespawnt wurden und alle enemys nicht mehr auf der welt sind,
	//wird das nächste level eingeleitet
	if s.world.enemyCount <= 0 && (s.world.round == 0 || len(s.world.enemies) == 0) {
		s.world.round++
		s.world.enemySpawnTimer = CreateTimer(time.Millisecond*1000 - time.Duration(s.world.round)*25*time.Millisecond)
		s.world.SetEnemyCount(s.world.round * 3)
	}

	if s.dragAndDrop != nil {
		s.dragAndDrop.Update()
	}
}

func (s *IngameScreen) render() {
	s.world.Render(240, 0)
	s.renderHud()
}

func (s *IngameScreen) renderHud() {
	//Vor.: -
	//Eff.: Rendert das Hud, also alles was sich über der Welt befindet, wie z.B. die Level anzeige

	gfx.Stiftfarbe(50, 50, 50)
	gfx.Vollrechteck(0, 0, 240, height)
	gfx.Vollrechteck(240+800, 0, 240, height)

	//Level anzeige
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("Runde: "+strconv.Itoa(s.world.round), 1050+220/2, 20)

	//1040px: rechter anfang

	//Leben anzeige
	gfx.Stiftfarbe(100, 100, 100)
	gfx.Vollrechteck(1050, 50, 220, 50)

	percentage := float64(s.world.health) / float64(s.world.maxHealth)
	if percentage > 1 {
		percentage = 1
	}

	gfx.Stiftfarbe(uint8((1-percentage)*255), 255-uint8((1-percentage)*255), 0)
	gfx.Vollrechteck(1050, 50, uint16(percentage*220), 50)
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText(strconv.Itoa(s.world.health)+"/"+strconv.Itoa(s.world.maxHealth)+" HP", 1050+220/2, 50+50/2-FontHeight/2)

	//Guthaben
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("Cash: "+strconv.Itoa(s.world.coins)+"€", 1050+220/2, 120)

	//tower slots rechts
	for i := range s.towerSlots {
		s.towerSlots[i].Render()
	}

	if s.dragAndDrop != nil {
		s.dragAndDrop.Render(&s.world)
	}

}

func (s *IngameScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	if s.dragAndDrop == nil {
		for i := range s.towerSlots {
			slot := s.towerSlots[i]
			if !slot.IsMouseOver() {
				continue
			}
			dragAndDrop := CreateDragAndDropTower(slot)
			s.dragAndDrop = &dragAndDrop
		}
	}
}

func (s *IngameScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {
	if s.dragAndDrop != nil {
		if s.dragAndDrop.CanPlaceTower(&s.world) {
			s.world.coins -= s.dragAndDrop.towerSlot.coasts
			s.world.SpawnTower(s.dragAndDrop.location, s.dragAndDrop.towerSlot.level)
		}
		s.dragAndDrop = nil
	}
}

func (s *IngameScreen) mouseMove(mouseX, mouseY uint16, deltaX, deltaY int) {

}

func (s *IngameScreen) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {

}
