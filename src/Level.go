package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Infos:
//Background-Image-size: 800px x 720px

var Levels []Level

type Level struct {
	BackGroundImage Image         `json:"background"`
	Regions         []LevelRegion `json:"regions"`
	Name            string        `json:"name"`
}

func (l *Level) Render(x, y uint16) {
	//Eff.: Rendert das level
	l.BackGroundImage.Render(x, y)
}

func (l *Level) AddRegion(region LevelRegion) {
	l.Regions = append(l.Regions, region)
}

func LoadAllLevels() {
	//Vor.: -
	//Eff.: Alle level werrden geladen

	fmt.Println("Loading levels...")
	levelsDir := GameFilePath + "/levels"
	if !FileExistsSafe(levelsDir) {
		err := os.Mkdir(levelsDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	files, err := os.ReadDir(levelsDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		level, err := LoadLevel(file.Name())
		if err == nil {
			Levels = append(Levels, level)
		} else {
			fmt.Println("Cant load level from file:", file.Name())
		}
	}
	fmt.Println("Loaded", len(Levels), "levels")
}

func LoadLevel(fileName string) (Level, error) {
	//Vor.: Ein pfad zu einem File, welches die Level daten in Form von json, enthält,
	//wird als string parameter übergeben
	//Eff.: Das level wird geladen
	var level Level

	filePath := GameFilePath + "/levels/" + fileName
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return level, err
	}

	err = json.Unmarshal(data, &level)
	if err != nil {
		return level, err
	}

	return level, nil
}

func (l Level) Save() {
	//Vor.: -
	//Eff.: Speichert das Level in form von json in einem file

	data, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	levelsDir := GameFilePath + "/levels"
	if !FileExistsSafe(levelsDir) {
		err := os.Mkdir(levelsDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	filePath := levelsDir + "/" + strings.ToLower(l.Name) + ".level"
	if FileExistsSafe(filePath) {
		err = os.Remove(filePath)
		if err != nil {
			panic(err)
		}
	}
	err = ioutil.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Levels = nil
	Levels = make([]Level, 0)
	LoadAllLevels()
}

func (l *Level) DeleteOldFileIfExists() {
	filePath := GameFilePath + "/levels/" + strings.ToLower(l.Name) + ".level"
	if FileExistsSafe(filePath) {
		err := os.Remove(filePath)
		if err != nil {
			panic(err)
		}
	}
}

func (l *Level) GetRegions() []LevelRegion {
	//Vor.: -
	//Eff.: Gibt alle LevelRegionen des zurpck
	return l.Regions
}

func (l *Level) HasAllRequiredRegions() bool {
	//Vor.: -
	//Eff.: Gibt zurück, ob das level alle regionen hat, die es braucht
	var path, goal, defense, spawn, waypoint bool
	for _, region := range l.GetRegions() {
		switch region.Type {
		case RTYPE_PATH:
			path = true
		case RTYPE_GOAL:
			goal = true
		case RTYPE_DEFENSE:
			defense = true
		case RTYPE_SPAWN:
			spawn = true
		case RTYPE_WAYPOINT:
			waypoint = true
		}
	}
	return path && goal && defense && spawn && waypoint
}

func (l *Level) GetRegionOfType(regionType int) LevelRegion {
	for _, region := range l.Regions {
		if region.Type == regionType {
			return region
		}
	}
	panic("Region with this type does not exists!")
}

func (l *Level) GetRegionsOfType(regionType int) []LevelRegion {
	var regions []LevelRegion
	for _, region := range l.Regions {
		if region.Type == regionType {
			regions = append(regions, region)
		}
	}
	return regions
}

func (l *Level) GetRandomSpawnLocation() Location {
	spawnRegion := l.GetRegionOfType(RTYPE_SPAWN)
	return spawnRegion.Region.GetRandomLocation()
}

func (l *Level) GetRegionAtLocation(location Location) LevelRegion {
	for _, region := range l.Regions {
		if region.Region.ContainsPosition(location.x, location.y) {
			return region
		}
	}
	return LevelRegion{Type: RTYPE_NOTHING}
}

func (l *Level) GetGoalLocation() Location {
	goalRegion := l.GetRegionOfType(RTYPE_GOAL)
	return goalRegion.Region.GetRandomLocation()
}
