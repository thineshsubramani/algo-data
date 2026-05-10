package main

import (
	"fmt"
	"github.com/thineshsubramani/algo-data/array"
	"github.com/thineshsubramani/algo-data/graph"
	"github.com/thineshsubramani/algo-data/hashtree"
	"github.com/thineshsubramani/algo-data/helper"
	"github.com/thineshsubramani/algo-data/linkedlist"
	"github.com/thineshsubramani/algo-data/tree"
)

// test the functions
func main() {
	// Get help for specific components
	helper.Describe("tree")
	helper.Describe("linkedlist")
	helper.Describe("array")
	helper.Describe("hashtree")
	helper.Describe("doublylinkedlist")
	helper.Describe("graph")

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

	// --- Array and HashTable Tests ---
	arr := array.Seq(5)
	fmt.Printf("Array: %v\n", arr)

	hash := hashtree.Random(5, hashtree.WithMultiplier(10))
	fmt.Printf("HashTree (Map) sample: %v\n", hash)

	// --- Doubly Linked List Tests ---
	dHead := linkedlist.SeqDoubly(5)
	fmt.Printf("Doubly Start: %v, Next: %v, Next.Prev: %v\n", dHead.Value, dHead.Next.Value, dHead.Next.Prev.Value)

	// --- Graph Tests ---
	g := graph.Random(5, graph.WithDirected(), graph.WithWeighted(), graph.WithDensity(0.5))
	fmt.Printf("Graph Nodes: %v\n", g.Nodes)
	for node, neighbors := range g.Adjacency {
		fmt.Printf("Node %d has %d neighbors\n", node, len(neighbors))
	}
}