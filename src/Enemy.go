package main

import (
	"fmt"
	"gfx"
)

const enemySize = 30

var enemyLevelTextures map[int]Image

func InitEnemyTextures() {
	enemyLevelTextures = make(map[int]Image)
	enemyLevelTextures[1] = CreateImageName("entity/enemy/1.bmp")
	enemyLevelTextures[2] = CreateImageName("entity/enemy/2.bmp")
	enemyLevelTextures[3] = CreateImageName("entity/enemy/3.bmp")
	enemyLevelTextures[4] = CreateImageName("entity/enemy/4.bmp")
	enemyLevelTextures[5] = CreateImageName("entity/enemy/5.bmp")
	enemyLevelTextures[6] = CreateImageName("entity/enemy/6.bmp")
}

type Enemy struct {
	location   *Location
	level      int
	pathfinder Pathfinder
	world      *World
	path       []Location
	dead       bool
	health     int
	maxHealth  int
	speed      float64
}

func CreateEnemy(spawnLocation Location, world *World, level int, speed float64) Enemy {
	//Eff.: Gibt ein Objekt der Klasse Enemy mit den übergebenen Werten zurück
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
	//Eff.: Findet den kürzesten Weg zum ziel und Speichert die Weg-Locations ab
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
	//Eff.: Bewegt das Objekt über den Weg
	nodeIndex := int(float64(deltaTime) * e.speed)
	if nodeIndex >= len(e.path) {
		nodeIndex = len(e.path) - 1
	}
	e.location.x = e.path[nodeIndex].x
	e.location.y = e.path[nodeIndex].y

	e.path = e.path[nodeIndex:]

}

func (e *Enemy) getClosestWayPoint(currentLocation Location, processedWayPoint []LevelRegion) LevelRegion {
	//Eff.: Gibt den am wenigstens entfernten Waypoint, von der übergebenen Location, des Levels zurück
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
	//Eff.: Gibt zurück, ob ein Waypoint bereits vom Spieler abgelaufen wurde
	for _, point := range processedWayPoints {
		if point.equals(wayPoint) {
			return true
		}
	}
	return false
}

func (e *Enemy) Render() {
	//Eff.: Rendert das Enemy
	var image Image
	if e.level > len(enemyLevelTextures) {
		image = enemyLevelTextures[len(enemyLevelTextures)]
	} else {
		image = enemyLevelTextures[e.level]
	}
	image.Render(e.location.x-enemySize/2, e.location.y-enemySize/2)
	if e.health < e.maxHealth {
		e.renderHealthBar()
	}
}

func (e *Enemy) renderHealthBar() {
	//Eff.: Render die Lebensanzeige über dem Enemy
	gfx.Stiftfarbe(100, 100, 100)
	gfx.Vollrechteck(e.location.x-enemySize/2, e.location.y-enemySize, enemySize, 3)
	gfx.Stiftfarbe(0, 255, 0)
	gfx.Vollrechteck(e.location.x-enemySize/2, e.location.y-enemySize, uint16(float64(e.health)/float64(e.maxHealth)*enemySize), 3)
}

func (e *Enemy) IsDead() bool {
	//Eff.: Gibt zurück, ob der Enemy Tod ist
	return e.dead
}

func (e *Enemy) SetDead(dead bool) {
	//Eff.: Setzt ob der Spieler Tod ist
	e.dead = dead
}

func (e *Enemy) IsLocationInHitBox(location Location) bool {
	//Eff.: Gibt zurück, ob eine Location in der Hitbox des Enemys ist.
	distance := location.Distance(*e.location)
	return distance < enemySize
}

func (e *Enemy) DecreaseHealth(amount int) {
	//Eff.: Entfernt eine übergebene anzahl an Leben
	//		Wenn das leben unter 0 ist wird es auf 0 gesetzt
	e.health -= amount
	if e.health <= 0 {
		e.dead = true
	}
}
