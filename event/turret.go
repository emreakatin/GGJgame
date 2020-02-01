package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var promptTicker = 0
var promptFlag = false

func TurretController() {
	// CREATE TURRET

	if rl.IsKeyPressed(rl.KeyT) && assets.PlayerInventory.MechanicParts > assets.TurretCost {
		assets.LastTurretID += 1
		assets.PlayerInventory.MechanicParts += -assets.TurretCost
		turret := assets.Turret{
			ID:       uint(assets.LastTurretID),
			OwnerID:  assets.PlayerID,
			Position: assets.PlayerPosition,
			Health:   0,
			Rotation: assets.PlayerRotation,
			Texture:  rl.LoadTexture("sprites/turret.png"),
		}

		assets.Turrets = append(assets.Turrets, turret)
	} else if rl.IsKeyPressed(rl.KeyT) && assets.PlayerInventory.MechanicParts < assets.TurretCost {
		promptFlag = true
	}

	if promptFlag {
		promptTicker++
		if promptTicker < 180 {
			assets.DrawPrompter("You do not have enough money", 23, 250)
		} else if promptTicker == 180 {
			promptFlag = false
			promptTicker = 0
		}
	}
}
