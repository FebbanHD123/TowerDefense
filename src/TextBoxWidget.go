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
	animationAlpha      int
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
	t.updateAnimation(currentDeltaTime)
	gfx.Stiftfarbe(70+uint8(t.animationAlpha), 70+uint8(t.animationAlpha), 70+uint8(t.animationAlpha))
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
	t.cursor = len(t.text)
}

//Achtung: Englische Tastatur
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
			}
			return
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

		if IsInvalidTextCharacter(taste) {
			return
		}
		if tiefe == 1 { // wenn shift gedrückt
			if taste >= 97 { //buchstabe
				taste -= 97 - 65
			} else if taste >= 48 { //zahl
				taste -= 48 - 32
			}

		}

		var text string

		if t.cursor == 0 {
			text += string(taste)
		}
		for i := 0; i < len(t.text); i++ {
			text += string(t.text[i])
			if i+1 == t.cursor {
				text += string(taste)
			}
		}
		t.text = text
		t.cursor++
	}
}

func IsInvalidTextCharacter(c uint16) bool {
	invalidChars := []uint16{
		304, 301,
	}
	for _, char := range invalidChars {
		if char == c {
			return true
		}
	}
	return false
}

func (t *TextBox) isEmpty() bool {
	return len(t.text) == 0
}

func (t *TextBox) SetText(text string) {
	//Vor.: text wird übergeben
	//Eff.: Text der textbox wird auf den übergeben string gesetzt
	t.text = text
}

func (t *TextBox) updateAnimation(deltaTime int64) {
	//Vor.: -
	//Eff.: Erhöht oder Sinkt die alpha variable des buttons,
	//welche den einfluss hat, dass der Farbton des Buttons heller (bei hohem alpha)
	//oder dunkler (bei niedrigem alpha) ist.
	var maxAnimation int = 80
	if t.isMouseHover(MouseX, MouseY) || t.selected {
		t.animationAlpha += 4 * int(deltaTime)
		if t.animationAlpha > maxAnimation {
			t.animationAlpha = maxAnimation
		}
	} else if t.animationAlpha > 0 {
		t.animationAlpha -= 4 * int(deltaTime)
		if t.animationAlpha < 0 {
			t.animationAlpha = 0
		}
	}
}

func (t *TextBox) isMouseHover(mouseX, mouseY uint16) bool {
	rect := CreateRect(t.x, t.y, t.width, t.height)
	return rect.ContainsPosition(mouseX, mouseY)
}
