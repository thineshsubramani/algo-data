// Package algo provides a simplified facade for generating data structures.
package algo

import (
	"github.com/thineshsubramani/algo-data/array"
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/graph"
	"github.com/thineshsubramani/algo-data/hashtree"
	"github.com/thineshsubramani/algo-data/helper"
	"github.com/thineshsubramani/algo-data/linkedlist"
	"github.com/thineshsubramani/algo-data/tree"
)

// Describe prints help information about a specific component.
var Describe = helper.Describe

// MaxValue is a reference to the central generator limit.
var MaxValue = &generator.MaxValue

// Array provides consistent access to slice generation.
var Array = struct {
	Random     func(int) []int
	Seq        func(int) []int
	SeqReverse func(int) []int
	FromSlice  func([]int) []int
}{
	Random:     array.Random,
	Seq:        array.Seq,
	SeqReverse: array.SeqReverse,
	FromSlice:  array.FromSlice,
}

// LinkedList provides access to singly linked list generation.
var LinkedList = struct {
	Random     func(int) *linkedlist.Node
	Seq        func(int) *linkedlist.Node
	SeqReverse func(int) *linkedlist.Node
	FromSlice  func([]int) *linkedlist.Node
}{
	Random:     linkedlist.Random,
	Seq:        linkedlist.Seq,
	SeqReverse: linkedlist.SeqReverse,
	FromSlice:  linkedlist.FromSlice,
}

// DoublyLinkedList provides access to doubly linked list generation.
var DoublyLinkedList = struct {
	Random     func(int) *linkedlist.DoublyNode
	Seq        func(int) *linkedlist.DoublyNode
	SeqReverse func(int) *linkedlist.DoublyNode
	FromSlice  func([]int) *linkedlist.DoublyNode
}{
	Random:     linkedlist.RandomDoubly,
	Seq:        linkedlist.SeqDoubly,
	SeqReverse: linkedlist.SeqReverseDoubly,
	FromSlice:  linkedlist.FromSliceDoubly,
}

// Tree provides access to binary tree generation and options.
var Tree = struct {
	Random        func(int, ...tree.Option) *tree.Node
	Seq           func(int, ...tree.Option) *tree.Node
	SeqReverse    func(int, ...tree.Option) *tree.Node
	FromSlice     func([]int) *tree.Node
	WithLeftBias  func() tree.Option
	WithRightBias func() tree.Option
}{
	Random:        tree.Random,
	Seq:           tree.Seq,
	SeqReverse:    tree.SeqReverse,
	FromSlice:     tree.FromSlice,
	WithLeftBias:  tree.WithLeftBias,
	WithRightBias: tree.WithRightBias,
}

// HashTree provides access to map generation and options.
var HashTree = struct {
	Random         func(int, ...hashtree.Option) map[int]int
	Seq            func(int, ...hashtree.Option) map[int]int
	SeqReverse     func(int, ...hashtree.Option) map[int]int
	FromSlice      func([]int) map[int]int
	WithMultiplier func(int) hashtree.Option
}{
	Random:         hashtree.Random,
	Seq:            hashtree.Seq,
	SeqReverse:     hashtree.SeqReverse,
	FromSlice:      hashtree.FromSlice,
	WithMultiplier: hashtree.WithMultiplier,
}

// Graph provides access to graph generation and options.
var Graph = struct {
	Random       func(int, ...graph.Option) *graph.Graph
	Seq          func(int, ...graph.Option) *graph.Graph
	SeqReverse   func(int, ...graph.Option) *graph.Graph
	FromSlice    func([]int, ...graph.Option) *graph.Graph
	WithDirected func() graph.Option
	WithWeighted func() graph.Option
	WithDensity  func(float64) graph.Option
}{
	Random:       graph.Random,
	Seq:          graph.Seq,
	SeqReverse:   graph.SeqReverse,
	FromSlice:    graph.FromSlice,
	WithDirected: graph.WithDirected,
	WithWeighted: graph.WithWeighted,
	WithDensity:  graph.WithDensity,
}