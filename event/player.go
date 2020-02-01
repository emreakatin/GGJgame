package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Speed float32 = 4
)

func PlayerController() {
	if rl.IsKeyDown(rl.KeyA) && float32(assets.PlayerPosition.X-Speed) > 0 {
		assets.PlayerPosition.X -= Speed
		assets.Camera.Offset.X -= Speed
	}

	if rl.IsKeyDown(rl.KeyD) && float32(assets.PlayerPosition.X-Speed) < float32(assets.Background.Width) {
		assets.PlayerPosition.X += Speed
		assets.Camera.Offset.X += Speed
	}

	if rl.IsKeyDown(rl.KeyW) && float32(assets.PlayerPosition.Y-Speed) > 0 {
		assets.PlayerPosition.Y -= Speed
		assets.Camera.Offset.Y -= Speed
	}

	if rl.IsKeyDown(rl.KeyS) && float32(assets.PlayerPosition.Y+Speed) < float32(assets.Background.Height) {
		assets.PlayerPosition.Y += Speed
		assets.Camera.Offset.Y += Speed
	}

	assets.UpdateCamera()
}
