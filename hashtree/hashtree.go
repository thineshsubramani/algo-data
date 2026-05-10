package hashtree

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "hashtree",
		Functions: []string{
			"Random(size int, opts ...Option) map[int]int // Random key-value pairs",
			"Seq(size int, opts ...Option) map[int]int    // Sequential keys (1..N)",
			"SeqReverse(size int, opts ...Option) map[int]int // Sequential keys (N..1)",
			"FromSlice(nums []int) map[int]int            // Slice values become keys",
		},
		Options: []string{
			"WithMultiplier(m int) // Sets value as key * multiplier",
		},
	})
}

type hashConfig struct {
	multiplier int
}

type Option func(*hashConfig)

// WithMultiplier allows controlling the value generation logic.
func WithMultiplier(m int) Option {
	return func(c *hashConfig) { c.multiplier = m }
}

// Random generates a map with random keys.
func Random(size int, opts ...Option) map[int]int {
	return buildWithOpts(generator.Random(size), opts...)
}

// Seq generates a map with sequential keys.
func Seq(size int, opts ...Option) map[int]int {
	return buildWithOpts(generator.Seq(size), opts...)
}

// SeqReverse generates a map with reverse sequential keys.
func SeqReverse(size int, opts ...Option) map[int]int {
	return buildWithOpts(generator.SeqReverse(size), opts...)
}

// FromSlice converts a slice into a map where slice values are keys.
func FromSlice(nums []int) map[int]int {
	res := make(map[int]int)
	for _, v := range nums {
		res[v] = v
	}
	return res
}

func buildWithOpts(nums []int, opts ...Option) map[int]int {
	cfg := &hashConfig{multiplier: 1} // Default
	for _, opt := range opts {
		opt(cfg)
	}

	res := make(map[int]int)
	for _, v := range nums {
		// Using the value as the key to ensure we have 'size' unique-ish entries
		// if the generator produces duplicates, the map size might be smaller.
		res[v] = v * cfg.multiplier
	}
	return res
}