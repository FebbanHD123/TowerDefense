package main

import (
	"gfx"
)

type ButtonWidget struct {
	title               string
	x, y, width, height uint16
	callback            func()
	activated           bool
}

func CreateButtonWidget(title string, x, y, width, height uint16, callback func()) ButtonWidget {
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
	if !b.activated {
		gfx.Stiftfarbe(60, 60, 60)
	} else if b.isMouseHover(mouseX, mouseY) {
		gfx.Stiftfarbe(110, 110, 110)
	} else {
		gfx.Stiftfarbe(90, 90, 90)
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
	if taste == 1 && b.isMouseHover(mouseX, mouseY) {
		b.callback()
	}

}

func (b *ButtonWidget) isMouseHover(mouseX, mouseY uint16) bool {
	rect := CreateRect(b.x, b.y, b.width, b.height)
	return rect.ContainsPosition(mouseX, mouseY)
}

func (b *ButtonWidget) setActivated(activated bool) {
	b.activated = activated
}
