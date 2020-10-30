package main

import (
	"fmt"
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
	nodes []node
	edges []edge
}

var infinity = math.MaxInt32

func bellmanFord(s string, g graph) map[string]int {
	d := make(map[string]int, len(g.nodes))

	for _, e := range g.nodes {
		d[e.key] = infinity
	}
	d[s] = 0
	updated := true

	for updated {
		updated = false
		for _, e := range g.edges {
			v := e.to
			u := e.from
			c := e.cost

			if d[u.key]+c < d[v.key] {
				d[v.key] = d[u.key] + c
				updated = true
			}
		}
		if !updated {
			break
		}
	}
	return d
}

func newEdge(from node, to node, cost int) edge {
	e := edge{from: from, to: to, cost: cost}
	return e
}

func main() {
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

	d := bellmanFord("a", g)
	for k, v := range d {
		fmt.Printf("key = %+v\n", k)
		fmt.Printf("cost = %+v\n", v)
		fmt.Println(strings.Repeat("-", 10))
	}
}
