package array

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "array",
		Functions: []string{
			"Random(size int, opts ...Option) []any      // Generates a slice with random values",
			"Seq(size int, opts ...Option) []any         // Generates a slice with sequential values",
			"SeqReverse(size int, opts ...Option) []any  // Generates a slice with sequential values",
			"FromSlice(nums []any) []any                // Identity function returning the slice",
		},
		Options: []string{
			"WithStrings() // Use IT-related strings (amazon, google...)",
			"WithBytes()   // Use characters (a, b, c...)",
		},
	})
}

type arrayConfig struct {
	dataType string // "int", "string", "byte"
}
type Option func(*arrayConfig)
func WithStrings() Option { return func(c *arrayConfig) { c.dataType = "string" } }
func WithBytes() Option   { return func(c *arrayConfig) { c.dataType = "byte" } }

func parseOpts(opts []Option) *arrayConfig {
	cfg := &arrayConfig{dataType: "int"}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// Random generates a slice of a specified size with random values.
func Random(size int, opts ...Option) []any {
	cfg := parseOpts(opts)
	switch cfg.dataType {
	case "string":
		return generator.ToAnySlice(generator.RandomString(size))
	case "byte":
		return generator.ToAnySlice(generator.RandomByte(size))
	default:
		return generator.ToAnySlice(generator.Random(size))
	}
}

// Seq generates a slice with sequential values from 1 to size.
func Seq(size int, opts ...Option) []any {
	cfg := parseOpts(opts)
	if cfg.dataType == "byte" {
		res := make([]any, size)
		for i := 0; i < size; i++ {
			res[i] = string(byte('a' + (i % 26)))
		}
		return res
	}
	if cfg.dataType == "string" {
		return generator.ToAnySlice(generator.RandomString(size))
	}
	return generator.ToAnySlice(generator.Seq(size))
}

// SeqReverse generates a slice with sequential values from size down to 1.
func SeqReverse(size int, opts ...Option) []any {
	return generator.ToAnySlice(generator.SeqReverse(size))
}

// FromSlice is a convenience function that simply returns the input.
func FromSlice(nums []any) []any {
	return nums
}