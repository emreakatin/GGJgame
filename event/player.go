package event

import (
	"fmt"
	"math"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Speed         float32 = 4
	mousePosition rl.Vector2
)

func PlayerController() {
	// ROTATION
	mousePosition = rl.GetMousePosition()

	// if mouse in screen
	if !(mousePosition.X < 0) && !(mousePosition.Y < 0) {
		radian := math.Atan2(float64(assets.PlayerPosition.Y-mousePosition.Y), float64(mousePosition.X-assets.PlayerPosition.X))
		degree := radian * 180 / math.Pi

		if degree > 0 && degree <= 90 {
			degree = 90 - degree
		} else if degree > 90 && degree <= 180 {
			degree = 270 - degree
		} else if degree > 180 && degree <= 270 {
			degree = 270 - degree
		} else if degree > 270 && degree <= 360 {
			degree = 360 - degree
		}

		assets.PlayerRotation = float32(degree)

		fmt.Println(mousePosition)
		fmt.Println(assets.PlayerPosition)

		fmt.Println(assets.PlayerRotation)
	}

	// MOVEMENT
	if rl.IsKeyDown(rl.KeyLeft) {
		assets.PlayerPosition.X -= Speed
		assets.Camera.Offset.X -= Speed
	}

	if rl.IsKeyDown(rl.KeyRight) {
		assets.PlayerPosition.X += Speed
		assets.Camera.Offset.X += Speed
	}

	if rl.IsKeyDown(rl.KeyUp) {
		assets.PlayerPosition.Y -= Speed
		assets.Camera.Offset.Y -= Speed
	}

	if rl.IsKeyDown(rl.KeyDown) {
		assets.PlayerPosition.Y += Speed
		assets.Camera.Offset.Y += Speed
	}

	assets.UpdateCamera()
}
