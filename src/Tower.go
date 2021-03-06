package main

import (
	"gfx"
	"time"
)

const towerSize = 64

var towerTextures map[int]Image

func InitTowerTextures() {
	towerTextures = make(map[int]Image)
	towerTextures[1] = CreateImageName("entity/tower/1.bmp")
	towerTextures[2] = CreateImageName("entity/tower/2.bmp")
	towerTextures[3] = CreateImageName("entity/tower/3.bmp")
	towerTextures[4] = CreateImageName("entity/tower/4.bmp")
	towerTextures[5] = CreateImageName("entity/tower/5.bmp")
	towerTextures[6] = CreateImageName("entity/tower/6.bmp")
}

type Tower struct {
	level      int
	location   Location
	texture    Image
	timer      Timer
	shootRange int
	damage     int
	world      *World
	bullets    []*Bullet
}

func CreateTower(world *World, level int, location Location) Tower {
	var fireRate time.Duration
	var damage int
	switch level {
	case 1:
		damage = 10
		fireRate = time.Second
	case 2:
		damage = 15
		fireRate = time.Millisecond * 800
	case 3:
		damage = 20
		fireRate = time.Millisecond * 800
	case 4:
		damage = 20
		fireRate = time.Millisecond * 600
	case 5:
		damage = 30
		fireRate = time.Millisecond * 500
	case 6:
		damage = 30
		fireRate = time.Millisecond * 250
	}
	return Tower{
		world:      world,
		level:      level,
		location:   location,
		texture:    towerTextures[level],
		damage:     damage,
		timer:      CreateTimer(fireRate),
		shootRange: GetTowerRange(level),
	}
}

func GetTowerCoasts(level int) int {
	var coasts int
	switch level {
	case 1:
		coasts = 20
	case 2:
		coasts = 35
	case 3:
		coasts = 75
	case 4:
		coasts = 150
	case 5:
		coasts = 250
	case 6:
		coasts = 400
	}
	return coasts
}

func GetTowerRange(level int) int {
	var shootRange int
	switch level {
	case 1:
		shootRange = 115
	case 2:
		shootRange = 150
	case 3:
		shootRange = 175
	case 4:
		shootRange = 200
	case 5:
		shootRange = 225
	case 6:
		shootRange = 250
	}
	return shootRange
}

func (t *Tower) GetUpgradeCoasts() int {
	return GetTowerCoasts(t.level+1) - GetTowerCoasts(t.level)
}

func (t *Tower) Update(deltaTime int64) {
	if t.timer.HasReached() {
		enemy, valid := t.getClosestEnemyInRange()
		if valid {
			t.shoot(enemy)
			t.timer.Reset()
		}
	}
	bulletsToRemove := make([]int, 0)
	for i := 0; i < len(t.bullets); i++ {
		bullet := t.bullets[i]
		bullet.Update(deltaTime)
		if bullet.goalEnemy.IsLocationInHitBox(bullet.location) {
			bullet.goalEnemy.DecreaseHealth(t.damage)
			bulletsToRemove = append(bulletsToRemove, i)
		}
	}
	for i := range bulletsToRemove {
		if bulletsToRemove[i] < len(t.bullets) {
			t.bullets = append(t.bullets[:bulletsToRemove[i]], t.bullets[bulletsToRemove[i]+1:]...)
		}
	}

}

func (t *Tower) Render() {
	for i := range t.bullets {
		t.bullets[i].Render()
	}
	t.texture.Render(t.location.x, t.location.y)
}

func (t *Tower) shoot(goal *Enemy) {
	bullet := CreateBullet(t.location, t.level, goal)
	t.bullets = append(t.bullets, &bullet)
}

func (t *Tower) RenderRange() {
	RenderTowerRange(t.location, t.level)
}

func RenderTowerRange(location Location, level int) {
	gfx.Stiftfarbe(180, 180, 180)
	gfx.Kreis(location.x+towerSize/2, location.y+towerSize/2, uint16(GetTowerRange(level)))
}

func (t *Tower) getClosestEnemyInRange() (*Enemy, bool) {
	//Vor.: -
	//Eff.: Gibt den enemy zur??ck, der am n??chsten an dem Tower ist, solange er
	//		in der Range des towers ist.
	var closest *Enemy = nil
	var distance float64 = 100000000000
	for i := range t.world.enemies {
		enemy := &t.world.enemies[i]
		d := enemy.location.Distance(t.location)
		if d < float64(t.shootRange) && d < distance {
			distance = d
			closest = enemy
		}
	}
	return closest, closest != nil
}

func (t *Tower) GetHitBox() Rect {
	return CreateRect(t.location.x, t.location.y, towerSize, towerSize)
}
