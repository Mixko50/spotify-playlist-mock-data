package text

import (
	"math/rand"
	"time"
)

func RandomNumber(length int) int {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and length inclusive
	randomNumber := rand.Intn(length)

	return randomNumber
}
