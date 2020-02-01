package event

import (
	"fmt"
	"math"
	"strconv"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
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
					Rotation: assets.PlayerRotation,
					Texture:  rl.LoadTexture("sprites/turret.png"),
					LockedID: -1,
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
			// for factoryIndex, factory := range assets.Factories {
			// 	if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, factory.Rectangle) {
			// 		radian := math.Atan2(float64(turret.Position.X-factory.Position.Y), float64(factory.Position.X-turret.Position.Y))
			// 		degree := radian * 180 / math.Pi

			// 		if degree >= 0 && degree <= 90 {
			// 			degree = 90 - degree
			// 		} else if degree <= -90 && degree >= -180 {
			// 			degree += 270 + 2*math.Abs(degree+90)
			// 		} else if degree <= 0 && degree >= -90 {
			// 			degree += 90 + 2*math.Abs(degree)
			// 		} else if degree <= 180 && degree >= 90 {
			// 			degree += 90 + 2*math.Abs(degree-180)
			// 		}

			// 		turret.Rotation = float32(degree)
			// 		turret.LockedID = int(factory.ID)
			// 		turret.LockedType = 2
			// 	}
			// }

			// STATION CHECK
			for _, station := range assets.Stations {
				if station.OwnerID != -1 && station.OwnerID != int(assets.PlayerID) {
					if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, rl.Rectangle{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2), float32(station.Texture.Width), float32(station.Texture.Height)}) {
						radian := math.Atan2(float64((turret.Position.X)-station.Position.Y), float64(station.Position.X-turret.Position.Y))
						degree := radian * 180 / math.Pi
						if degree >= 0 && degree <= 90 {
							degree = -1*degree + 90
						} else if degree <= -90 && degree >= -180 {
							degree += 270 + 2*math.Abs(degree+90)
						} else if degree <= 0 && degree >= -90 {
							degree += 90 + 2*math.Abs(degree)
						} else if degree <= 180 && degree >= 90 {
							degree += 90 + 2*math.Abs(degree-180)
						}

						fmt.Println(degree)
						assets.Turrets[index].Rotation = float32(degree)
						assets.Turrets[index].LockedID = int(station.ID)
						assets.Turrets[index].LockedType = 1
					}
				}
			}

			// OTHER TURRETS CHECK
			for _, otherTurret := range assets.Turrets {
				if otherTurret.OwnerID != assets.PlayerID {
					if rl.CheckCollisionCircleRec(rl.Vector2{turret.Position.X + float32(0)/2, turret.Position.Y + float32(0)/2}, assets.FireRadius, rl.Rectangle{turret.Position.X, turret.Position.Y, float32(turret.Texture.Width), float32(turret.Texture.Height)}) {
						radian := math.Atan2(float64(turret.Position.X-otherTurret.Position.Y), float64(otherTurret.Position.X-turret.Position.Y))
						degree := radian * 180 / math.Pi

						// if degree >= 0 && degree <= 90 {
						// 	degree = 90 - degree
						// } else if degree <= -90 && degree >= -180 {
						// 	degree += 270 + 2*math.Abs(degree+90)
						// } else if degree <= 0 && degree >= -90 {
						// 	degree += 90 + 2*math.Abs(degree)
						// } else if degree <= 180 && degree >= 90 {
						// 	degree += 90 + 2*math.Abs(degree-180)
						// }

						degree = 0
						assets.Turrets[index].Rotation = float32(degree)
						assets.Turrets[index].LockedID = int(otherTurret.ID)
						assets.Turrets[index].LockedType = 3
					}
				}
			}
		} else {
			assets.Turrets[index].LockedID = -1
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
