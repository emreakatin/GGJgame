package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Camera rl.Camera2D

func GetCameraX() float32 {
	return -1 * Camera.Offset.X
}

func GetCameraY() float32 {
	return Camera.Offset.Y
}

func CreateCamera() {
	Camera = rl.NewCamera2D(rl.Vector2{float32(Background.Width) / 4, float32(Background.Height) / 4}, rl.Vector2{PlayerPosition.X, PlayerPosition.Y}, 0, 1.3)
}

func UpdateCamera() {
	Camera.Target.X = PlayerPosition.X
	Camera.Target.Y = PlayerPosition.Y
}
