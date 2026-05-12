package generator

import (
	"math/rand"
	"time"
)

var (
	// MaxValue allows central control over the range of generated numbers
	MaxValue = 100
)

var itNames = []string{
	"amazon", "google", "microsoft", "apple", "meta",
	"netflix", "adobe", "oracle", "intel", "nvidia",
	"docker", "kubernetes", "golang", "python", "react",
	"aws", "azure", "linux", "git", "github",
}

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

// RandomString returns a slice of random IT-related names.
func RandomString(size int) []string {
	result := make([]string, size)
	for i := range result {
		result[i] = itNames[rand.Intn(len(itNames))]
	}
	return result
}

// RandomByte returns a slice of random characters (a-z).
func RandomByte(size int) []byte {
	result := make([]byte, size)
	for i := range result {
		result[i] = byte('a' + rand.Intn(26))
	}
	return result
}

// Seq returns a sequential slice of integers.
func Seq(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = i + 1
	}
	return result
}

// SeqReverse returns a reverse sequential slice of integers.
func SeqReverse(size int) []int {
	result := make([]int, size)
	for i := range result {
		result[i] = size - i
	}
	return result
}

// Generate is a generic helper that takes a "builder" function (like FromSlice)
// and populates it with centrally generated random data.
func Generate[T any](size int, builder func([]int) T) T {
	data := Random(size)
	return builder(data)
}

func ToAnySlice[T any](slice []T) []any {
	res := make([]any, len(slice))
	for i, v := range slice {
		res[i] = v
	}
	return res
}