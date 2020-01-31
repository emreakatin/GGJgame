package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Camera rl.Camera2D

func CreateCamera() {
	Camera = rl.NewCamera2D(rl.Vector2{0, 0}, rl.Vector2{0, 0}, 0, 2)
}
