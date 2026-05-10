package algodata

import (
	"github.com/thineshsubramani/algo-data/array"
	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/graph"
	"github.com/thineshsubramani/algo-data/hashtree"
	"github.com/thineshsubramani/algo-data/helper"
	"github.com/thineshsubramani/algo-data/linkedlist"
	"github.com/thineshsubramani/algo-data/tree"
)

// Describe is a shortcut to the helper's Describe function.
func Describe(name string) {
	helper.Describe(name)
}

// MaxValue is a pointer to the central generator limit.
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
	Random     func(int) *doublylinkedlist.Node
	Seq        func(int) *doublylinkedlist.Node
	SeqReverse func(int) *doublylinkedlist.Node
	FromSlice  func([]int) *doublylinkedlist.Node
}{
	Random:     doublylinkedlist.Random,
	Seq:        doublylinkedlist.Seq,
	SeqReverse: doublylinkedlist.SeqReverse,
	FromSlice:  doublylinkedlist.FromSlice,
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

// Stack provides access to stack generation.
var Stack = struct {
	Random     func(int) *stack.Stack
	Seq        func(int) *stack.Stack
	SeqReverse func(int) *stack.Stack
	FromSlice  func([]int) *stack.Stack
}{
	Random:     stack.Random,
	Seq:        stack.Seq,
	SeqReverse: stack.SeqReverse,
	FromSlice:  stack.FromSlice,
}