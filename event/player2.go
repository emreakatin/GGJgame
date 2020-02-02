package event

import (
	"math"
	"strconv"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Degree2          float32 = 0
	Speed2           float32 = 4
	RepairRadius2    float32 = 80
	mousePosition2   rl.Vector2
	Animation2               = false
	Animation2Ticker float32 = 0
)

func Player2Controller() {
	// MOVEMENT
	if rl.IsKeyDown(rl.KeyLeft) && float32(assets.Player2Position.X-Speed2) > 0 {
		if !isColliding2(rl.Vector2{assets.Player2Position.X - Speed2, assets.Player2Position.Y}) {
			assets.Player2Position.X -= Speed2
			// assets.Camera.Offset.X -= Speed2
			assets.Player2Rotation = 270
			Animation2 = true
			Player2Animation2()
		}
	}

	if rl.IsKeyDown(rl.KeyRight) && float32(assets.Player2Position.X+float32(assets.Player2.Width)+Speed2) < float32(assets.Background.Width) {
		if !isColliding2(rl.Vector2{assets.Player2Position.X + Speed2, assets.Player2Position.Y}) {
			assets.Player2Position.X += Speed2
			// assets.Camera.Offset.X += Speed2
			assets.Player2Rotation = 90
			Animation2 = true
			Player2Animation2()
		}
	}

	if rl.IsKeyDown(rl.KeyUp) && float32(assets.Player2Position.Y-Speed2) > 0 {
		if !isColliding2(rl.Vector2{assets.Player2Position.X, assets.Player2Position.Y - Speed2}) {
			assets.Player2Position.Y -= Speed2
			// assets.Camera.Offset.Y -= Speed2
			assets.Player2Rotation = 0
			Animation2 = true
			Player2Animation2()
		}
	}

	if rl.IsKeyDown(rl.KeyDown) && float32(assets.Player2Position.Y+float32(assets.Player2.Height)+Speed2) < float32(assets.Background.Height) {
		if !isColliding2(rl.Vector2{assets.Player2Position.X, assets.Player2Position.Y + Speed2}) {
			assets.Player2Position.Y += Speed2
			// assets.Camera.Offset.Y += Speed2
			assets.Player2Rotation = 180
			Animation2 = true
			Player2Animation2()
		}
	}

	// IF NEAR BY ANY TOWER
	for index, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, RepairRadius2, rl.Rectangle{float32(assets.Player2Position.X), float32(assets.Player2Position.Y), float32(assets.Player2.Width * assets.Player2Scale), float32(assets.Player2.Height * assets.Player2Scale)}) {
			if station.OwnerID == -1 && !rl.IsKeyDown(rl.KeyRightShift) && assets.Player2Inventory.MechanicParts >= 0 {
				assets.DrawPrompter("Repairing the tower will cost "+strconv.Itoa(int(100.0-assets.Stations[index].Health))+" Mechanical Part. Press \"E\" for repair this tower", 23, 250)
			} else if station.OwnerID == -1 && rl.IsKeyDown(rl.KeyRightShift) && assets.Player2Inventory.MechanicParts <= 0 {
				assets.DrawPrompter("You do not have enough money!", 23, 250)
			} else if station.OwnerID == int(assets.Player2ID) {
				assets.DrawPrompter("You're in safe!", 20, 410)
			}

			if assets.Stations[index].Health <= 100.0 && assets.Player2Inventory.MechanicParts > 0 {
<<<<<<< HEAD
				if rl.IsKeyDown(rl.KeyRightShift) {
					assets.Player2 = rl.LoadTexture("sprites/p1_8.png")
=======
				if rl.IsKeyDown(rl.KeyE) {
					assets.Player2 = rl.LoadTexture("sprites/p2_8.png")
>>>>>>> 9cf52652e1e7288d6701863c87790ec0235a8601
					if station.OwnerID == -1 {
						assets.Stations[index].Health += 0.25
						assets.Player2Inventory.MechanicParts -= 0.2
						assets.DrawPrompter("Repairing! %"+strconv.Itoa(int(assets.Stations[index].Health)), 23, 250)
						if assets.Stations[index].Health == 100 {
							assets.Stations[index].OwnerID = int(assets.Player2ID)
						}
					} else {
						if station.OwnerID != int(assets.Player2ID) && station.OwnerID != -1 && assets.Stations[index].Health > 0 {
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

func isColliding2(nextPosition rl.Vector2) bool {
	// STATIONS and TURRETS or etc. will be added
	for _, station := range assets.Stations {

		if rl.CheckCollisionRecs(rl.Rectangle{station.Position.X, station.Position.Y, float32(station.Texture.Width), float32(station.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.Player2Rectangle.Width, assets.Player2Rectangle.Height}) {
			return true
		}
	}
	for _, factory := range assets.Factories {
		if rl.CheckCollisionRecs(rl.Rectangle{factory.Position.X, factory.Position.Y, float32(factory.Texture.Width), float32(factory.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.Player2Rectangle.Width, assets.Player2Rectangle.Height}) {
			return true
		}
	}
	return false
}

func Player2Animation2() {
	if Animation2 == true {
		Animation2Ticker++
		if Animation2Ticker <= 4 {
			assets.Player2 = rl.LoadTexture("sprites/p2_0.png")
		} else if Animation2Ticker <= 8 {
			assets.Player2 = rl.LoadTexture("sprites/p2_1.png")
		} else if Animation2Ticker <= 12 {
			assets.Player2 = rl.LoadTexture("sprites/p2_2.png")
		} else if Animation2Ticker <= 16 {
			assets.Player2 = rl.LoadTexture("sprites/p2_3.png")
		} else if Animation2Ticker <= 20 {
			assets.Player2 = rl.LoadTexture("sprites/p2_4.png")
		} else if Animation2Ticker <= 24 {
			assets.Player2 = rl.LoadTexture("sprites/p2_5.png")
		} else if Animation2Ticker <= 28 {
			assets.Player2 = rl.LoadTexture("sprites/p2_6.png")
		} else if Animation2Ticker <= 32 {
			assets.Player2 = rl.LoadTexture("sprites/p2_7.png")
		} else if Animation2Ticker <= 36 {
			assets.Player2 = rl.LoadTexture("sprites/p2_6.png")
		} else if Animation2Ticker <= 40 {
			assets.Player2 = rl.LoadTexture("sprites/p2_5.png")
		} else if Animation2Ticker <= 44 {
			assets.Player2 = rl.LoadTexture("sprites/p2_4.png")
		} else if Animation2Ticker <= 48 {
			assets.Player2 = rl.LoadTexture("sprites/p2_3.png")
		} else if Animation2Ticker <= 52 {
			assets.Player2 = rl.LoadTexture("sprites/p2_2.png")
		} else if Animation2Ticker <= 56 {
			assets.Player2 = rl.LoadTexture("sprites/p2_1.png")
		} else if Animation2Ticker <= 60 {
			assets.Player2 = rl.LoadTexture("sprites/p2_0.png")
			Animation2 = false
			Animation2Ticker = 0
		}
	}
}

func Rotation2() {
	// ROTATION
	mousePosition2 = rl.GetMousePosition()

	// if mouse in screen
	if !(mousePosition2.X < 0) && !(mousePosition2.Y < 0) {
		radian := math.Atan2(float64(assets.Player2Position.Y-mousePosition2.Y), float64(mousePosition2.X-assets.Player2Position.X))
		Degree2 := radian * 180 / math.Pi

		if Degree2 >= 0 && Degree2 <= 90 {
			Degree2 = 90 - Degree2
		} else if Degree2 <= -90 && Degree2 >= -180 {
			Degree2 += 270 + 2*math.Abs(Degree2+90)
		} else if Degree2 <= 0 && Degree2 >= -90 {
			Degree2 += 90 + 2*math.Abs(Degree2)
		} else if Degree2 <= 180 && Degree2 >= 90 {
			Degree2 += 90 + 2*math.Abs(Degree2-180)
		}
	}
	assets.Player2Rotation = Degree2
}
