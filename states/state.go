package states

import (
	"fmt"

	"github.com/emreakatin/GGJgame/assets"
	"github.com/emreakatin/GGJgame/event"
	"github.com/emreakatin/GGJgame/scripts"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Load() {
	// BACKGROUND
	assets.CreateBackground()

	// CAMERA
	assets.CreateCamera()

	// PLAYER
	assets.CreatePlayer()

	// STATIONS
	assets.Stations = scripts.GenerateStations()
}

func Run() {
	rl.BeginMode2D(assets.Camera)

	// BACKGROUND
	assets.DrawBackground()

	// PLAYER
	assets.DrawPlayer()

	// STATIONS
	for _, station := range assets.Stations {
		station.DrawStation()
	}

	rl.EndMode2D()

	rl.DrawText(fmt.Sprintf("Camera: X: %f Y: %f", assets.Camera.Offset.X, assets.Camera.Offset.Y), 190, 200, 20, rl.Black)

	event.PlayerController()

}
