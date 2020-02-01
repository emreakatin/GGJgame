package assets

import rl "github.com/gen2brain/raylib-go/raylib"

type Turret struct {
	ID uint

	Texture rl.Texture2D

	Rotation float32

	OwnerID  uint
	Position rl.Vector2
	Health   uint

	ReloadFPS int
}

var Turrets []Turret
var LastTurretID = 0

func (turret Turret) DrawTurret() {
	rl.DrawTexturePro(turret.Texture, rl.Rectangle{0, 0, float32(turret.Texture.Width), float32(turret.Texture.Height)}, rl.Rectangle{turret.Position.X, turret.Position.Y, float32(turret.Texture.Width), float32(turret.Texture.Height)}, rl.Vector2{float32(turret.Texture.Width / 2), float32(turret.Texture.Height / 2)}, turret.Rotation, rl.White)
}
