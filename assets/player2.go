package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Player2ID uint = 2

var Player2Name string = "yeygym"

var Player2Position rl.Vector2

var Player2 rl.Texture2D
var Player2Rectangle rl.Rectangle

var Player2Rotation float32

var Player2Scale int32

func CreatePlayer2() {
	Player2 = rl.LoadTexture("sprites/p1_0.png")

	Player2Position = rl.Vector2{float32(Background.Width / 2), float32(Background.Height / 2)}

	Player2Rotation = 180
	Player2Scale = 3
}

func DrawPlayer2() {
	rl.DrawTexturePro(Player2, rl.Rectangle{0, 0, float32(Player2.Width), float32(Player2.Height)}, rl.Rectangle{float32(Player2Position.X), float32(Player2Position.Y), float32(Player2.Width * Player2Scale), float32(Player2.Height * Player2Scale)},
		rl.Vector2{float32(Player2.Width / 2 * Player2Scale), float32(Player2.Height / 2 * Player2Scale)}, Player2Rotation, rl.White)

	textSize := rl.MeasureText(Player2Name, 12)
	rl.DrawText(Player2Name, int32(Player2Position.X)-(textSize/2), int32(Player2Position.Y)+40, 12, rl.White)
	rl.DrawRectangle(int32(Player2Position.X)-(textSize/2)-7, int32(Player2Position.Y)+40-7, textSize+15, 15+12, rl.Color{0, 0, 0, 60})
}
