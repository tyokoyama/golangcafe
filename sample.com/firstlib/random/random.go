package random

import "math/rand"

func RandInt(x int) int {
	rand.Seed(int64(x))

	return rand.Int()
}
