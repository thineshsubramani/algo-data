package generator

import (
	"math/rand"
	"time"
)

func Random(size int) []int {
	rand.Seed(time.Now().UnixNano())

	result := make([]int, size)

	for i := range result {
		result[i] = rand.Intn(100)
	}

	return result
}