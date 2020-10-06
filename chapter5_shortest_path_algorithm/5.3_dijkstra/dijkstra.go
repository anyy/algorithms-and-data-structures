package main

import (
	"fmt"
	"math"
)

type edge struct {
	from *node
	to   *node
	cost int
}

type node struct {
	key string
}

type graph struct {
	nodes     []*node
	edges     []*edge
	costTable map[*node]int
}

func newEdge(from *node, to *node, cost int) *edge {
	e := edge{from: from, to: to, cost: cost}
	return &e
}

func main() {
	graph := graph{}

	nodeA := &node{key: "a"}
	nodeB := &node{key: "b"}
	nodeC := &node{key: "c"}
	nodeD := &node{key: "d"}
	nodeE := &node{key: "e"}
	nodeF := &node{key: "f"}
	graph.nodes = append(graph.nodes, nodeA, nodeB, nodeC, nodeD, nodeE, nodeF)

	graph.addEdge(newEdge(nodeA, nodeB, 4), true)
	graph.addEdge(newEdge(nodeA, nodeC, 1), true)
	graph.addEdge(newEdge(nodeB, nodeD, 3), true)
	graph.addEdge(newEdge(nodeC, nodeD, 1), true)
	graph.addEdge(newEdge(nodeC, nodeE, 3), true)
	graph.addEdge(newEdge(nodeC, nodeF, 2), true)
	graph.addEdge(newEdge(nodeD, nodeE, 1), true)
	graph.addEdge(newEdge(nodeE, nodeF, 1), true)
	graph.addEdge(newEdge(nodeB, nodeC, 2), true)

	graph.shortestPath(nodeA)
	for k, c := range graph.costTable {
		fmt.Printf("from = %+v, to = %+v, cost = %+v\n", nodeA.key, k.key, c)
	}
}

func (g *graph) addEdge(edge *edge, directed bool) {
	g.edges = append(g.edges, edge)
	if directed {
		e := newEdge(edge.to, edge.from, edge.cost)
		g.edges = append(g.edges, e)
	}
}

// dijkstra algorithm
func (g *graph) shortestPath(s *node) {
	// initialize cost table
	g.initCostTable(s)
	// U is a set of vertices with a shortest path from the start point
	U := []*node{s}
	for len(U) != len(g.nodes) {
		neighborEdges := g.neighborEdge(U)

		min := neighborEdges[0]
		for _, ne := range neighborEdges {
			c := g.costTable[ne.from] + ne.cost
			if c < g.costTable[ne.to] {
				g.costTable[ne.to] = c
			}
			min = g.min(min, ne)
		}

		U = append(U, min.to)
	}
}

// initCostTable initializes a cost table
func (g *graph) initCostTable(start *node) {
	g.costTable = make(map[*node]int)
	for _, n := range g.nodes {
		g.costTable[n] = math.MaxInt32
	}
	g.costTable[start] = 0
}

// neighborEdge returns edges nearby U
func (g *graph) neighborEdge(U []*node) []*edge {
	var neighbors []*edge
	for _, u := range U {
		for _, e := range g.edges {
			if e.from.key == u.key && !visited(e.to, U) {
				neighbors = append(neighbors, e)
			}
		}
	}
	return neighbors
}

// visited returns if the node has already been visited
func visited(x *node, y []*node) bool {
	for _, n := range y {
		if x.key == n.key {
			return true
		}
	}
	return false
}

// min returns a edge that has minimum cost between two edges
func (g *graph) min(x, y *edge) *edge {
	if (g.costTable[y.from] + y.cost) < (g.costTable[x.from] + x.cost) {
		return y
	}
	return x
}
