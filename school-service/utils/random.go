package utils

import "math/rand"

func GetIntBetween(start, end int) int {

	max := (end - start) + 1

	return start + (rand.Intn(max))
}

func LowProbability() bool {

	probability := 0.1

	return rand.Float64() < probability
}
