package main

type Rect struct {
	X, Y, Width, Height uint16
}

func (r Rect) ContainsPosition(x, y uint16) bool {
	//Vor.: eine position wird in form zweier uint8 parameter (x und y in px) übergeben
	//Eff.: Es wird zurückgegeben, ob die Position in den Rect ist
	return x >= r.X && y >= r.Y && x < r.X+r.Width && y < r.Y+r.Width
}

func CreateRect(x, y, width, height uint16) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
