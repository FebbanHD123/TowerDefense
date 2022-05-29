package main

import "bytes"

type Rect struct {
	X, Y, Width, Height uint16
}

func (r Rect) ContainsPosition(x, y uint16) bool {
	//Vor.: eine position wird in form zweier uint8 parameter (x und y in px) übergeben
	//Eff.: Es wird zurückgegeben, ob die Position in den Rect ist
	return x >= r.X && y >= r.Y && x < r.X+r.Width && y < r.Y+r.Height
}

func (r *Rect) WriteToBuffer(buffer *bytes.Buffer) {
	buffer.WriteRune(rune(r.X))
	buffer.WriteRune(rune(r.Y))
	buffer.WriteRune(rune(r.Width))
	buffer.WriteRune(rune(r.Height))
}

func ReadRectFromBuffer(buffer bytes.Buffer) Rect {
	x, _, _ := buffer.ReadRune()
	y, _, _ := buffer.ReadRune()
	width, _, _ := buffer.ReadRune()
	height, _, _ := buffer.ReadRune()

	return CreateRect(uint16(x), uint16(y), uint16(width), uint16(height))
}

func CreateRect(x, y, width, height uint16) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
