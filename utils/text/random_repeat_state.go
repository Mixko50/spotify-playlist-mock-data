package text

import "math/rand"

func RandomRepeatState() string {
	devices := [10]string{"off", "track", "context"}
	return devices[rand.Intn(3)]
}
