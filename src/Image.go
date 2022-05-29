package main

import (
	"fmt"
	"gfx"
	"os"
	"strings"
)

type Image struct {
	Id Identifier
}

func CreateImage(id Identifier) Image {
	//Vor.: Ein identifier, welcher de Pfad zu einem bpm-image hat, wird übergeben
	//Eff. Ein image Objekt wird zurückgegeben

	fileInfo, err := os.Stat(id.getPath())

	if err != nil {
		fmt.Println("Es kann kein Image Object erstellt werden, da der Pfad auf keine Image Dateo zeigt")
		return Image{}
	}
	if !strings.HasSuffix(fileInfo.Name(), ".bmp") {
		fmt.Println("Die hinterlegte Datei ist kein .bmp Image")
		return Image{}
	}

	return Image{
		Id: id,
	}
}

func CreateImageName(name string) Image {
	return CreateImage(NewIdentifier(name))
}

func (i *Image) Render(x, y uint16) {
	gfx.LadeBild(x, y, i.Id.Path)
}
