package main

import (
	"fmt"
	"log"
)

const enemyWidth = 16

var enemyLevelTextures map[int]Image

func InitEnemyTextures() {
	enemyLevelTextures = make(map[int]Image)
	enemyLevelTextures[1] = CreateImageName("entity/enemy_1.bmp")
}

type Enemy struct {
	currentWaypoint    int
	ProcessedWayPoints []int
	location           *Location
	level              int
	pathfinder         Pathfinder
	world              World
}

func CreateEnemy(spawnLocation Location, world World) Enemy {
	fmt.Println("Create enemy")
	enemy := Enemy{
		location:   &spawnLocation,
		level:      1,
		world:      world,
		pathfinder: CreatePathFinder(world.level),
	}
	return enemy
}

func (e *Enemy) Update(deltaTime int64) {
	fmt.Println(len(e.ProcessedWayPoints))
	region := e.GetCurrentWayPoint().Region
	if region.ContainsPosition(e.location.x, e.location.y) {
		e.MarkAsProcessed(e.currentWaypoint)
		e.SetNextWayPoint(false)
	}

	path, found := e.pathfinder.FindPath(*e.location, region.GetRandomLocation())
	if !found {
		log.Fatalln("No path found to next waypoint!")
		return
	}

	nodeIndex := int(deltaTime)
	if nodeIndex >= len(path.nodes) {
		nodeIndex = len(path.nodes) - 1
	}
	e.location.x = path.nodes[nodeIndex].location.x
	e.location.y = path.nodes[nodeIndex].location.y
}

func (e *Enemy) SetNextWayPoint(isFirst bool) {
	_, i := e.getClosestWayPoint()
	e.currentWaypoint = i
}

func (e *Enemy) MarkAsProcessed(wayPoint int) {
	e.ProcessedWayPoints = append(e.ProcessedWayPoints, wayPoint)
}

func (e *Enemy) getClosestWayPoint() (LevelRegion, int) {
	var closest int = -1
	var closestDistance float64
	regions := e.world.level.GetRegionsOfType(RTYPE_WAYPOINT)
	for i, region := range regions {
		distance := e.location.Distance(region.Region.GetRandomLocation())
		if closest == -1 || (closestDistance > distance && !e.IsWayPointProcessed(closest)) {
			closest = i
			closestDistance = distance
		}
	}
	return regions[closest], closest
}

func (e *Enemy) IsWayPointProcessed(wayPointIndex int) bool {
	for _, point := range e.ProcessedWayPoints {
		if point == wayPointIndex {
			return true
		}
	}
	return false
}

func (e *Enemy) Render() {
	image := enemyLevelTextures[e.level]
	image.Render(e.location.x-enemyWidth/2, e.location.y-enemyWidth/2)
}

func (e *Enemy) GetCurrentWayPoint() LevelRegion {
	return e.world.level.GetRegionsOfType(RTYPE_WAYPOINT)[e.currentWaypoint]
}
