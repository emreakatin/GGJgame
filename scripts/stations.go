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
		},
		assets.Station{
			Position: rl.Vector2{
				330,
				75,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				245,
				390,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				490,
				600,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				700,
				860,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				630,
				200,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				935,
				300,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				1220,
				170,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				1040,
				670,
			},
			Texture: texture,
		},
		assets.Station{
			Position: rl.Vector2{
				1530,
				300,
			},
			Texture: texture,
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
