package event

import (
	"fmt"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var prompt2TickerMoney = 0
var prompt2FlagMoney = false

func Turret2Controller() {

	// CREATE TURRET
	for _, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, assets.StationRadius, rl.Rectangle{float32(assets.Player2Position.X), float32(assets.Player2Position.Y), float32(assets.Player2.Width * assets.Player2Scale), float32(assets.Player2.Height * assets.Player2Scale)}) && station.OwnerID == int(assets.Player2ID) {
			if rl.IsKeyPressed(rl.KeyRightControl) && assets.Player2Inventory.MechanicParts > assets.TurretCost {
				fmt.Println("lalalooo")
				assets.LastTurretID += 1
				assets.Player2Inventory.MechanicParts += -assets.TurretCost
				turret := assets.Turret{
					ID:       uint(assets.LastTurretID),
					OwnerID:  assets.Player2ID,
					Position: assets.Player2Position,
					Health:   100,
					Rotation: 0,
					Texture:  rl.LoadTexture("sprites/turret.png"),
					LockedID: -1,
					Ticker:   0,
					Thick:    0,
				}

				assets.Turrets = append(assets.Turrets, turret)
			}
		}
	}

	// TURRET COLLIDING PROMPT
	for _, turret := range assets.Turrets {

		if rl.CheckCollisionCircleRec(rl.Vector2{(turret.Position.X + float32(0)/2), (turret.Position.Y + float32(0)/2)}, assets.TurretRadius, rl.Rectangle{float32(assets.Player2Position.X), float32(assets.Player2Position.Y), float32(assets.Player2.Width * assets.Player2Scale), float32(assets.Player2.Height * assets.Player2Scale)}) {
			if !rl.IsKeyDown(rl.KeyRightShift) {
				if turret.OwnerID == assets.Player2ID {
					//assets.DrawPrompter("Hold press \"E\" to remove this turret", 23, 235)
				} else if turret.OwnerID != assets.Player2ID {
					//assets.DrawPrompter("Hold press \"E\" to destroy enemy turret", 23, 235)
				}
			}
		}

	}

	// PROMPTER FOR OUT OF SAFE ZONE AND MONEY
	if prompt2FlagMoney {
		prompt2TickerMoney++
		if prompt2TickerMoney < 90 {
			assets.DrawPrompter("You do not have enough mechanical part", 20, 250)
		} else if prompt2TickerMoney == 90 {
			prompt2FlagMoney = false
			prompt2TickerMoney = 0
		}
	}
}
