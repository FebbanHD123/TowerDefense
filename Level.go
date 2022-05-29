package main

//Infos:
//Background-Image-size: 800px x 720px

type Level struct {
	BackGroundImage Image
	regions         []LevelRegion
}

func (l *Level) Render(x, y uint16) {
	l.BackGroundImage.Render(x, y)
}

func (l *Level) AddRegion(region LevelRegion) {
	l.regions = append(l.regions, region)
}

func LoadLevel(path string) Level {
	//Vor.: Ein pfad zu einem File, welches die Level daten in sich trägt,
	//wird als string parameter übergeben
	//Eff.: Das level wird geladen und zurückgegeben
	return Level{}
}

func (l *Level) GetRegions() []LevelRegion {
	return l.regions
}
