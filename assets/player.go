package assets

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var PlayerPosition rl.Vector2

var Player rl.Texture2D
var PlayerRectangle rl.Rectangle

var PlayerRotation float32

var PlayerScale int32

func CreatePlayer() {
	Player = rl.LoadTexture("sprites/player.png")

	PlayerPosition = rl.Vector2{float32(Background.Width / 2), float32(Background.Height / 2)}

	PlayerRotation = 90
	PlayerScale = 3
}

func DrawPlayer() {
	rl.DrawTexturePro(Player, rl.Rectangle{0, 0, float32(Player.Width), float32(Player.Height)}, rl.Rectangle{float32(PlayerPosition.X), float32(PlayerPosition.Y), float32(Player.Width * PlayerScale), float32(Player.Height * PlayerScale)},
		rl.Vector2{float32(Player.Width / 2 * PlayerScale), float32(Player.Height / 2 * PlayerScale)}, PlayerRotation, rl.White)

	fmt.Println("---", rl.Vector2{float32(Player.Width / 2), float32(Player.Height / 2)})

}
