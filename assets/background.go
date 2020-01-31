package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var background rl.Texture2D
var backgroundRectangle rl.Rectangle

func CreateBackground() {
	background = rl.LoadTexture("sprites/background_1600_900.png")

	backgroundRectangle = rl.NewRectangle(0, 0, float32(background.Width), float32(background.Height))
}

func DrawBackground() {
	rl.DrawTextureRec(background, backgroundRectangle, rl.Vector2{0, 0}, rl.White)
}
