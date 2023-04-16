package text

import "math/rand"

func RandomDevice() string {
	devices := [10]string{"Macbook Pro", "iPhone 14 Pro", "Windows 11", "Google Chrome"}
	return devices[rand.Intn(4)]
}
