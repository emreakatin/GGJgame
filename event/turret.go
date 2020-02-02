package event

import (
	"fmt"
	"strconv"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TurretDamage = 10
)

var promptTickerMoney = 0
var promptFlagMoney = false

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
					Health:   100,
					Rotation: 0,
					Texture:  rl.LoadTexture("sprites/turret.png"),
					LockedID: -1,
					Thick:    0,
					Ticker:   0,
				}

				assets.Turrets = append(assets.Turrets, turret)
			} else if rl.IsKeyPressed(rl.KeyT) && assets.PlayerInventory.MechanicParts < assets.TurretCost {
				promptFlagMoney = true
			} else if rl.IsKeyDown(rl.KeyE) {
				// TURRET COLLIDING
				for index, turret := range assets.Turrets {

					if rl.CheckCollisionCircleRec(rl.Vector2{(turret.Position.X + float32(0)/2), (turret.Position.Y + float32(0)/2)}, assets.TurretRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) {
						if turret.OwnerID == assets.PlayerID {
							if turret.Health > 0 {
								assets.Turrets[index].Health -= 1
								assets.DrawPrompter("Removing turret. %"+strconv.Itoa(100-int(turret.Health)), 23, 235)
							} else if turret.Health <= 0 {
								assets.Turrets = append(assets.Turrets[:index], assets.Turrets[index+1:]...)
								assets.PlayerInventory.MechanicParts += assets.TurretCost * 2 / 3
							}

						} else if turret.OwnerID != assets.PlayerID {
							if turret.Health > 0 && assets.PlayerInventory.MechanicParts >= float32(turret.Health) {
								assets.Turrets[index].Health -= 1
								assets.PlayerInventory.MechanicParts -= 15 / 100
								assets.DrawPrompter("Destroying turret. %"+strconv.Itoa(100-int(turret.Health)), 23, 235)
							} else if turret.Health <= 0 {
								assets.Turrets = append(assets.Turrets[:index], assets.Turrets[index+1:]...)
							}
						}
					}

				}
			}
		}
	}

	// TURRET COLLIDING PROMPT
	for _, turret := range assets.Turrets {

		if rl.CheckCollisionCircleRec(rl.Vector2{(turret.Position.X + float32(0)/2), (turret.Position.Y + float32(0)/2)}, assets.TurretRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) {
			if !rl.IsKeyDown(rl.KeyE) {
				if turret.OwnerID == assets.PlayerID {
					assets.DrawPrompter("Hold press \"E\" for remove this turret", 23, 235)
				} else if turret.OwnerID != assets.PlayerID {
					assets.DrawPrompter("Hold press \"E\" for destroy this enemy turret", 23, 235)
				}
			}
		}

	}

	// TURRET FIRE
	for index, turret := range assets.Turrets {
		if turret.LockedID == -1 {
			// FACTORY CHECK
			for _, factory := range assets.Factories {
				if factory.OwnerID != int(turret.OwnerID) {
					if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, rl.Rectangle{factory.Position.X, factory.Position.Y, float32(factory.Texture.Width), float32(factory.Texture.Height)}) {

						turret.LockedID = int(factory.ID)
						turret.LockedType = 2
						assets.Turrets[index].Ticker = 0
						assets.Turrets[index].Thick = 0
						assets.Turrets[index].ReloadFPS = 0
					}
				}
			}

			// STATION CHECK
			for _, station := range assets.Stations {
				if station.OwnerID != -1 && station.OwnerID != int(turret.OwnerID) {
					// if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, rl.Rectangle{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2), float32(station.Texture.Width), float32(station.Texture.Height)}) {
					if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, rl.Rectangle{station.Position.X, station.Position.Y, float32(station.Texture.Width), float32(station.Texture.Height)}) {

						assets.Turrets[index].LockedID = int(station.ID)
						assets.Turrets[index].LockedType = 1
						assets.Turrets[index].Ticker = 0
						assets.Turrets[index].Thick = 0
						assets.Turrets[index].ReloadFPS = 0
					}
				}
			}

			for _, otherTurret := range assets.Turrets {
				if otherTurret.OwnerID != turret.OwnerID && otherTurret.ID != turret.ID {
					if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, rl.Rectangle{otherTurret.Position.X, otherTurret.Position.Y, float32(otherTurret.Texture.Width), float32(otherTurret.Texture.Height)}) {

						assets.Turrets[index].LockedID = int(otherTurret.ID)
						assets.Turrets[index].LockedType = 3
						assets.Turrets[index].ReloadFPS = 0
					}
				}
			}
		} else {
			//assets.Turrets[index].LockedID = -1

			var trgtX float32
			var trgtY float32

			if assets.Turrets[index].LockedType == 1 {
				for _, station := range assets.Stations {
					if int(station.ID) == assets.Turrets[index].LockedID {
						trgtX = station.Position.X + float32(station.Texture.Width)/2
						trgtY = station.Position.Y + float32(station.Texture.Height)/2
					}
				}
			}

			if assets.Turrets[index].LockedType == 2 {
				for _, factory := range assets.Factories {
					if int(factory.ID) == assets.Turrets[index].LockedID {
						trgtX = factory.Position.X + float32(factory.Texture.Width)/2
						trgtY = factory.Position.Y + float32(factory.Texture.Height)/2
					}
				}
			}

			if assets.Turrets[index].LockedType == 3 {
				for _, otherTurret := range assets.Turrets {
					if int(otherTurret.ID) == assets.Turrets[index].LockedID {
						trgtX = otherTurret.Position.X + float32(otherTurret.Texture.Width)/2
						trgtY = otherTurret.Position.Y + float32(otherTurret.Texture.Height)/2
					}
				}
			}

			if assets.Turrets[index].Ticker <= 90 {
				assets.Turrets[index].Ticker++
			} else if assets.Turrets[index].Ticker > 90 && assets.Turrets[index].Ticker < 120 {
				assets.Turrets[index].Ticker++
				assets.Turrets[index].Thick++
			} else if assets.Turrets[index].Ticker == 120 {
				assets.Turrets[index].Ticker = 0
				assets.Turrets[index].Thick = 0

				if assets.Turrets[index].LockedType == 1 {
					for stationIndex, station := range assets.Stations {
						if int(station.ID) == assets.Turrets[index].LockedID {
							assets.Stations[stationIndex].Health -= TurretDamage

							if assets.Stations[stationIndex].Health <= 0 {
								assets.Stations[stationIndex].OwnerID = -1
								assets.Turrets[index].LockedID = -1
							}
						}
					}
				}

				if assets.Turrets[index].LockedType == 3 {
					for turretIndex, otherTurret := range assets.Turrets {
						if int(otherTurret.ID) == assets.Turrets[index].LockedID {
							assets.Turrets[turretIndex].Health -= TurretDamage

							if assets.Turrets[turretIndex].Health <= 0 {
								assets.Turrets = append(assets.Turrets[:turretIndex], assets.Turrets[turretIndex+1:]...)
								assets.Turrets[index].LockedID = -1
							}
						}
					}
				}

				if assets.Turrets[index].LockedType == 2 {
					for factoryIndex, factory := range assets.Factories {
						if int(factory.ID) == assets.Turrets[index].LockedID {
							assets.Factories[factoryIndex].Health -= TurretDamage

							if assets.Factories[factoryIndex].Health <= 0 {
								assets.Factories = append(assets.Factories[:factoryIndex], assets.Factories[factoryIndex+1:]...)
								assets.Turrets[index].LockedID = -1
							}
						}
					}
				}

				rl.DrawLineEx(rl.Vector2{turret.Position.X, turret.Position.Y}, rl.Vector2{trgtX, trgtY}, float32(turret.Thick)/5, rl.Red)
			}

			fmt.Println("dıkşın")
		}
	}

	// PROMPTER FOR OUT OF SAFE ZONE AND MONEY
	if promptFlagMoney {
		promptTickerMoney++
		if promptTickerMoney < 90 {
			assets.DrawPrompter("You do not have enough mechanical part", 20, 250)
		} else if promptTickerMoney == 90 {
			promptFlagMoney = false
			promptTickerMoney = 0
		}
	}
}
