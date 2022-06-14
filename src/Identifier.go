package main

import (
	"strings"
)

//Ziel: Einfache benutzung von assets/resources.

const assetsPath = "assets/"

type Identifier struct {
	Path string
}

func CreateIdentifier(path string) Identifier {
	//Eff.: Gibt einen neuen Identifier mit dem übergebenen Path zurück
	if !strings.HasPrefix(path, assetsPath) {
		path = assetsPath + path
	}
	return Identifier{
		Path: path,
	}
}

func (i *Identifier) getPath() string {
	//Eff.: Gibt den Path des Identifiers zurück
	return i.Path
}
