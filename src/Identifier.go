package main

import (
	"strings"
)

//Ziel: Einfache benutzung von assets/resources.

const assetsPath = "assets/"

type Identifier struct {
	Path string
}

func NewIdentifier(path string) Identifier {
	if !strings.HasPrefix(path, assetsPath) {
		path = assetsPath + path
	}
	return Identifier{
		Path: path,
	}
}

func (i *Identifier) getPath() string {
	return i.Path
}
