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
	Camera = rl.NewCamera2D(rl.Vector2{0, 0}, rl.Vector2{0, 0}, 0, 2)
}
