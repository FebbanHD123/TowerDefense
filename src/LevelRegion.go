package main

import (
	"fmt"
)

const (
	RTYPE_PATH    int = 0
	RTYPE_DEFENSE     = 1
	RTYPE_GOAL        = 2
	RTYPE_SPAWN       = 3
)

type LevelRegion struct {
	Region Rect
	Type   int
}

func CreateLevelRegion(regionType int, region Rect) LevelRegion {
	//Vor.: region type wird übergeben: Siehe oben bei den konstanten
	//		Ein region in Form eines Rects wird übergeben
	//Eff.: Es wird eine Level-Region zurückgegeben, mit den übergeben attributen
	if regionType != RTYPE_PATH && regionType != RTYPE_DEFENSE && regionType != RTYPE_GOAL {
		fmt.Println("LevelRegion konnte nicht erstellt werden:", "Übergebener Region type existiert nicht")
	}
	return LevelRegion{
		Region: region,
		Type:   regionType,
	}
}
