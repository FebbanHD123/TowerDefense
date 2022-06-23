package main

import (
	"gfx"
)

type ButtonWidget struct {
	title               string
	x, y, width, height uint16
	callback            func()
	activated           bool
	animationAlpha      int
}

func CreateButtonWidget(title string, x, y, width, height uint16, callback func()) ButtonWidget {
	//Eff.: Gibt ein neuen Button mit den übergebenen Attributen zurück
	return ButtonWidget{
		title:     title,
		x:         x,
		y:         y,
		width:     width,
		height:    height,
		callback:  callback,
		activated: true,
	}
}

func (b *ButtonWidget) Render(mouseX, mouseY uint16) {
	//Eff.: Rendert den Button
	b.updateAnimation(currentDeltaTime)
	if b.activated {
		gfx.Stiftfarbe(115+uint8(b.animationAlpha), 115+uint8(b.animationAlpha), 115+uint8(b.animationAlpha))
	} else {
		gfx.Stiftfarbe(100, 100, 100)
	}

	gfx.Vollrechteck(b.x-2, b.y-2, b.width+2*2, b.height+2*2)
	if b.activated {
		gfx.Stiftfarbe(100, 100, 100)
	} else {
		gfx.Stiftfarbe(60, 60, 60)
	}
	gfx.Vollrechteck(b.x, b.y, b.width, b.height)
	if b.activated {
		gfx.Stiftfarbe(255, 255, 255)
	} else {
		gfx.Stiftfarbe(150, 150, 150)
	}
	RenderCenteredText(b.title, b.x+b.width/2, b.y+b.height/2-FontHeight/2)
}

func (b ButtonWidget) MousePress(taste uint8, mouseX, mouseY uint16) {
	//Eff.: Wenn gedrückt wird, und die Maus über den Button ist wird das callback ausgeführt
	if taste == 1 && b.isMouseHover(mouseX, mouseY) && b.activated {
		b.callback()
	}

}

func (b *ButtonWidget) isMouseHover(mouseX, mouseY uint16) bool {
	//Eff.: Gibt zurück, ob die Maus über den Button ist
	rect := CreateRect(b.x, b.y, b.width, b.height)
	return rect.ContainsPosition(mouseX, mouseY)
}

func (b *ButtonWidget) SetActivated(activated bool) {
	//Eff.: Setzt den status des Buttons
	b.activated = activated
}

func (b *ButtonWidget) updateAnimation(deltaTime int64) {
	//Vor.: -
	//Eff.: Erhöht oder Sinkt die alpha variable des buttons,
	//welche den einfluss hat, dass der Farbton des Buttons heller (bei hohem alpha)
	//oder dunkler (bei niedrigem alpha) ist.
	var maxAnimation int = 50
	if b.isMouseHover(MouseX, MouseY) {
		b.animationAlpha += 3 * int(deltaTime)
		if b.animationAlpha > maxAnimation {
			b.animationAlpha = maxAnimation
		}
	} else if b.animationAlpha > 0 {
		b.animationAlpha -= 3 * int(deltaTime)
		if b.animationAlpha < 0 {
			b.animationAlpha = 0
		}
	}
}

func (b *ButtonWidget) SetTitle(title string) {
	//Eff.: Title des Buttons wird auf den übergebenen string gesetzt
	b.title = title
}
