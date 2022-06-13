package main

import (
	"fmt"
	"gfx"
)

const enemyWidth = 30

var enemyLevelTextures map[int]Image

func InitEnemyTextures() {
	enemyLevelTextures = make(map[int]Image)
	enemyLevelTextures[1] = CreateImageName("entity/enemy/1.bmp")
	enemyLevelTextures[2] = CreateImageName("entity/enemy/2.bmp")
	enemyLevelTextures[3] = CreateImageName("entity/enemy/3.bmp")
}

type Enemy struct {
	currentWaypoint    int
	ProcessedWayPoints []int
	location           *Location
	level              int
	pathfinder         Pathfinder
	world              *World
	path               []Location
	dead               bool
	health             int
	maxHealth          int
	speed              float64
}

func CreateEnemy(spawnLocation Location, world *World, level int, speed float64) Enemy {
	maxHealth := 10 + 10*level
	enemy := Enemy{
		location:   &spawnLocation,
		level:      level,
		world:      world,
		pathfinder: CreatePathFinder(world.level),
		dead:       false,
		health:     maxHealth,
		maxHealth:  maxHealth,
		speed:      speed,
	}
	enemy.initPath()
	return enemy
}

func (e *Enemy) initPath() {
	processedWayPoints := make([]LevelRegion, 0)
	path := make([]Location, 0)
	path = append(path, *e.location)
	for {
		currentLocation := path[len(path)-1]

		if e.world.level.GetRegionOfType(RTYPE_GOAL).Region.ContainsPosition(currentLocation.x, currentLocation.y) {
			break
		}
		wayPoint := e.getClosestWayPoint(currentLocation, processedWayPoints)
		result, found := e.pathfinder.FindPath(currentLocation, wayPoint.Region.GetRandomLocation())
		if !found {
			fmt.Println("No path found to goal!")
			break
		}
		for _, node := range result.nodes {
			path = append(path, node.location)
		}
		processedWayPoints = append(processedWayPoints, wayPoint)
	}
	e.path = path
}

func (e *Enemy) Update(deltaTime int64) {
	nodeIndex := int(float64(deltaTime) * e.speed)
	if nodeIndex >= len(e.path) {
		nodeIndex = len(e.path) - 1
	}
	e.location.x = e.path[nodeIndex].x
	e.location.y = e.path[nodeIndex].y

	e.path = e.path[nodeIndex:]

}

func (e *Enemy) getClosestWayPoint(currentLocation Location, processedWayPoint []LevelRegion) LevelRegion {
	var closest int = -1
	var closestDistance float64
	regions := e.world.level.GetRegionsOfType(RTYPE_WAYPOINT)
	for i, region := range regions {
		if e.IsWayPointProcessed(region, processedWayPoint) {
			continue
		}
		distance := currentLocation.Distance(region.Region.GetRandomLocation())
		if closest == -1 || (closestDistance > distance) {
			closest = i
			closestDistance = distance
		}
	}
	return regions[closest]
}

func (e *Enemy) IsWayPointProcessed(wayPoint LevelRegion, processedWayPoints []LevelRegion) bool {
	for _, point := range processedWayPoints {
		if point.equals(wayPoint) {
			return true
		}
	}
	return false
}

func (e *Enemy) Render() {
	var image Image
	if e.level > len(enemyLevelTextures) {
		image = enemyLevelTextures[len(enemyLevelTextures)]
	} else {
		image = enemyLevelTextures[e.level]
	}
	image.Render(e.location.x-enemyWidth/2, e.location.y-enemyWidth/2)
	if e.health < e.maxHealth {
		e.renderHealthBar()
	}
}

func (e *Enemy) renderHealthBar() {
	gfx.Stiftfarbe(100, 100, 100)
	gfx.Vollrechteck(e.location.x-enemyWidth/2, e.location.y-enemyWidth, enemyWidth, 3)
	gfx.Stiftfarbe(0, 255, 0)
	gfx.Vollrechteck(e.location.x-enemyWidth/2, e.location.y-enemyWidth, uint16(float64(e.health)/float64(e.maxHealth)*enemyWidth), 3)
}

func (e *Enemy) IsDead() bool {
	return e.dead
}

func (e *Enemy) SetDead(dead bool) {
	e.dead = dead
}

func (e *Enemy) IsLocationInHitBox(location Location) bool {
	distance := location.Distance(*e.location)
	return distance < enemyWidth
}

func (e *Enemy) DecreaseHealth(amount int) {
	e.health -= amount
	if e.health <= 0 {
		e.dead = true
	}
}
