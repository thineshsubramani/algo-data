package linkedlist

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "linkedlist",
		Functions: []string{
			"Random(size int, opts ...Option) *Node      // Random values linked list",
			"Seq(size int, opts ...Option) *Node         // Sequential values",
			"SeqReverse(size int, opts ...Option) *Node  // Sequential values",
			"FromSlice(nums []any) *Node // Convert slice to linked list",
		},
		Options: []string{
			"WithStrings() // Use IT-related strings",
			"WithBytes()   // Use characters",
		},
	})
}

type listConfig struct {
	dataType string
}
type Option func(*listConfig)
func WithStrings() Option { return func(c *listConfig) { c.dataType = "string" } }
func WithBytes() Option   { return func(c *listConfig) { c.dataType = "byte" } }

func parseOpts(opts []Option) *listConfig {
	cfg := &listConfig{dataType: "int"}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// Random generates a linked list of a specified size with random values.
func Random(size int, opts ...Option) *Node {
	cfg := parseOpts(opts)
	switch cfg.dataType {
	case "string":
		return FromSlice(generator.ToAnySlice(generator.RandomString(size)))
	case "byte":
		return FromSlice(generator.ToAnySlice(generator.RandomByte(size)))
	default:
		return FromSlice(generator.ToAnySlice(generator.Random(size)))
	}
}

// Seq generates a linked list with sequential values from 1 to size.
func Seq(size int, opts ...Option) *Node {
	cfg := parseOpts(opts)
	if cfg.dataType == "byte" {
		res := make([]any, size)
		for i := 0; i < size; i++ {
			res[i] = string(byte('a' + (i % 26)))
		}
		return FromSlice(res)
	}
	return FromSlice(generator.ToAnySlice(generator.Seq(size)))
}

// SeqReverse generates a linked list with sequential values from size down to 1.
func SeqReverse(size int, opts ...Option) *Node {
	return FromSlice(generator.ToAnySlice(generator.SeqReverse(size)))
}

type Node struct {
	Value any
	Next  *Node
}

func FromSlice(nums []any) *Node {
	if len(nums) == 0 {
		return nil
	}

	head := &Node{
		Value: nums[0],
	}

	current := head

	for _, val := range nums[1:] {
		current.Next = &Node{
			Value: val,
		}

		current = current.Next
	}

	return head
}