package states

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Load() {
	// CAMERA
	assets.CreateCamera()

	// BACKGROUND
	assets.CreateBackground()
}

func Run() {
	rl.BeginMode2D(assets.Camera)

	// BACKGROUND
	assets.DrawBackground()

	rl.EndMode2D()

}
