package event

import (
	"math"
	"strconv"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Degree3          float32 = 0
	Speed3           float32 = 4
	RepairRadius3    float32 = 80
	mousePosition3   rl.Vector2
	Animation3               = false
	Animation3Ticker float32 = 0
)

func Player3Controller() {
	// MOVEMENT
	if rl.IsKeyDown(rl.KeyH) && float32(assets.Player3Position.X-Speed3) > 0 {
		if !isColliding3(rl.Vector2{assets.Player3Position.X - Speed3, assets.Player3Position.Y}) {
			assets.Player3Position.X -= Speed3
			// assets.Camera.Offset.X -= Speed3
			assets.Player3Rotation = 270
			Animation3 = true
			Player3Animation3()
		}
	}

	if rl.IsKeyDown(rl.KeyK) && float32(assets.Player3Position.X+float32(assets.Player3.Width)+Speed3) < float32(assets.Background.Width) {
		if !isColliding3(rl.Vector2{assets.Player3Position.X + Speed3, assets.Player3Position.Y}) {
			assets.Player3Position.X += Speed3
			// assets.Camera.Offset.X += Speed3
			assets.Player3Rotation = 90
			Animation3 = true
			Player3Animation3()
		}
	}

	if rl.IsKeyDown(rl.KeyU) && float32(assets.Player3Position.Y-Speed3) > 0 {
		if !isColliding3(rl.Vector2{assets.Player3Position.X, assets.Player3Position.Y - Speed3}) {
			assets.Player3Position.Y -= Speed3
			// assets.Camera.Offset.Y -= Speed3
			assets.Player3Rotation = 0
			Animation3 = true
			Player3Animation3()
		}
	}

	if rl.IsKeyDown(rl.KeyJ) && float32(assets.Player3Position.Y+float32(assets.Player3.Height)+Speed3) < float32(assets.Background.Height) {
		if !isColliding3(rl.Vector2{assets.Player3Position.X, assets.Player3Position.Y + Speed3}) {
			assets.Player3Position.Y += Speed3
			// assets.Camera.Offset.Y += Speed3
			assets.Player3Rotation = 180
			Animation3 = true
			Player3Animation3()
		}
	}

	// IF NEAR BY ANY TOWER
	for index, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, RepairRadius3, rl.Rectangle{float32(assets.Player3Position.X), float32(assets.Player3Position.Y), float32(assets.Player3.Width * assets.Player3Scale), float32(assets.Player3.Height * assets.Player3Scale)}) {
			if station.OwnerID == -1 && !rl.IsKeyDown(rl.KeyE) && assets.Player3Inventory.MechanicParts >= 0 {
				assets.DrawPrompter("Repairing the tower will cost "+strconv.Itoa(int(100.0-assets.Stations[index].Health))+" Mechanical Part. Press \"E\" for repair this tower", 23, 250)
			} else if station.OwnerID == -1 && rl.IsKeyDown(rl.KeyE) && assets.Player3Inventory.MechanicParts <= 0 {
				assets.DrawPrompter("You do not have enough money!", 23, 250)
			} else if station.OwnerID == int(assets.Player3ID) {
				assets.DrawPrompter("You're in safe!", 20, 410)
			}

			if assets.Stations[index].Health <= 100.0 && assets.Player3Inventory.MechanicParts > 0 {
				if rl.IsKeyDown(rl.KeyE) {
					assets.Player3 = rl.LoadTexture("sprites/p1_8.png")
					if station.OwnerID == -1 {
						assets.Stations[index].Health += 0.25
						assets.Player3Inventory.MechanicParts -= 0.2
						assets.DrawPrompter("Repairing! %"+strconv.Itoa(int(assets.Stations[index].Health)), 23, 250)
						if assets.Stations[index].Health == 100 {
							assets.Stations[index].OwnerID = int(assets.Player3ID)
						}
					} else {
						if station.OwnerID != int(assets.Player3ID) && station.OwnerID != -1 && assets.Stations[index].Health > 0 {
							assets.Stations[index].Health -= 0.5
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
}

func isColliding3(nextPosition rl.Vector2) bool {
	// STATIONS and TURRETS or etc. will be added
	for _, station := range assets.Stations {

		if rl.CheckCollisionRecs(rl.Rectangle{station.Position.X, station.Position.Y, float32(station.Texture.Width), float32(station.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.Player3Rectangle.Width, assets.Player3Rectangle.Height}) {
			return true
		}
	}
	for _, factory := range assets.Factories {
		if rl.CheckCollisionRecs(rl.Rectangle{factory.Position.X, factory.Position.Y, float32(factory.Texture.Width), float32(factory.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.Player3Rectangle.Width, assets.Player3Rectangle.Height}) {
			return true
		}
	}
	return false
}

func Player3Animation3() {
	if Animation3 == true {
		Animation3Ticker++
		if Animation3Ticker <= 4 {
			assets.Player3 = rl.LoadTexture("sprites/p1_1.png")
		} else if Animation3Ticker <= 8 {
			assets.Player3 = rl.LoadTexture("sprites/p1_2.png")
		} else if Animation3Ticker <= 12 {
			assets.Player3 = rl.LoadTexture("sprites/p1_3.png")
		} else if Animation3Ticker <= 16 {
			assets.Player3 = rl.LoadTexture("sprites/p1_4.png")
		} else if Animation3Ticker <= 20 {
			assets.Player3 = rl.LoadTexture("sprites/p1_5.png")
		} else if Animation3Ticker <= 24 {
			assets.Player3 = rl.LoadTexture("sprites/p1_6.png")
		} else if Animation3Ticker <= 28 {
			assets.Player3 = rl.LoadTexture("sprites/p1_7.png")
		} else if Animation3Ticker <= 32 {
			assets.Player3 = rl.LoadTexture("sprites/p1_8.png")
		} else if Animation3Ticker <= 36 {
			assets.Player3 = rl.LoadTexture("sprites/p1_7.png")
		} else if Animation3Ticker <= 40 {
			assets.Player3 = rl.LoadTexture("sprites/p1_6.png")
		} else if Animation3Ticker <= 44 {
			assets.Player3 = rl.LoadTexture("sprites/p1_5.png")
		} else if Animation3Ticker <= 48 {
			assets.Player3 = rl.LoadTexture("sprites/p1_4.png")
		} else if Animation3Ticker <= 52 {
			assets.Player3 = rl.LoadTexture("sprites/p1_3.png")
		} else if Animation3Ticker <= 56 {
			assets.Player3 = rl.LoadTexture("sprites/p1_2.png")
		} else if Animation3Ticker <= 60 {
			assets.Player3 = rl.LoadTexture("sprites/p1_1.png")
			Animation3 = false
			Animation3Ticker = 0
		}
	}
}

func Rotation3() {
	// ROTATION
	mousePosition3 = rl.GetMousePosition()

	// if mouse in screen
	if !(mousePosition3.X < 0) && !(mousePosition3.Y < 0) {
		radian := math.Atan2(float64(assets.Player3Position.Y-mousePosition3.Y), float64(mousePosition3.X-assets.Player3Position.X))
		Degree3 := radian * 180 / math.Pi

		if Degree3 >= 0 && Degree3 <= 90 {
			Degree3 = 90 - Degree3
		} else if Degree3 <= -90 && Degree3 >= -180 {
			Degree3 += 270 + 2*math.Abs(Degree3+90)
		} else if Degree3 <= 0 && Degree3 >= -90 {
			Degree3 += 90 + 2*math.Abs(Degree3)
		} else if Degree3 <= 180 && Degree3 >= 90 {
			Degree3 += 90 + 2*math.Abs(Degree3-180)
		}
	}
	assets.Player3Rotation = Degree3
}
