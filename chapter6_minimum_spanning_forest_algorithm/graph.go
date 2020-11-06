package graph

import "sort"

type Graph struct {
	edges     int32
	vertices  int32
	adjacency map[string]map[string]int32
}

func (g *Graph) Vertices() []string {
	vertices := make([]string, 0)
	for k := range g.adjacency {
		vertices = append(vertices, k)
	}
	return vertices
}

func (g *Graph) Adjacency(v string) []string {
	adj := make([]string, 0)
	if g.adjacency == nil {
		return adj
	}

	for k := range g.adjacency[v] {
		adj = append(adj, k)
	}

	return adj
}

func (g *Graph) AddEdge(v, w string, cost int32) {
	if g.adjacency == nil {
		g.adjacency = make(map[string]map[string]int32)
	}
	if g.adjacency[v] == nil {
		g.adjacency[v] = make(map[string]int32)
	}
	g.adjacency[v][w] = cost
}

func (g *Graph) SortAdjacency(v string) []string {
	if g.adjacency == nil {
		return []string{}
	}
	if len(g.adjacency[v]) <= 1 {
		return []string{}
	}
	type kv struct {
		Key   string
		Value int32
	}
	var ss []kv
	for k, v := range g.adjacency[v] {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})
	ranked := make([]string, len(g.adjacency[v]))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
