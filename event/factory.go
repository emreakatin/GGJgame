package event

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var promptTicker = 0
var promptFlag = false
var MPFlag = false
var MPTicker = 0

func FactoryController() {
	EarningMechanicalParts()
	// CREATE FACTORY
	for _, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, assets.StationRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) && station.OwnerID == int(assets.PlayerID) {
			if rl.IsKeyPressed(rl.KeyR) && assets.PlayerInventory.MechanicParts > assets.FactoryCost {
				assets.LastFactoryID += 1
				assets.PlayerInventory.MechanicParts += -assets.FactoryCost
				factory := assets.Factory{
					ID:       uint(assets.LastFactoryID),
					OwnerID:  assets.PlayerID,
					Position: assets.PlayerPosition,
					Health:   0,
					Rotation: 0,
					Texture:  rl.LoadTexture("sprites/factory.png"),
				}

				assets.Factories = append(assets.Factories, factory)
			} else if rl.IsKeyPressed(rl.KeyF) && assets.PlayerInventory.MechanicParts < assets.FactoryCost {
				promptFlag = true
			}
		}
	}
	// PROMPTER FOR OUT OF SAFE ZONE AND MONEY
	if promptFlag {
		promptTicker++
		if promptTicker < 90 {
			assets.DrawPrompter("You do not have enough mechanical part", 20, 250)
		} else if promptTicker == 90 {
			promptFlag = false
			promptTicker = 0
		}
	}
}

func EarningMechanicalParts() {
	for _, factory := range assets.Factories {
		if factory.OwnerID == uint(assets.PlayerID) {
			MPTicker++
			if MPTicker >= 90 {
				assets.PlayerInventory.MechanicParts++
				MPTicker = 0
			}
		}
	}
}
