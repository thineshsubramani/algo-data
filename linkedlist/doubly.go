package doublylinkedlist

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "doublylinkedlist",
		Functions: []string{
			"Random(size int) *Node      // Random values doubly linked list",
			"Seq(size int) *Node         // Sequential values (1..N)",
			"SeqReverse(size int) *Node  // Sequential values (N..1)",
			"FromSlice(nums []int) *Node // Convert slice to doubly linked list",
		},
	})
}

// Node represents a node in a doubly linked list.
type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

// Random generates a doubly linked list with random values.
func Random(size int) *Node {
	return generator.Generate(size, FromSlice)
}

// Seq generates a doubly linked list with sequential values from 1 to size.
func Seq(size int) *Node {
	return FromSlice(generator.Seq(size))
}

// SeqReverse generates a doubly linked list with sequential values from size down to 1.
func SeqReverse(size int) *Node {
	return FromSlice(generator.SeqReverse(size))
}

// FromSlice converts a slice of integers into a doubly linked list.
func FromSlice(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	head := &Node{Value: nums[0]}
	current := head

	for _, val := range nums[1:] {
		newNode := &Node{
			Value: val,
			Prev:  current,
		}
		current.Next = newNode
		current = newNode
	}

	return head
}