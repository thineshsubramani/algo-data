package linkedlist

import (
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "doublylinkedlist",
		Functions: []string{
			"RandomDoubly(size int) *DoublyNode      // Random values doubly linked list",
			"SeqDoubly(size int) *DoublyNode         // Sequential values (1..N)",
			"SeqReverseDoubly(size int) *DoublyNode  // Sequential values (N..1)",
			"FromSliceDoubly(nums []int) *DoublyNode // Convert slice to doubly linked list",
		},
	})
}

// DoublyNode represents a node in a doubly linked list.
type DoublyNode struct {
	Value int
	Next  *DoublyNode
	Prev  *DoublyNode
}

// RandomDoubly generates a doubly linked list with random values.
func RandomDoubly(size int) *DoublyNode {
	return generator.Generate(size, FromSliceDoubly)
}

// SeqDoubly generates a doubly linked list with sequential values from 1 to size.
func SeqDoubly(size int) *DoublyNode {
	return FromSliceDoubly(generator.Seq(size))
}

// SeqReverseDoubly generates a doubly linked list with sequential values from size down to 1.
func SeqReverseDoubly(size int) *DoublyNode {
	return FromSliceDoubly(generator.SeqReverse(size))
}

// FromSliceDoubly converts a slice of integers into a doubly linked list.
func FromSliceDoubly(nums []int) *DoublyNode {
	if len(nums) == 0 {
		return nil
	}

	head := &DoublyNode{Value: nums[0]}
	current := head

	for _, val := range nums[1:] {
		newNode := &DoublyNode{
			Value: val,
			Prev:  current,
		}
		current.Next = newNode
		current = newNode
	}

	return head
}