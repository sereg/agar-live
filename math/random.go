package math

import "math/rand"

// Random returns a random value between the min and max
func Random(min, max int) int {
	return int(rand.Intn(max-min) + min)
}
