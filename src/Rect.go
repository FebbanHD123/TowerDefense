package main

import "math/rand"

type Rect struct {
	X, Y, Width, Height uint16
}

func (r Rect) ContainsPosition(x, y uint16) bool {
	//Vor.: eine position wird in form zweier uint8 parameter (x und y in px) übergeben
	//Eff.: Es wird zurückgegeben, ob die Position in den Rect ist
	return x >= r.X && y >= r.Y && x < r.X+r.Width && y < r.Y+r.Height
}

func (r *Rect) GetRandomLocation() Location {
	x := r.X + uint16(rand.Intn(int(r.Width+1)))
	y := r.Y + uint16(rand.Intn(int(r.Height+1)))
	return CreateLocation(x, y)
}

func CreateRect(x, y, width, height uint16) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r *Rect) equals(other Rect) bool {
	return r.X == other.X && r.Y == other.Y && r.Width == other.Width && r.Height == other.Height
}
