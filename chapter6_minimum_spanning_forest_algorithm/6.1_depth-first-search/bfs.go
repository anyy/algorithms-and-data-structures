package main

import (
	"fmt"

	graph "github.com/anyy/data-structures-and-algorithms/chapter6_minimum_spanning_forest_algorithm"
)

type depthFirstSearch struct {
	marked map[string]bool
}

func (d *depthFirstSearch) init(g graph.Graph) {
	if d.marked == nil {
		d.marked = make(map[string]bool)
	}
	for _, v := range g.Vertices() {
		d.marked[v] = false
	}
}

func (d *depthFirstSearch) DFS(start string, g graph.Graph) {
	d.marked[start] = true
	fmt.Printf("marked = %+v\n", start)

	for _, v := range g.Adjacency(start) {
		if !d.marked[v] {
			d.DFS(v, g)
		}
	}
}

func main() {
	g := new(graph.Graph)
	g.AddEdge("a", "b", 0)
	g.AddEdge("b", "c", 0)
	g.AddEdge("b", "d", 0)

	g.AddEdge("a", "e", 0)
	g.AddEdge("e", "f", 0)
	g.AddEdge("f", "g", 0)
	g.AddEdge("f", "h", 0)
	g.AddEdge("e", "i", 0)
	g.AddEdge("i", "j", 0)
	g.AddEdge("i", "k", 0)

	g.AddEdge("a", "l", 0)
	g.AddEdge("l", "m", 0)
	g.AddEdge("l", "n", 0)
	g.AddEdge("n", "o", 0)

	d := new(depthFirstSearch)
	d.init(*g)
	d.DFS("a", *g)
}
