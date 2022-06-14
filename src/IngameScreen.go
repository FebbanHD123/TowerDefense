package main

import (
	"gfx"
	"strconv"
	"time"
)

type IngameScreen struct {
	world                     World
	towerSlots                []TowerSlot
	dragAndDrop               *DragAndDropTower
	slectedTower              *Tower
	upgradeButton, sellButton ButtonWidget
}

func CreateIngameScreen(level Level) IngameScreen {
	//Eff.: Gibt ein Objekt der Klasse IngameScreen zurück
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
	s.upgradeButton = CreateButtonWidget("Upgrade", 1050+220/2-150/2, 200, 150, 50, func() {
		coasts := s.slectedTower.GetUpgradeCoasts()
		if s.world.coins >= coasts {
			s.world.coins -= coasts
		}
		s.world.SpawnTower(s.slectedTower.location, s.slectedTower.level+1)
		s.world.RemoveTower(s.slectedTower.location)
		s.slectedTower = nil
	})
	s.sellButton = CreateButtonWidget("Sell", 1050+220/2-150/2, 200+70, 150, 50, func() {
		coins := GetTowerCoasts(s.slectedTower.level) / 2
		s.world.coins += coins
		s.world.RemoveTower(s.slectedTower.location)
	})
}

func (s *IngameScreen) update(deltaTime int64) {
	s.world.Update(deltaTime)

	//wenn keine enemys mehr erstellt werden sollen, was bedeutet, dass alle
	//enemys des levels gespawnt wurden und alle enemys nicht mehr auf der welt sind,
	//wird das nächste level eingeleitet
	if s.world.enemyCount <= 0 && (s.world.round == 0 || len(s.world.enemies) == 0) {
		s.world.round++
		if s.world.round%10 == 0 && s.world.health < s.world.maxHealth {
			s.world.health++
		}
		s.world.enemySpawnTimer = CreateTimer(time.Millisecond*1000 - time.Duration(s.world.round)*25*time.Millisecond)
		s.world.SetEnemyCount(s.world.round * 3)
	}

	if s.dragAndDrop != nil {
		s.dragAndDrop.Update()
	}

	if s.slectedTower != nil {
		s.upgradeButton.setActivated(s.world.coins >= s.slectedTower.GetUpgradeCoasts())
		s.upgradeButton.SetTitle("Upgrade: " + strconv.Itoa(s.slectedTower.GetUpgradeCoasts()) + "€")
		s.sellButton.setActivated(true)
		s.sellButton.SetTitle("Sell: " + strconv.Itoa(GetTowerCoasts(s.slectedTower.level)/2) + "€")
	} else {
		s.upgradeButton.setActivated(false)
		s.upgradeButton.SetTitle("Upgrade")
		s.sellButton.setActivated(false)
		s.sellButton.SetTitle("Sell")
	}
}

func (s *IngameScreen) render() {
	s.world.Render(240, 0)
	if s.slectedTower != nil {
		s.slectedTower.RenderRange()
	}
	s.renderHud()
}

func (s *IngameScreen) renderHud() {
	//Vor.: -
	//Eff.: Rendert das Hud, also alles was sich über der Welt befindet, wie z.B. die Level anzeige

	gfx.Stiftfarbe(50, 50, 50)
	gfx.Vollrechteck(0, 0, 240, height)
	gfx.Vollrechteck(240+800, 0, 240, height)

	//1040px: rechter anfang

	//Leben anzeige
	gfx.Stiftfarbe(100, 100, 100)
	gfx.Vollrechteck(1050, 20, 220, 50)

	percentage := float64(s.world.health) / float64(s.world.maxHealth)
	if percentage > 1 {
		percentage = 1
	}

	gfx.Stiftfarbe(uint8((1-percentage)*255), 255-uint8((1-percentage)*255), 0)
	gfx.Vollrechteck(1050, 20, uint16(percentage*220), 50)
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText(strconv.Itoa(s.world.health)+"/"+strconv.Itoa(s.world.maxHealth)+" HP", 1050+220/2, 20+20/2-FontHeight/2)

	//Runde anzeigen
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("Runde: "+strconv.Itoa(s.world.round), 1050+220/2, 90)

	//Score
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("Score: "+strconv.Itoa(s.world.score), 1050+220/2, 120)

	//Guthaben
	gfx.Stiftfarbe(255, 255, 255)
	RenderCenteredText("Cash: "+strconv.Itoa(s.world.coins)+"€", 1050+220/2, 150)

	//tower slots rechts
	for i := range s.towerSlots {
		s.towerSlots[i].Render()
	}

	if s.dragAndDrop != nil {
		s.dragAndDrop.Render(s.world)
	}

	s.upgradeButton.Render(MouseX, MouseY)
	s.sellButton.Render(MouseX, MouseY)
}

func (s *IngameScreen) mousePress(taste uint8, mouseX, mouseY uint16) {
	s.upgradeButton.MousePress(taste, mouseX, mouseY)
	s.sellButton.MousePress(taste, mouseX, mouseY)
	if s.dragAndDrop == nil {
		for i := range s.towerSlots {
			slot := s.towerSlots[i]
			if !slot.IsMouseOver() || slot.coasts > s.world.coins {
				continue
			}
			dragAndDrop := CreateDragAndDropTower(slot)
			s.dragAndDrop = &dragAndDrop
			return
		}
	}
	for i := range s.world.towers {
		tower := &s.world.towers[i]
		if tower.GetHitBox().ContainsPosition(mouseX, mouseY) {
			s.slectedTower = tower
			return
		}
	}
	s.slectedTower = nil
}

func (s *IngameScreen) mouseRelease(taste uint8, mouseX, mouseY uint16) {
	if s.dragAndDrop != nil {
		if s.dragAndDrop.CanPlaceTower(s.world) {
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
