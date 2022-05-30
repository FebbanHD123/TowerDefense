package main

import "math"

type Location struct {
	x, y uint16
}

func CreateLocation(x, y uint16) Location {
	return Location{
		x: x,
		y: y,
	}
}

func (l *Location) Distance(other Location) float64 {
	//Eff.: Die distanz zwischen den beiden Locations wird zurückgegeben
	return math.Sqrt(math.Pow(float64(other.x)-float64(l.x), 2) + math.Pow(float64(other.y)-float64(l.y), 2))
}
