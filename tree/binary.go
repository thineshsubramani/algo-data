package tree

import (
	"sort"

	"github.com/thineshsubramani/algo-data/generator"
)

// Node represents a node in a binary tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type treeConfig struct {
	bias string // "", "left", or "right"
}

type Option func(*treeConfig)

// WithLeftBias forces the tree to grow exclusively to the left, 
// effectively creating a linked-list structure with maximum depth.
func WithLeftBias() Option {
	return func(c *treeConfig) { c.bias = "left" }
}

// WithRightBias forces the tree to grow exclusively to the right.
func WithRightBias() Option {
	return func(c *treeConfig) { c.bias = "right" }
}

// Random generates a binary tree with random values. 
// By default, it returns a balanced Binary Search Tree.
func Random(size int, opts ...Option) *Node {
	return buildWithOpts(generator.Random(size), opts...)
}

// Seq generates a binary tree with sequential values from 1 to size.
func Seq(size int, opts ...Option) *Node {
	return buildWithOpts(generator.Seq(size), opts...)
}

// FromSlice builds a balanced Binary Search Tree from a slice of integers.
func FromSlice(nums []int) *Node {
	sort.Ints(nums)
	return buildBalanced(nums)
}

func buildWithOpts(nums []int, opts ...Option) *Node {
	cfg := &treeConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.bias == "left" {
		return buildSkewed(nums, true)
	}
	if cfg.bias == "right" {
		return buildSkewed(nums, false)
	}

	// Default: Balanced BST
	return FromSlice(nums)
}

func buildBalanced(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &Node{
		Value: nums[mid],
		Left:  buildBalanced(nums[:mid]),
		Right: buildBalanced(nums[mid+1:]),
	}
}

func buildSkewed(nums []int, left bool) *Node {
	if len(nums) == 0 {
		return nil
	}
	root := &Node{Value: nums[0]}
	current := root
	for i := 1; i < len(nums); i++ {
		newNode := &Node{Value: nums[i]}
		if left { current.Left = newNode; current = current.Left } else { 
			current.Right = newNode; current = current.Right 
		}
	}
	return root
}