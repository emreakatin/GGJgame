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
	OwnerID   int     // -1: Unowned
	Situation uint    // 0: BROKEN, 1: REPAIRING, 2: REPAIRED
	Health    float32 // OVER 100
}

var Stations []Station

func (station Station) CreateStation() {
	station.Texture = rl.LoadTexture("sprites/station.png")
}

func (station Station) DrawStation() {
	station.Rectangle = rl.NewRectangle(0, 0, float32(station.Texture.Width), float32(station.Texture.Height))

	// rl.DrawRectangleRec(station.Rectangle, rl.Red)
	rl.DrawTextureRec(station.Texture, station.Rectangle, station.Position, rl.White)
	if station.OwnerID != -1 {
		rl.DrawCircle(int32(station.Position.X+float32(station.Texture.Width)/2), int32(station.Position.Y+float32(station.Texture.Height)/2), StationRadius, rl.Color{uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+1) * uint8(40), uint8(station.OwnerID+1) * uint8(40), 50})
	}
	// rl.DrawCircle(int32(station.Position.X+float32(station.Texture.Width)/2), int32(station.Position.Y+float32(station.Texture.Height)/2), StationRadius, rl.Color{uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+2) * uint8(40), uint8(station.OwnerID+2) * uint8(40), 50})

}
func (station Station) UpdateStation() {
	if station.Health <= 25 {
		station.Texture = rl.LoadTexture("sprites/station0.png")
	} else if station.Health <= 50 && station.Health > 25 {
		station.Texture = rl.LoadTexture("sprites/station1.png")
	} else if station.Health <= 75 && station.Health > 50 {
		station.Texture = rl.LoadTexture("sprites/station2.png")
	} else if station.Health < 100 && station.Health > 75 {
		station.Texture = rl.LoadTexture("sprites/station3.png")
	} else if station.Health >= 100 {
		station.Texture = rl.LoadTexture("sprites/station4.png")
	}
	station.DrawStation()
}
