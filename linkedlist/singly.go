package linkedlist

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "linkedlist",
		Functions: []string{
			"Random(size int) *Node      // Random values linked list",
			"Seq(size int) *Node         // Sequential values (1..N)",
			"SeqReverse(size int) *Node  // Sequential values (N..1)",
			"FromSlice(nums []int) *Node // Convert slice to linked list",
		},
	})
}

// Random generates a linked list of a specified size with random values.
func Random(size int) *Node {
	return generator.Generate(size, FromSlice)
}

// Seq generates a linked list with sequential values from 1 to size.
func Seq(size int) *Node {
	return FromSlice(generator.Seq(size))
}

// SeqReverse generates a linked list with sequential values from size down to 1.
func SeqReverse(size int) *Node {
	return FromSlice(generator.SeqReverse(size))
}

type Node struct {
	Value int
	Next  *Node
}

func FromSlice(nums []int) *Node {
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