package main

import "gfx"
import "time"

var xTower float64 = 500
var yTower float64 = 400
var xGegner float64 = 100
var yGegner float64 = 200
var speed float64 = 0.5

func main1() {
	gfx.Fenster(800, 600)
	for {
		gfx.UpdateAus()
		gfx.Stiftfarbe(0, 0, 0)
		gfx.Cls()
		gfx.Stiftfarbe(255, 2, 25)
		gfx.Vollkreis(uint16(xTower), uint16(yTower), 5)
		gfx.Vollkreis(uint16(xGegner), uint16(yGegner), 5)

		gk := yTower - yGegner
		ak := xGegner - xTower
		tan := gk / ak
		Update(tan)

		xGegner += 5
		yGegner += 0.5

		gfx.UpdateAn()
		time.Sleep(15 * time.Millisecond)
	}
}

func Update(anstieg float64) {
	dX := 5.0
	if xTower >= xGegner {
		dX = -dX
	}
	dY := anstieg * dX
	xTower += dX * speed
	yTower -= dY * speed

}
