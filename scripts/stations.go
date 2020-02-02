package scripts

import (
	"github.com/emreakatin/GGJgame/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxStation = 5
)

func GenerateStations() []assets.Station {

	texture := rl.LoadTexture("sprites/station.png")

	stations := []assets.Station{
		assets.Station{
			Position: rl.Vector2{
				25,
				25,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      0,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				330,
				75,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      1,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				150,
				360,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      2,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				370,
				580,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      3,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				622,
				743,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      4,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				630,
				200,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      5,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				935,
				300,
			},
			Texture: texture,
<<<<<<< HEAD
			OwnerID: -1,
=======
			OwnerID: 2,
>>>>>>> ad8897e20f23171f427d14781b9a65bd28109426
			ID:      6,
			Health:  90,
		},
		assets.Station{
			Position: rl.Vector2{
				1220,
				170,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      7,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				930,
				625,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      8,
			Health:  0,
		},
		assets.Station{
			Position: rl.Vector2{
				1530,
				300,
			},
			Texture: texture,
<<<<<<< HEAD
			OwnerID: -1,
=======
			OwnerID: 2,
>>>>>>> ad8897e20f23171f427d14781b9a65bd28109426
			ID:      9,
			Health:  90,
		},
		assets.Station{
			Position: rl.Vector2{
				1320,
				615,
			},
			Texture: texture,
			OwnerID: -1,
			ID:      10,
			Health:  0,
		},
	}

	return stations
}

// var station assets.Station
// 	var stations []assets.Station

// 	var flag bool = false

// 	for i := 0; i < MaxStation; i++ {
// 		randX := randomInt(0, int(assets.Background.Width))
// 		randY := randomInt(0, int(assets.Background.Height))

// 		for _, existStation := range stations {
// 			distance := math.Sqrt(math.Pow(float64(existStation.Position.X-float32(randX)), 2) - math.Pow(float64(existStation.Position.Y-float32(randY)), 2))

// 			if distance < 500 {
// 				break
// 				fmt.Println("distance:", distance, "i:", i)

// 				flag = true
// 			}
// 		}

// 		if !flag {
// 			station.ID = uint(i)
// 			station.Health = 0
// 			station.Situation = 0
// 			station.OwnerID = -1

// 			station.Position = rl.Vector2{float32(randX), float32(randY)}

// 			station.CreateStation()

// 			stations = append(stations, station)
// 		}

// 		flag = false
// 	}
