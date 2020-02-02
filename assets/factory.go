package assets

import rl "github.com/gen2brain/raylib-go/raylib"

type Factory struct {
	ID uint

	Texture rl.Texture2D

	OwnerID   int
	Position  rl.Vector2
	Health    float32
	Rotation  float32
	ReloadFPS int
}

var Factories []Factory
var LastFactoryID = 0

func (factory Factory) DrawFactory() {
	rl.DrawTexturePro(factory.Texture, rl.Rectangle{0, 0, float32(factory.Texture.Width), float32(factory.Texture.Height)}, rl.Rectangle{factory.Position.X, factory.Position.Y, float32(factory.Texture.Width), float32(factory.Texture.Height)}, rl.Vector2{float32(factory.Texture.Width / 2), float32(factory.Texture.Height / 2)}, factory.Rotation, rl.White)
}
