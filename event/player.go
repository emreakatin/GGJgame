package event

import (
	"math"
	"strconv"

	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Speed         float32 = 4
	RepairRadius  float32 = 80
	mousePosition rl.Vector2
)

func PlayerController() {
	// ROTATION
	mousePosition = rl.GetMousePosition()

	// if mouse in screen
	if !(mousePosition.X < 0) && !(mousePosition.Y < 0) {
		radian := math.Atan2(float64(assets.PlayerPosition.Y-mousePosition.Y), float64(mousePosition.X-assets.PlayerPosition.X))
		degree := radian * 180 / math.Pi

		if degree >= 0 && degree <= 90 {
			degree = 90 - degree
		} else if degree <= -90 && degree >= -180 {
			degree += 270 + 2*math.Abs(degree+90)
		} else if degree <= 0 && degree >= -90 {
			degree += 90 + 2*math.Abs(degree)
		} else if degree <= 180 && degree >= 90 {
			degree += 90 + 2*math.Abs(degree-180)
		}

		assets.PlayerRotation = float32(degree)
	}

	// MOVEMENT
	if rl.IsKeyDown(rl.KeyA) && float32(assets.PlayerPosition.X-Speed) > 0 {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X - Speed, assets.PlayerPosition.Y}) {
			assets.PlayerPosition.X -= Speed
			assets.Camera.Offset.X -= Speed
		}
	}

	if rl.IsKeyDown(rl.KeyD) && float32(assets.PlayerPosition.X+float32(assets.Player.Width)+Speed) < float32(assets.Background.Width) {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X + Speed, assets.PlayerPosition.Y}) {
			assets.PlayerPosition.X += Speed
			assets.Camera.Offset.X += Speed
		}
	}

	if rl.IsKeyDown(rl.KeyW) && float32(assets.PlayerPosition.Y-Speed) > 0 {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X, assets.PlayerPosition.Y - Speed}) {
			assets.PlayerPosition.Y -= Speed
			assets.Camera.Offset.Y -= Speed
		}
	}

	if rl.IsKeyDown(rl.KeyS) && float32(assets.PlayerPosition.Y+float32(assets.Player.Height)+Speed) < float32(assets.Background.Height) {
		if !isColliding(rl.Vector2{assets.PlayerPosition.X, assets.PlayerPosition.Y + Speed}) {
			assets.PlayerPosition.Y += Speed
			assets.Camera.Offset.Y += Speed
		}
	}

	// IF NEAR BY ANY TOWER
	for index, station := range assets.Stations {
		if rl.CheckCollisionCircleRec(rl.Vector2{station.Position.X + float32(station.Texture.Width/2), station.Position.Y + float32(station.Texture.Height/2)}, RepairRadius, rl.Rectangle{float32(assets.PlayerPosition.X), float32(assets.PlayerPosition.Y), float32(assets.Player.Width * assets.PlayerScale), float32(assets.Player.Height * assets.PlayerScale)}) {
			if station.OwnerID == -1 && !rl.IsKeyDown(rl.KeyE) && assets.PlayerInventory.MechanicParts >= 0 {
				assets.DrawPrompter("Repairing the tower will cost "+strconv.Itoa(int(100.0-assets.Stations[index].Health))+" Mechanical Part. Press \"E\" for repair this tower", 23, 250)
			} else if station.OwnerID == -1 && rl.IsKeyDown(rl.KeyE) && assets.PlayerInventory.MechanicParts <= 0 {
				assets.DrawPrompter("You do not have enough money!", 23, 250)
			} else if station.OwnerID == int(assets.PlayerID) {
				assets.DrawPrompter("You're in safe!", 23, 250)
			}

			if assets.Stations[index].Health <= 100.0 && assets.PlayerInventory.MechanicParts > 0 {
				if rl.IsKeyDown(rl.KeyE) {
					if station.OwnerID != int(assets.PlayerID) {
						assets.Stations[index].Health += 0.25
						assets.PlayerInventory.MechanicParts -= 0.2
						assets.DrawPrompter("Repairing! %"+strconv.Itoa(int(assets.Stations[index].Health)), 23, 250)
						if assets.Stations[index].Health == 100 {
							assets.Stations[index].OwnerID = int(assets.PlayerID)
						}
					} else {
						if assets.Stations[index].Health > 0.0 && station.OwnerID != int(assets.PlayerID) {
							assets.Stations[index].Health -= 0.2
							assets.DrawPrompter("Destroying enemy tower! %"+strconv.Itoa(int(100-assets.Stations[index].Health)), 23, 250)
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
	// station.UpdateStation()
	assets.UpdateCamera()
}

func isColliding(nextPosition rl.Vector2) bool {
	// STATIONS and TURRETS or etc. will be added
	for _, station := range assets.Stations {

		if rl.CheckCollisionRecs(rl.Rectangle{station.Position.X, station.Position.Y, float32(station.Texture.Width), float32(station.Texture.Height)}, rl.Rectangle{nextPosition.X, nextPosition.Y, assets.PlayerRectangle.Width, assets.PlayerRectangle.Height}) {
			return true
		}
	}

	return false
}
