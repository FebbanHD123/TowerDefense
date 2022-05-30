package main

import (
	"fmt"
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
	if w.enemySpawnTimer.HasReached() {
		w.enemySpawnTimer.Reset()
		spawnLocation := w.level.GetRandomSpawnLocation()
		w.enemies = append(w.enemies, CreateEnemy(spawnLocation))
	}

}

func (w *World) Render(x, y uint16) {
	w.level.Render(x, y)
	w.renderEntities()
}

func (w *World) renderEntities() {
	for _, enemy := range w.enemies {
		fmt.Println("render enemy at", enemy.location)
		enemy.Render()
	}
}
