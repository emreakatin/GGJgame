package assets

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	StationRadius = 150
)

type Station struct {
	ID uint

	Texture   rl.Texture2D
	Rectangle rl.Rectangle

	Position  rl.Vector2
	OwnerID   int
	Situation uint // 0: BROKEN, 1: REPAIRING, 2: REPAIRED
	Health    uint // OVER 100
}

var Stations []Station

func (station Station) CreateStation() {
	station.Texture = rl.LoadTexture("sprites/station.png")
}

func (station Station) DrawStation() {
	// station.Rectangle = rl.Rectangle{station.Position.X, station.Position.Y, float32(station.Texture.Width), float32(station.Texture.Height)}
	station.Rectangle = rl.NewRectangle(0, 0, float32(station.Texture.Width), float32(station.Texture.Height))

	// rl.DrawRectangleRec(station.Rectangle, rl.Red)
	rl.DrawTextureRec(station.Texture, station.Rectangle, station.Position, rl.Red)
	rl.DrawCircle(int32(station.Position.X+float32(station.Texture.Width)/2), int32(station.Position.Y+float32(station.Texture.Height)/2), StationRadius, rl.Color{uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+2) * uint8(40), 50})

	// rl.DrawCircle(int32(station.Position.X+float32(station.Texture.Width)/2), int32(station.Position.Y+float32(station.Texture.Height)/2), StationRadius, rl.Color{uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+2) * uint8(40), 50})

}
