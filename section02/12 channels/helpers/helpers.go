package helpers

import (
	"math/rand"
)

func RandomNumber(n int) int {
	/* rand.Seed is deprecated: As of Go 1.20 there is no reason to call Seed with a random value.
	Programs that call Seed with a known value to get a specific sequence of results
	should use New(NewSource(seed)) to obtain a local random generator*/
	// rand.Seed(time.Now().UnixNano())

	value := rand.Intn(n)
	return value
}
