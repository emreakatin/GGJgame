package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Speed float32 = 4
	// cameraPositionX float32
	// cameraPositionY float32
)

func PlayerController() {
	// cameraPositionX = assets.GetCameraX()
	// cameraPositionY = assets.GetCameraY()
	// if rl.IsKeyDown(rl.KeyLeft) {
	// 	PlayerMoveLeft()
	// }
	// if rl.IsKeyDown(rl.KeyUp) {
	// 	PlayerMoveUp()
	// }
	// if rl.IsKeyDown(rl.KeyDown) {
	// 	PlayerMoveDown()
	// }
	// if rl.IsKeyDown(rl.KeyRight) {
	// 	PlayerMoveRight()
	// }

	if rl.IsKeyDown(rl.KeyLeft) {
		assets.PlayerPosition.X -= Speed
		assets.Camera.Offset.X += Speed
	}

	if rl.IsKeyDown(rl.KeyRight) {
		assets.PlayerPosition.X += Speed
		assets.Camera.Offset.X -= Speed
	}

	if rl.IsKeyDown(rl.KeyUp) {
		assets.PlayerPosition.Y -= Speed
		assets.Camera.Offset.Y += Speed
	}

	if rl.IsKeyDown(rl.KeyDown) {
		assets.PlayerPosition.Y += Speed
		assets.Camera.Offset.Y -= Speed
	}

	assets.UpdateCamera()
}

// func PlayerMoveUp() {
// 	if cameraPositionY+Speed < 0 {
// 		assets.Camera.Offset.Y += Speed
// 	} else {
// 		assets.Camera.Offset.Y = 0
// 	}
// }

// func PlayerMoveDown() {
// 	if (cameraPositionY-Speed)*-1 < float32(assets.Background.Height) {
// 		assets.Camera.Offset.Y -= Speed
// 	} else {
// 		assets.Camera.Offset.Y = float32(-assets.Background.Height)
// 	}
// }

// func PlayerMoveRight() {
// 	if cameraPositionX+Speed < float32(assets.Background.Width) {
// 		assets.Camera.Offset.X -= Speed
// 	} else {
// 		assets.Camera.Offset.X = float32(-assets.Background.Width)
// 	}
// }

// func PlayerMoveLeft() {
// 	if cameraPositionX-Speed > 0 {
// 		assets.Camera.Offset.X += Speed
// 	} else {
// 		assets.Camera.Offset.X = 0
// 	}
// }

// func PlayerMoveUp() {
// 	if cameraPositionY+Speed < 0 {
// 		assets.Camera.Offset.Y += Speed
// 	} else {
// 		assets.Camera.Offset.Y = 0
// 	}
// }

// func PlayerMoveDown() {
// 	if (cameraPositionY-Speed)*-1 < float32(assets.Background.Height) {
// 		assets.Camera.Offset.Y -= Speed
// 	} else {
// 		assets.Camera.Offset.Y = float32(-assets.Background.Height)
// 	}
// }

// func PlayerMoveRight() {
// 	if cameraPositionX+Speed < float32(assets.Background.Width) {
// 		assets.Camera.Offset.X -= Speed
// 	} else {
// 		assets.Camera.Offset.X = float32(-assets.Background.Width)
// 	}
// }

// func PlayerMoveLeft() {
// 	if cameraPositionX-Speed > 0 {
// 		assets.Camera.Offset.X += Speed
// 	} else {
// 		assets.Camera.Offset.X = 0
// 	}
// }
