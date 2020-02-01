package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TurretController() {
	// CREATE TURRET

	if rl.IsKeyPressed(rl.KeyT) && assets.PlayerInventory.MechanicParts > assets.FactoryCost {
		assets.LastTurretID += 1

		turret := assets.Turret{
			ID:       uint(assets.LastTurretID),
			OwnerID:  assets.PlayerID,
			Position: assets.PlayerPosition,
			Health:   0,
			Rotation: assets.PlayerRotation,
			Texture:  rl.LoadTexture("sprites/turret.png"),
		}

		assets.Turrets = append(assets.Turrets, turret)
	}
}
