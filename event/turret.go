package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var promptTicker = 0
var promptFlag = false

func TurretController() {

	// CREATE TURRET
	for _, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, assets.StationRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) && station.OwnerID == int(assets.PlayerID) {
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
		}
	}

	// PROMPTER FOR MOMEY
	if promptFlag {
		promptTicker++
		if promptTicker < 90 {
			assets.DrawPrompter("You do not have enough money", 23, 250)
		} else if promptTicker == 90 {
			promptFlag = false
			promptTicker = 0
		}
	}
}
