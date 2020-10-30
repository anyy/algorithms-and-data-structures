package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
)

type edge struct {
	from node
	to   node
	cost int
}

type node struct {
	key string
}

type graph struct {
	nodes     []node
	edges     []edge
	distances map[string]int
}

var infinity = math.MaxInt32

func (g *graph) printDistance() {
	for k, v := range g.distances {
		fmt.Printf("node = %+v\n", k)
		fmt.Printf("cost = %+v\n", v)
		fmt.Println(strings.Repeat("-", 10))
	}
}

// O(|V|·|E|)
func bellmanFord(s string, g graph) (graph, error) {
	g.distances = make(map[string]int, len(g.nodes))

	for _, e := range g.nodes {
		g.distances[e.key] = infinity
	}
	g.distances[s] = 0

	// |V| − 1
	for i := 0; i < len(g.nodes); i++ {
		for _, e := range g.edges {
			v := e.to
			u := e.from
			c := e.cost

			if g.distances[u.key]+c < g.distances[v.key] {
				g.distances[v.key] = g.distances[u.key] + c
				if i >= len(g.nodes)-1 {
					return g, errors.New("the graph has negative cycle")
				}
			}
		}
	}
	return g, nil
}

func newEdge(from node, to node, cost int) edge {
	e := edge{from: from, to: to, cost: cost}
	return e
}

func main() {
	g := newGraph()
	g, err := bellmanFord("a", g)
	if err != nil {
		log.Fatal(err.Error())
	}
	g.printDistance()

	g = newGraphHasNegativeCycle()
	g, err = bellmanFord("a", g)
	if err != nil {
		log.Fatal(err.Error())
	}
	g.printDistance()
}

func newGraph() graph {
	nodeA := node{key: "a"}
	nodeB := node{key: "b"}
	nodeC := node{key: "c"}
	nodeD := node{key: "d"}
	nodeE := node{key: "e"}
	nodeF := node{key: "f"}

	g := graph{}
	g.edges = append(g.edges,
		newEdge(nodeA, nodeB, 4),
		newEdge(nodeA, nodeC, 3),
		newEdge(nodeB, nodeD, 1),
		newEdge(nodeC, nodeB, -2),
		newEdge(nodeC, nodeD, 2),
		newEdge(nodeC, nodeF, 2),
		newEdge(nodeC, nodeE, 3),
		newEdge(nodeD, nodeE, 1),
		newEdge(nodeE, nodeF, 2),
	)
	g.nodes = append(g.nodes, nodeA, nodeB, nodeC, nodeD, nodeE, nodeF)
	return g
}

func newGraphHasNegativeCycle() graph {
	nodeA := node{key: "a"}
	nodeB := node{key: "b"}
	nodeC := node{key: "c"}
	nodeD := node{key: "d"}
	nodeE := node{key: "e"}

	g := graph{}
	g.edges = append(g.edges,
		newEdge(nodeA, nodeB, 6),
		newEdge(nodeB, nodeC, -2),
		newEdge(nodeC, nodeA, -5),
		newEdge(nodeA, nodeD, 2),
		newEdge(nodeD, nodeC, 7),
		newEdge(nodeC, nodeE, 2),
		newEdge(nodeE, nodeD, 4),
	)
	g.nodes = append(g.nodes, nodeA, nodeB, nodeC, nodeD, nodeE)
	return g
}
