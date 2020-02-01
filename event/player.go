package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Speed float32 = 4
)

func PlayerController() {
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
