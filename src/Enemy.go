package main

const enemyWidth = 16

var enemyLevelTextures map[int]Image

func InitEnemyTextures() {
	enemyLevelTextures = make(map[int]Image)
	enemyLevelTextures[1] = CreateImageName("entity/enemy_1.bmp")
}

type Enemy struct {
	location Location
	level    int
}

func CreateEnemy(spawnLocation Location) Enemy {
	return Enemy{
		location: spawnLocation,
		level:    1,
	}
}

func (e *Enemy) Render() {
	image := enemyLevelTextures[e.level]
	image.Render(e.location.x-enemyWidth/2, e.location.y-enemyWidth/2)

}
