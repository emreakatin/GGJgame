package event

import (
	"math"
	"strconv"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Degree          float32 = 0
	Speed           float32 = 4
	RepairRadius    float32 = 80
	mousePosition   rl.Vector2
	Animation               = false
	AnimationTicker float32 = 0
)

func PlayerController() {
	// MOVEMENT
	if rl.IsKeyDown(rl.KeyA) && float32(assets.PlayerPosition.X-Speed) > 0 {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X - Speed, assets.PlayerPosition.Y}) {
			assets.PlayerPosition.X -= Speed
			assets.Camera.Offset.X -= Speed
			assets.PlayerRotation = 270
			Animation = true
			PlayerAnimation()
		}
	}

	if rl.IsKeyDown(rl.KeyD) && float32(assets.PlayerPosition.X+float32(assets.Player.Width)+Speed) < float32(assets.Background.Width) {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X + Speed, assets.PlayerPosition.Y}) {
			assets.PlayerPosition.X += Speed
			assets.Camera.Offset.X += Speed
			assets.PlayerRotation = 90
			Animation = true
			PlayerAnimation()
		}
	}

	if rl.IsKeyDown(rl.KeyW) && float32(assets.PlayerPosition.Y-Speed) > 0 {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X, assets.PlayerPosition.Y - Speed}) {
			assets.PlayerPosition.Y -= Speed
			assets.Camera.Offset.Y -= Speed
			assets.PlayerRotation = 0
			Animation = true
			PlayerAnimation()
		}
	}

	if rl.IsKeyDown(rl.KeyS) && float32(assets.PlayerPosition.Y+float32(assets.Player.Height)+Speed) < float32(assets.Background.Height) {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X, assets.PlayerPosition.Y + Speed}) {
			assets.PlayerPosition.Y += Speed
			assets.Camera.Offset.Y += Speed
			assets.PlayerRotation = 180
			Animation = true
			PlayerAnimation()
		}
	}

	// IF NEAR BY ANY TOWER
	for index, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, RepairRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) {
			if station.OwnerID == -1 && !rl.IsKeyDown(rl.KeyE) && assets.PlayerInventory.MechanicParts >= 0 {
				assets.DrawPrompter("Repairing the tower will cost "+strconv.Itoa(int(100.0-assets.Stations[index].Health))+" Mechanical Part. Press \"E\" to repair this tower", 23, 250)
			} else if station.OwnerID == -1 && rl.IsKeyDown(rl.KeyE) && assets.PlayerInventory.MechanicParts <= 0 {
				assets.DrawPrompter("You do not have enough money!", 23, 250)
			} else if station.OwnerID == int(assets.PlayerID) {
				assets.DrawPrompter("You're in safe!", 20, 410)
			} else if station.OwnerID != int(assets.PlayerID) && station.OwnerID != -1 && assets.Stations[index].Health > 0 && !rl.IsKeyDown(rl.KeyE) {
				assets.DrawPrompter("Destroying the enemy tower will cost "+strconv.Itoa(int(assets.Stations[index].Health))+" Mechanical Part. Press \"E\" to destroy this tower", 23, 250)
			}
			if station.OwnerID == int(assets.PlayerID) && !rl.IsKeyDown(rl.KeyE) && assets.Stations[index].Health < 100 {
				assets.DrawPrompter("Somebody attacked your tower! Repairing the tower will cost "+strconv.Itoa(int(100.0-assets.Stations[index].Health))+" Mechanical Part. Press \"E\" to repair this tower", 23, 250)
			}
			if assets.Stations[index].Health < 100.0 && assets.PlayerInventory.MechanicParts > 0 {
				if rl.IsKeyDown(rl.KeyE) {
					assets.Player = rl.LoadTexture("sprites/p1_8.png")
					if station.OwnerID == -1 {
						assets.Stations[index].Health += 0.25
						assets.PlayerInventory.MechanicParts -= 0.25
						assets.DrawPrompter("Repairing! %"+strconv.Itoa(int(assets.Stations[index].Health)), 23, 250)
						if assets.Stations[index].Health == 100 {
							assets.Stations[index].OwnerID = int(assets.PlayerID)
						}
					} else if station.OwnerID == int(assets.PlayerID) {
						assets.Stations[index].Health += 0.25
						assets.PlayerInventory.MechanicParts -= 0.25
						assets.DrawPrompter("Repairing! %"+strconv.Itoa(int(assets.Stations[index].Health)), 23, 250)
					} else {
						if station.OwnerID != int(assets.PlayerID) && station.OwnerID != -1 && assets.Stations[index].Health > 0 {
							assets.Stations[index].Health -= 0.5
							assets.PlayerInventory.MechanicParts -= 0.5
							assets.DrawPrompter("Destroying enemy tower! %"+strconv.Itoa(int(assets.Stations[index].Health)), 23, 250)
						} else if assets.Stations[index].Health == 0 {
							assets.Stations[index].OwnerID = -1
						}
					}
				}

				// STATION ANIMATION
				if station.Health <= 25 {
					assets.Stations[index].Texture = rl.LoadTexture("sprites/station0.png")
				} else if station.Health <= 50 && station.Health > 25 {
					assets.Stations[index].Texture = rl.LoadTexture("sprites/station1.png")
				} else if station.Health <= 75 && station.Health > 50 {
					assets.Stations[index].Texture = rl.LoadTexture("sprites/station2.png")
				} else if station.Health < 100 && station.Health > 75 {
					assets.Stations[index].Texture = rl.LoadTexture("sprites/station3.png")
				} else if station.Health >= 100 {
					assets.Stations[index].Texture = rl.LoadTexture("sprites/station4.png")
				}

			}
		}
	}
	assets.UpdateCamera()
	// DESTROYING ENEMIES FACTORIES
	for index, factory := range assets.Factories {
		if rl.CheckCollisionCircleRec(rl.Vector2{factory.Position.X + float32(factory.Texture.Width/2), factory.Position.Y + float32(factory.Texture.Height/2)}, RepairRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) {
			if factory.OwnerID != int(assets.PlayerID) && !rl.IsKeyDown(rl.KeyE) && assets.PlayerInventory.MechanicParts >= 0 {
				assets.DrawPrompter("Destroying the enemy factory will cost "+strconv.Itoa(int(assets.Factories[index].Health))+" Mechanical Part. Press \"E\" to destroy this factory", 23, 250)
			}
			if factory.OwnerID != int(assets.PlayerID) && rl.IsKeyDown(rl.KeyE) && assets.PlayerInventory.MechanicParts >= 0 {
				assets.DrawPrompter("Destroying enemy factory! %"+strconv.Itoa(int(assets.Factories[index].Health)), 23, 250)
				assets.Factories[index].Health -= 0.5
				assets.PlayerInventory.MechanicParts -= 0.5
				if assets.Factories[index].Health == 0 {
					assets.Factories = append(assets.Factories[:index], assets.Factories[index+1:]...)
				}
			}
		}
	}
}

func isColliding(nextPosition rl.Vector2) bool {
	// STATIONS and TURRETS or etc. will be added
	for _, station := range assets.Stations {

		if rl.CheckCollisionRecs(rl.Rectangle{station.Position.X, station.Position.Y, float32(station.Texture.Width), float32(station.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.PlayerRectangle.Width, assets.PlayerRectangle.Height}) {
			return true
		}
	}
	for _, factory := range assets.Factories {
		if rl.CheckCollisionRecs(rl.Rectangle{factory.Position.X, factory.Position.Y - 15.0, float32(factory.Texture.Width), float32(factory.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.PlayerRectangle.Width, assets.PlayerRectangle.Height}) {
			return true
		}
	}
	return false
}

func PlayerAnimation() {
	if Animation == true {
		AnimationTicker++
		if AnimationTicker <= 4 {
			assets.Player = rl.LoadTexture("sprites/p1_1.png")
		} else if AnimationTicker <= 8 {
			assets.Player = rl.LoadTexture("sprites/p1_2.png")
		} else if AnimationTicker <= 12 {
			assets.Player = rl.LoadTexture("sprites/p1_3.png")
		} else if AnimationTicker <= 16 {
			assets.Player = rl.LoadTexture("sprites/p1_4.png")
		} else if AnimationTicker <= 20 {
			assets.Player = rl.LoadTexture("sprites/p1_5.png")
		} else if AnimationTicker <= 24 {
			assets.Player = rl.LoadTexture("sprites/p1_6.png")
		} else if AnimationTicker <= 28 {
			assets.Player = rl.LoadTexture("sprites/p1_7.png")
		} else if AnimationTicker <= 32 {
			assets.Player = rl.LoadTexture("sprites/p1_8.png")
		} else if AnimationTicker <= 36 {
			assets.Player = rl.LoadTexture("sprites/p1_7.png")
		} else if AnimationTicker <= 40 {
			assets.Player = rl.LoadTexture("sprites/p1_6.png")
		} else if AnimationTicker <= 44 {
			assets.Player = rl.LoadTexture("sprites/p1_5.png")
		} else if AnimationTicker <= 48 {
			assets.Player = rl.LoadTexture("sprites/p1_4.png")
		} else if AnimationTicker <= 52 {
			assets.Player = rl.LoadTexture("sprites/p1_3.png")
		} else if AnimationTicker <= 56 {
			assets.Player = rl.LoadTexture("sprites/p1_2.png")
		} else if AnimationTicker <= 60 {
			assets.Player = rl.LoadTexture("sprites/p1_1.png")
			Animation = false
			AnimationTicker = 0
		}
	}
}

func Rotation() {
	// ROTATION
	mousePosition = rl.GetMousePosition()

	// if mouse in screen
	if !(mousePosition.X < 0) && !(mousePosition.Y < 0) {
		radian := math.Atan2(float64(assets.PlayerPosition.Y-mousePosition.Y), float64(mousePosition.X-assets.PlayerPosition.X))
		Degree := radian * 180 / math.Pi

		if Degree >= 0 && Degree <= 90 {
			Degree = 90 - Degree
		} else if Degree <= -90 && Degree >= -180 {
			Degree += 270 + 2*math.Abs(Degree+90)
		} else if Degree <= 0 && Degree >= -90 {
			Degree += 90 + 2*math.Abs(Degree)
		} else if Degree <= 180 && Degree >= 90 {
			Degree += 90 + 2*math.Abs(Degree-180)
		}
	}
	assets.PlayerRotation = Degree
}
