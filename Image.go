package main

import (
	"gfx"
	"log"
	"os"
	"strings"
)

type Image struct {
	id Identifier
}

func CreateImage(id Identifier) Image {
	//Vor.: Ein identifier, welcher de Pfad zu einem bpm-image hat, wird übergeben
	//Eff. Ein image Objekt wird zurückgegeben

	fileInfo, err := os.Stat(id.getPath())

	if err != nil {
		log.Fatal("Es kann kein Image Object erstellt werden, da der Pfad auf keine Image Dateo zeigt")
		return Image{}
	}
	if !strings.HasSuffix(fileInfo.Name(), ".bmp") {
		log.Fatal("Die hinterlegte Datei ist kein .bmp Image")
		return Image{}
	}

	return Image{
		id: id,
	}
}

func CreateImageName(name string) Image {
	return CreateImage(NewIdentifier(name))
}

func (i *Image) Render(x, y uint16) {
	gfx.LadeBild(x, y, i.id.path)
}
