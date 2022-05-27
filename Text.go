package main

import "gfx"
import "github.com/golang/freetype/truetype"
import "os"
import "fmt"

//Ziel: Auslesen der ttf/font datei um die width von texten herrauszufinden
//      und so auch text mittig rendern können, da diese FUnktion wichtig ist
//      und von gfx nicht uterstützt wird.
//      Außerdem wird es benötigt, um eine y-verschiebung von fonts einzubauen.
//      Wenn z.b eine Font immer 3 pixel zu tief gerendert wird, muss dies ausgeglichen werden.

const (
	FontYVerschiebung = 5
	FontHeight = 20 - FontYVerschiebung
)

var fontIdentifier = NewIdentifier("fonts/Skranji-Regular.ttf")
var font *truetype.Font

func InitFont() {
	
	file, err := os.Open(fontIdentifier.getPath())
	if err != nil {
		panic(err)
	}
	
	//erstelle byte array (größe 1 mb)
	ttf := make([]byte, 1024 * 1024)
	//Fülle das array mit den daten des files
	size, err := file.Read(ttf)
	if err != nil {
		panic(err)
	}
	//entferne überschüssige array indices
	ttf = ttf[:size]
	
	gfx.SetzeFont(fontIdentifier.getPath(), FontHeight + FontYVerschiebung)
	f, err := truetype.Parse(ttf)
	font = f
	if err != nil {
		panic(err)
	}
	fmt.Println(font.HMetric(FontHeight, font.Index(rune("B"[0]))).AdvanceWidth)
}

func RenderText(text string, x, y uint16) {
	//Vor.: text und die position werden übergeben
	//Eff.: text wird an der poisiton gerendert 
	gfx.SchreibeFont(x, y, text)
}

func RenderCenteredText(text string, x, y uint16) {
	//Vor.: text und die position werden übergeben
	//Eff.: text wird mittig gerendert. Also nicht wie bei RenderText
	//an der x- und y-position sondern die x- und y-position ist in der mitte des textes
	width := GetTextWidth(text)
	RenderText(text, x - width / 2, y - FontYVerschiebung)
}

func GetTextWidth(text string) uint16 {
	//Vor.: text wird als string übergeben
	//Eff.: Horizontale pixelanzahl wird als uint16 zurückgegeben
	var width uint16
	for i := 0; i < len(text); i++ {
		char := rune(text[i])
		charWidth := uint16(font.HMetric(FontHeight, font.Index(char)).AdvanceWidth)
		width += charWidth
	}
	return width
}
 