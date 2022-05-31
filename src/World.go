package main

import (
	"time"
)

type World struct {
	level           Level
	enemySpawnTimer Timer
	enemies         []Enemy
}

func CreateWorld(level Level) World {
	return World{
		level:           level,
		enemySpawnTimer: CreateTimer(time.Second * 5),
		enemies:         make([]Enemy, 0),
	}
}

func (w *World) Update(deltaTime int64) {
	if len(w.enemies) == 0 {
		spawnLocation := w.level.GetRandomSpawnLocation()
		w.enemies = append(w.enemies, CreateEnemy(spawnLocation, *w))
	}
	//if w.enemySpawnTimer.HasReached() {
	//	w.enemySpawnTimer.Reset()
	//	spawnLocation := w.level.GetRandomSpawnLocation()
	//	w.enemies = append(w.enemies, CreateEnemy(spawnLocation, *w))
	//}
	for _, enemy := range w.enemies {
		enemy.Update(deltaTime)
	}

}

func (w *World) Render(x, y uint16) {
	w.level.Render(x, y)
	w.renderEntities()
}

func (w *World) renderEntities() {
	for _, enemy := range w.enemies {
		enemy.Render()
	}
}
