package generator

import (
	"math/rand"
	"time"
)

var (
	// MaxValue allows central control over the range of generated numbers
	MaxValue = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random generates a raw slice of random integers based on central logic.
func Random(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = rand.Intn(MaxValue)
	}
	return result
}

// Generate is a generic helper that takes a "builder" function (like FromSlice)
// and populates it with centrally generated random data.
func Generate[T any](size int, builder func([]int) T) T {
	data := Random(size)
	return builder(data)
}