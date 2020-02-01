package assets

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TurretRadius = 50
	FireRadius   = 600
)

type Turret struct {
	ID uint

	Texture rl.Texture2D

	Rotation float32

	LockedID   int
	LockedType int // 1 Station, 2 Factory, 3 Turret

	OwnerID  uint
	Position rl.Vector2
	Health   uint

	ReloadFPS int
}

var Turrets []Turret
var LastTurretID = 0

func (turret Turret) DrawTurret() {
	rl.DrawTexturePro(turret.Texture, rl.Rectangle{0, 0, float32(turret.Texture.Width), float32(turret.Texture.Height)}, rl.Rectangle{turret.Position.X, turret.Position.Y, float32(turret.Texture.Width), float32(turret.Texture.Height)}, rl.Vector2{float32(turret.Texture.Width / 2), float32(turret.Texture.Height / 2)}, turret.Rotation, rl.White)
	rl.DrawCircle(int32(turret.Position.X+float32(0)/2), int32(turret.Position.Y+float32(0)/2), TurretRadius, rl.Color{uint8(turret.OwnerID+2) * uint8(40), uint8(turret.OwnerID+1) * uint8(40), uint8(turret.OwnerID+1) * uint8(40), 50})

	rl.DrawRectangleLinesEx(rl.Rectangle{turret.Position.X - float32(turret.Texture.Width)/2, turret.Position.Y - float32(turret.Texture.Height)/2, float32(turret.Texture.Width), float32(turret.Texture.Height)}, 3, rl.Red)

	// rl.DrawCircleLines(int32(turret.Position.X), int32(turret.Position.Y), FireRadius, rl.White)

	// rl.DrawCircle(int32(turret.Position.X)+turret.Texture.Width/2, int32(turret.Position.Y)+turret.Texture.Height/2, 5, rl.Red)

	rl.DrawLine(int32(turret.Position.X), int32(turret.Position.Y), 967, 332, rl.Red)

}
