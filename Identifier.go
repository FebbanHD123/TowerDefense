package main

//Ziel: Einfache benutzung von assets/resources.

const assetsPath = "assets/"

type Identifier struct {
	path string
}

func NewIdentifier(name string) Identifier {
	path := assetsPath + name
	return Identifier{
		path: path,
	}
}

func (i *Identifier) getPath() string {
	return i.path
}
