package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var PlayerName string = "nukeci"

var PlayerPosition rl.Vector2

var Player rl.Texture2D
var PlayerRectangle rl.Rectangle

var PlayerRotation float32

var PlayerScale int32

func CreatePlayer() {
	Player = rl.LoadTexture("sprites/player.png")

	PlayerPosition = rl.Vector2{float32(Background.Width / 2), float32(Background.Height / 2)}

	PlayerRotation = 180
	PlayerScale = 3
}

func DrawPlayer() {
	rl.DrawTexturePro(Player, rl.Rectangle{0, 0, float32(Player.Width), float32(Player.Height)}, rl.Rectangle{float32(PlayerPosition.X), float32(PlayerPosition.Y), float32(Player.Width * PlayerScale), float32(Player.Height * PlayerScale)},
		rl.Vector2{float32(Player.Width / 2 * PlayerScale), float32(Player.Height / 2 * PlayerScale)}, PlayerRotation, rl.White)

	textSize := rl.MeasureText(PlayerName, 12)
	rl.DrawText(PlayerName, int32(PlayerPosition.X)-(textSize/2), int32(PlayerPosition.Y)+40, 12, rl.White)
	rl.DrawRectangle(int32(PlayerPosition.X)-(textSize/2)-7, int32(PlayerPosition.Y)+40-7, textSize+15, 15+12, rl.Color{0, 0, 0, 60})
}
