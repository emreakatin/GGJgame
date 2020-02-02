package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Player3ID uint = 1

var Player3Name string = "makatin"

var Player3Position rl.Vector2

var Player3 rl.Texture2D
var Player3Rectangle rl.Rectangle

var Player3Rotation float32

var Player3Scale int32

func CreatePlayer3() {
	Player3 = rl.LoadTexture("sprites/p3_1.png")

	Player3Position = rl.Vector2{float32(Background.Width / 2), float32(Background.Height / 2)}

	Player3Rotation = 180
	Player3Scale = 3
}

func DrawPlayer3() {
	rl.DrawTexturePro(Player3, rl.Rectangle{0, 0, float32(Player3.Width), float32(Player3.Height)}, rl.Rectangle{float32(Player3Position.X), float32(Player3Position.Y), float32(Player3.Width * Player3Scale), float32(Player3.Height * Player3Scale)},
		rl.Vector2{float32(Player3.Width / 2 * Player3Scale), float32(Player3.Height / 2 * Player3Scale)}, Player3Rotation, rl.White)

	textSize := rl.MeasureText(Player3Name, 12)
	rl.DrawText(Player3Name, int32(Player3Position.X)-(textSize/2), int32(Player3Position.Y)+40, 12, rl.White)
	rl.DrawRectangle(int32(Player3Position.X)-(textSize/2)-7, int32(Player3Position.Y)+40-7, textSize+15, 15+12, rl.Color{0, 0, 0, 60})
}
