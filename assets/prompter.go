package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreatePrompter() {

}

func DrawPrompter(text string, fontSize int32, marginTop int32) {
	textSize := rl.MeasureText(text, fontSize)

	rl.DrawText(text, (Background.Width/2)-(textSize/2), (Background.Height/2)+marginTop, fontSize, rl.White)
	rl.DrawRectangle((Background.Width/2)-(textSize/2)-15, (Background.Height/2)+marginTop-15, textSize+30, 30+fontSize, rl.Color{0, 0, 0, 60})
}
