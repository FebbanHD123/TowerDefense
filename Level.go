package main

type Level struct {
	BackGroundImage Image
	Regions         []LevelRegion
}

func (l *Level) Render(x, y uint16) {
	l.BackGroundImage.Render(x, y)
}

func LoadLevel(path string) Level {
	//Vor.: Ein pfad zu einem File, welches die Level daten in sich trägt,
	//wird als string parameter übergeben
	//Eff.: Das level wird geladen und zurückgegeben
	return Level{}
}
