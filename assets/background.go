package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Background rl.Texture2D
var BackgroundRectangle rl.Rectangle

func CreateBackground() {
	Background = rl.LoadTexture("sprites/background_1600_900.png")

	BackgroundRectangle = rl.NewRectangle(0, 0, float32(Background.Width), float32(Background.Height))
}

func DrawBackground() {
	rl.DrawTextureRec(Background, BackgroundRectangle, rl.Vector2{0, 0}, rl.White)
}
