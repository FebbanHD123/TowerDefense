package main

import (
	"time"
)

type World struct {
	level             Level
	enemySpawnTimer   Timer
	enemies           []Enemy
	towers            []Tower
	enemyCount        int
	health, maxHealth int
	coins             int
	round             int
}

func CreateWorld(level Level) World {
	return World{
		level:           level,
		enemySpawnTimer: CreateTimer(time.Second),
		enemies:         make([]Enemy, 0),
		towers:          make([]Tower, 0),
		maxHealth:       10000,
		health:          10000,
		round:           0,
		coins:           20,
	}
}

func (w *World) Update(deltaTime int64) {
	if w.enemyCount > 0 && w.enemySpawnTimer.HasReached() {
		w.enemyCount--
		w.enemySpawnTimer.Reset()
		spawnLocation := w.level.GetRandomSpawnLocation()
		w.enemies = append(w.enemies, CreateEnemy(spawnLocation, w, int(float64(w.round)/5.0+1), 1+0.2*float64(w.round)))
	}

	enemiesToRemove := make([]int, 0)
	for i := range w.enemies {
		enemy := &w.enemies[i]
		enemy.Update(deltaTime)
		if w.level.GetRegionOfType(RTYPE_GOAL).Region.ContainsPosition(enemy.location.x, enemy.location.y) {
			enemiesToRemove = append(enemiesToRemove, i)
			w.coins--
			w.health--
		}
		if enemy.IsDead() {
			enemiesToRemove = append(enemiesToRemove, i)
			w.coins++
		}
	}
	for i := range enemiesToRemove {
		if enemiesToRemove[i] < len(w.enemies) {
			w.enemies = append(w.enemies[:enemiesToRemove[i]], w.enemies[enemiesToRemove[i]+1:]...)
		}
	}

	for i := range w.towers {
		w.towers[i].Update(deltaTime)
	}

	if w.health <= 0 {
		screen := CreateGameOverScreen()
		SetScreen(&screen)
	}

}

func (w *World) Render(x, y uint16) {
	w.level.Render(x, y)
	w.renderEntities()
}

func (w *World) renderEntities() {
	for i := range w.towers {
		w.towers[i].Render()
	}
	for i := range w.enemies {
		w.enemies[i].Render()
	}
}

func (w *World) SetEnemyCount(enemyCount int) {
	w.enemyCount = enemyCount
}

func (w *World) GetEnemyCount() int {
	return w.enemyCount
}

func (w *World) SpawnTower(location Location, level int) {
	w.towers = append(w.towers, CreateTower(w, level, location))
}
