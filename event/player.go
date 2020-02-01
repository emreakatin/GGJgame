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
	if rl.IsKeyDown(rl.KeyA) && float32(assets.PlayerPosition.X-Speed) > 0 {
		assets.PlayerPosition.X -= Speed
		assets.Camera.Offset.X -= Speed
	}

	if rl.IsKeyDown(rl.KeyD) && float32(assets.PlayerPosition.X+float32(assets.Player.Width)+Speed) < float32(assets.Background.Width) {
		assets.PlayerPosition.X += Speed
		assets.Camera.Offset.X += Speed
	}

	if rl.IsKeyDown(rl.KeyW) && float32(assets.PlayerPosition.Y-Speed) > 0 {
		assets.PlayerPosition.Y -= Speed
		assets.Camera.Offset.Y -= Speed
	}

	if rl.IsKeyDown(rl.KeyS) && float32(assets.PlayerPosition.Y+float32(assets.Player.Height)+Speed) < float32(assets.Background.Height) {
		assets.PlayerPosition.Y += Speed
		assets.Camera.Offset.Y += Speed
	}

	assets.UpdateCamera()
}
