package graph

import (
	"math/rand"

	"github.com/thineshsubramani/algo-data/generator"
	"github.com/thineshsubramani/algo-data/helper"
)

func init() {
	helper.Register(helper.ComponentInfo{
		Name: "graph",
		Functions: []string{
			"Random(size int, opts ...Option) *Graph // Randomly connected nodes",
			"Seq(size int, opts ...Option) *Graph    // Linear path graph (1->2->3...)",
			"SeqReverse(size int, opts ...Option) *Graph // Linear path graph (N->N-1->...1)",
			"FromSlice(nums []int, opts ...Option) *Graph // Builds graph from slice",
		},
		Options: []string{
			"WithDirected()   // Edges are one-way",
			"WithWeighted()   // Random weights assigned to edges",
			"WithDensity(f)   // Probability (0..1) of connection (default 0.25)",
		},
	})
}

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Nodes      []int
	Adjacency  map[int][]Edge
	IsDirected bool
}

type graphConfig struct {
	directed bool
	weighted bool
	density  float64
}

type Option func(*graphConfig)

func WithDirected() Option         { return func(c *graphConfig) { c.directed = true } }
func WithWeighted() Option         { return func(c *graphConfig) { c.weighted = true } }
func WithDensity(density float64) Option { return func(c *graphConfig) { c.density = density } }

func Random(size int, opts ...Option) *Graph {
	cfg := parseOpts(opts)
	raw := generator.Random(size)
	return buildGraph(raw, cfg, true)
}

func Seq(size int, opts ...Option) *Graph {
	cfg := parseOpts(opts)
	raw := generator.Seq(size)
	return buildGraph(raw, cfg, false)
}

// FromSlice builds a graph from a slice of integers, connecting them sequentially.
func FromSlice(nums []int, opts ...Option) *Graph {
	cfg := parseOpts(opts)
	return buildGraph(nums, cfg, false)
}

// SeqReverse generates a linear path graph with nodes in reverse order.
func SeqReverse(size int, opts ...Option) *Graph {
	cfg := parseOpts(opts)
	raw := generator.SeqReverse(size)
	return buildGraph(raw, cfg, false)
}

func parseOpts(opts []Option) *graphConfig {
	cfg := &graphConfig{density: 0.25}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func buildGraph(nodeValues []int, cfg *graphConfig, isRandom bool) *Graph {
	g := &Graph{
		Nodes:      nodeValues,
		Adjacency:  make(map[int][]Edge),
		IsDirected: cfg.directed,
	}

	for _, v := range nodeValues {
		g.Adjacency[v] = []Edge{}
	}

	if isRandom {
		for _, u := range nodeValues {
			for _, v := range nodeValues {
				if u != v && rand.Float64() < cfg.density {
					addEdge(g, u, v, cfg)
				}
			}
		}
	} else {
		for i := 0; i < len(nodeValues)-1; i++ {
			addEdge(g, nodeValues[i], nodeValues[i+1], cfg)
		}
	}
	return g
}

func addEdge(g *Graph, u, v int, cfg *graphConfig) {
	weight := 1
	if cfg.weighted {
		weight = rand.Intn(generator.MaxValue) + 1
	}
	g.Adjacency[u] = append(g.Adjacency[u], Edge{To: v, Weight: weight})
	if !cfg.directed {
		g.Adjacency[v] = append(g.Adjacency[v], Edge{To: u, Weight: weight})
	}
}