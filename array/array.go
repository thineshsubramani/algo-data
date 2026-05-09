package array

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "array",
		Functions: []string{
			"Random(size int) []int      // Generates a slice with random values",
			"Seq(size int) []int         // Generates a slice with sequential values (1..N)",
			"FromSlice(nums []int) []int // Identity function returning the slice",
		},
	})
}

// Random generates a slice of a specified size with random values.
func Random(size int) []int {
	return generator.Random(size)
}

// Seq generates a slice with sequential values from 1 to size.
func Seq(size int) []int {
	return generator.Seq(size)
}

// FromSlice is a convenience function that simply returns the input.
func FromSlice(nums []int) []int {
	return nums
}