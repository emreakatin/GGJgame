package engine

import (
	"github.com/emreakatin/GGJgame/states"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  int32 = 1600
	screenHeight int32 = 900

	screenTitle string = "Battle Royale"
)

func init() {

}

func Run() {
	rl.InitWindow(screenWidth, screenHeight, screenTitle)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		states.Run()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
