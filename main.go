package main

import (
	"fmt"
	"github.com/thineshsubramani/algo-data/linkedlist"
	"github.com/thineshsubramani/algo-data/tree"
)

// test the functions
func main() {
	// Directly generate a random linked list
	head := linkedlist.Random(10)
	fmt.Printf("Random Head: %v\n", head.Value)

	// Generate a sequential linked list (1, 2, 3...)
	seqHead := linkedlist.Seq(5)
	fmt.Printf("Sequence Start: %v, Next: %v\n", seqHead.Value, seqHead.Next.Value)

	// Generate a reverse sequential linked list (5, 4, 3...)
	revHead := linkedlist.SeqReverse(5)
	fmt.Printf("Reverse Start: %v, Next: %v\n", revHead.Value, revHead.Next.Value)

	// --- Tree Tests ---
	// Balanced Tree (Default)
	root := tree.Random(10)
	fmt.Printf("Balanced Tree Root: %v\n", root.Value)

	// Skewed Tree (Maximum Depth / Left Bias)
	leftSkewed := tree.Seq(5, tree.WithLeftBias())
	fmt.Printf("Left Skewed Root: %v, Left Child: %v\n", leftSkewed.Value, leftSkewed.Left.Value)
}