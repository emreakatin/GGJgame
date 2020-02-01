package scripts

import (
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
