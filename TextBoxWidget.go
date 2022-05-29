package main

import (
	"gfx"
	"time"
)

type TextBox struct {
	text                string
	cursor              int
	x, y, width, height uint16
	showCursor          bool
	lastCursorUpdate    time.Time
	selected            bool
}

func CreateTextBox(x, y, width, height uint16) TextBox {
	return TextBox{
		text:             "",
		cursor:           0,
		x:                x,
		y:                y,
		width:            width,
		height:           height,
		showCursor:       true,
		lastCursorUpdate: time.Now(),
	}
}

func (t *TextBox) Update() {
	if time.Now().Sub(t.lastCursorUpdate).Milliseconds() > 500 {
		t.lastCursorUpdate = time.Now()
		t.showCursor = !t.showCursor
	}
}

func (t *TextBox) Render() {
	gfx.Stiftfarbe(70, 70, 70)
	gfx.Vollrechteck(t.x-2, t.y-2, t.width+2*2, t.height+2*2)
	gfx.Stiftfarbe(10, 10, 10)
	gfx.Vollrechteck(t.x, t.y, t.width, t.height)
	text := t.text
	textWidth := GetTextWidth(text + "_")
	var startIndex int
	if textWidth > t.width {
		var width uint16 = 15
		for i := t.cursor - 1; i > 0; i-- {
			if width+GetTextWidth(string(t.text[i])) >= t.width {
				startIndex = i
				break
			}
			width += GetTextWidth(string(t.text[i]))
		}
	}
	text = ""
	if t.cursor == 0 && t.selected && t.showCursor {
		text += "|"
	}
	var width uint16
	for i := startIndex; i < len(t.text); i++ {
		if width+GetTextWidth(string(t.text[i])) < t.width+10 {
			width += GetTextWidth(string(t.text[i]))
			text += string(t.text[i])
		}
		if i+1 == t.cursor && t.showCursor && t.selected {
			text += "|"
		}
	}
	gfx.Stiftfarbe(250, 250, 250)
	RenderText(text, t.x+5, t.y+t.height/2-FontHeight/2)
}

func (t *TextBox) mousePress(taste uint8, mouseX, mouseY uint16) {
	rect := CreateRect(t.x, t.y, t.width, t.height)
	if taste != 1 {
		return
	}
	t.selected = rect.ContainsPosition(mouseX, mouseY)
	t.showCursor = true
	t.lastCursorUpdate = time.Now()
}

func (t *TextBox) keyPressed(taste uint16, gedrueckt bool, tiefe uint16) {
	if !t.selected {
		return
	}
	if gedrueckt {

		if taste == KEY_DELETE {
			if len(t.text) > 0 {
				var text string
				for i := 0; i < len(t.text); i++ {
					if i+1 != t.cursor {
						text += string(t.text[i])

					}
				}
				t.text = text
				if t.cursor > 0 {
					t.cursor--
				}
				return
			}
		}
		if taste == KEY_ARROR_RIGHT {
			if t.cursor != len(t.text) {
				t.cursor++
			}
			return
		}
		if taste == KEY_ARROR_LEFT {
			if t.cursor > 0 {
				t.cursor--
			}
			return
		}

		t.text += string(taste)
		t.cursor++
	}
}
