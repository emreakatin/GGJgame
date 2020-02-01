package assets

import rl "github.com/gen2brain/raylib-go/raylib"

type Turret struct {
	ID uint

	Texture   rl.Texture2D
	Rectangle rl.Rectangle

	Rotation float32

	OwnerID  uint
	Position rl.Vector2
	Health   uint

	ReloadFPS int
}

var Turrets []Turret

func (turret Turret) LoadTurret() {
	turret.Texture = rl.LoadTexture("sprites/turret.png")

	turret.Rotation = 0
}

func (turret Turret) DrawTurret() {
	rl.DrawTexturePro(turret.Texture, rl.Rectangle{0, 0, float32(turret.Texture.Width), float32(turret.Texture.Height)}, turret.Rectangle, rl.Vector2{float32(turret.Texture.Width / 2), float32(turret.Texture.Height / 2)}, turrent.Rotation, rl.White)
}
