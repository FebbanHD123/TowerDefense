package main

import (
	"gfx"
	"strconv"
)

type Bullet struct {
	note      string
	location  Location
	goalEnemy *Enemy
	damage    int
}

func CreateBullet(location Location, level int, goalEnemy *Enemy) Bullet {
	var damage int
	switch level {
	default:
		damage = 1
	}
	return Bullet{
		note:      strconv.Itoa(level),
		location:  location,
		goalEnemy: goalEnemy,
		damage:    damage,
	}
}

func (b *Bullet) Update(deltaTime int64) {
	gk := float64(b.location.y) - float64(b.goalEnemy.location.y)
	ak := float64(b.goalEnemy.location.x) - float64(b.location.x)
	anstieg := gk / ak

	dX := 2.0
	if b.location.x >= b.goalEnemy.location.x {
		dX = -dX
	}
	dY := anstieg * dX
	b.location.x += uint16(dX * float64(deltaTime))
	b.location.y -= uint16(dY * float64(deltaTime))
}

func (b *Bullet) Render() {
	gfx.Stiftfarbe(230, 0, 0)
	RenderCenteredText(b.note, b.location.x, b.location.y-FontHeight/2)
}
