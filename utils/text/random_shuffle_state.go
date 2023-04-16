package text

import "math/rand"

func RandomShuffleState() bool {
	return rand.Intn(2) == 0
}
