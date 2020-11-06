package main

import (
	"fmt"

	graph "github.com/anyy/data-structures-and-algorithms/chapter6_minimum_spanning_forest_algorithm"
)

type bestFirstSearch struct {
	marked map[string]bool
}

func (b *bestFirstSearch) init(g graph.Graph) {
	if b.marked == nil {
		b.marked = make(map[string]bool)
	}
	for _, v := range g.Vertices() {
		b.marked[v] = false
	}
}

func (b *bestFirstSearch) BestFirstSearch(start string, g graph.Graph) {
	b.marked[start] = true
	fmt.Printf("marked = %+v\n", start)

	for _, v := range g.SortAdjacency(start) {
		if !b.marked[v] {
			b.BestFirstSearch(v, g)
		}
	}
}

func main() {
	g := new(graph.Graph)
	g.AddEdge("a", "c", 1)
	g.AddEdge("c", "a", 1)

	g.AddEdge("a", "b", 4)
	g.AddEdge("b", "a", 4)

	g.AddEdge("b", "c", 2)
	g.AddEdge("c", "b", 2)

	g.AddEdge("b", "d", 3)
	g.AddEdge("d", "b", 3)

	g.AddEdge("d", "e", 1)
	g.AddEdge("e", "d", 1)

	g.AddEdge("c", "d", 1)
	g.AddEdge("d", "c", 1)

	g.AddEdge("e", "f", 2)
	g.AddEdge("f", "e", 2)

	g.AddEdge("c", "f", 2)
	g.AddEdge("f", "c", 2)

	g.AddEdge("c", "e", 3)
	g.AddEdge("e", "c", 3)

	b := new(bestFirstSearch)
	b.init(*g)
	b.BestFirstSearch("a", *g)
}
