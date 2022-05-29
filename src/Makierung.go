package main

import (
	"math"
)

type Makierung struct {
	X1, Y1, X2, Y2 uint16
	step           int
}

func CreateMakierung() Makierung {
	return Makierung{
		X1:   0,
		Y1:   0,
		X2:   0,
		Y2:   0,
		step: -1,
	}
}

func (m *Makierung) IsFinished() bool {
	return m.step >= 1
}

func (m *Makierung) AddPos(x, y uint16) {
	m.step++
	switch m.step {
	case 0:
		m.X1 = x
		m.Y1 = y
	case 1:
		m.X2 = x
		m.Y2 = y
	}
}

func (m *Makierung) ToRect() Rect {
	//Vor.: -
	//Eff.: Wandelt die makierung in ein Rect um,
	//also in eine Position (x, y) und in Breite und höhe

	var x, y, width, height uint16

	//math.Abs: Wenn der wert im negativen bereich ist, wird das vorzeichen positiv
	//Es ist also egal ob man x - y oder y - x rechnet
	width = uint16(math.Abs(float64(int(m.X2) - int(m.X1))))
	height = uint16(math.Abs(float64(int(m.Y2) - int(m.Y1))))

	if m.X1 < m.X2 && m.Y1 < m.Y2 {
		x = m.X1
		y = m.Y1
	} else if m.X1 < m.X2 {
		x = m.X1
		y = m.Y2
	} else if m.X1 > m.X2 && m.Y1 < m.Y2 {
		x = m.X2
		y = m.Y1
	} else {
		x = m.X2
		y = m.Y2
	}

	//return rect
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}

}

func (m *Makierung) Reset() {
	m.step = -1
	m.X1 = 0
	m.Y1 = 0
	m.X2 = 0
	m.Y2 = 0
}
