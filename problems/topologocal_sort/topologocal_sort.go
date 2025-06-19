// 3. Topological Sort (Directed Acyclic Graph – DAG)

// Problem Statement
// Given a directed graph, return a valid topological ordering of its nodes (if one exists).

// Input: A directed graph (could be disconnected).
// Output: An array/slice of nodes, ordered such that for every edge u → v, node u appears before v in the ordering.
// If the graph has a cycle, no valid topological ordering exists.
// Example 1:
// Nodes: A, B, C, D
// Edges: (A → B), (B → C), (A → C), (C → D)

// Possible valid orderings:

// [A, B, C, D]
// [A, B, D, C] (not valid, because C→D, so C must be before D)

package main

import "fmt"

type AdjListGraph[T comparable] struct {
	Nodes   map[T]struct{}
	AdjList map[T]map[T]struct{}
}

// AddEdge adds an undirected edge between nodes a and b.
func (g *AdjListGraph[T]) AddEdge(a, b T) {
	if _, ok := g.Nodes[a]; !ok {
		g.Nodes[a] = struct{}{}
		g.AdjList[a] = map[T]struct{}{}
	}
	if _, ok := g.Nodes[b]; !ok {
		g.Nodes[b] = struct{}{}
		g.AdjList[b] = map[T]struct{}{}
	}
	g.AdjList[a][b] = struct{}{}
}

func (g *AdjListGraph[T]) TopologicalSort() []T {
	visited := map[T]struct{}{}
	order := []T{}
	recStack := map[T]struct{}{}
	hasCycle := false
	var dfs func(node T)
	dfs = func(node T) {
		if hasCycle {
			return
		}
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			recStack[node] = struct{}{}
			for k := range g.AdjList[node] {
				if _, done := visited[k]; !done {
					dfs(k)
				} else if _, done := recStack[k]; done {
					hasCycle = true
				}
			}
			order = append(order, node)
		}
		delete(recStack, node)
	}
	for k := range g.Nodes {
		if _, ok := visited[k]; !ok {
			dfs(k)
		}
	}
	if hasCycle {
		return []T{}
	}
	if len(order) == len(g.Nodes) {
		clondedOrder := []T{}
		for i := len(order) - 1; i >= 0; i-- {
			clondedOrder = append(clondedOrder, order[i])
		}
		return clondedOrder
	} else {
		return []T{}
	}
}

func main() {
	// Example 1: DAG
	g := &AdjListGraph[string]{
		Nodes:   make(map[string]struct{}),
		AdjList: make(map[string]map[string]struct{}),
	}
	// A → B → C
	//  \         ↘
	//   → D → E → F
	g.AddEdge("A", "B")
	g.AddEdge("A", "D")
	g.AddEdge("B", "C")
	g.AddEdge("D", "E")
	g.AddEdge("E", "F")
	g.AddEdge("C", "F")

	order := g.TopologicalSort()
	fmt.Printf("Topological Sort (should be valid): %v\n", order)

	// Example 2: Graph with a cycle
	g2 := &AdjListGraph[string]{
		Nodes:   make(map[string]struct{}),
		AdjList: make(map[string]map[string]struct{}),
	}
	// A → B → C → A (cycle)
	g2.AddEdge("A", "B")
	g2.AddEdge("B", "C")
	g2.AddEdge("C", "A")

	order2 := g2.TopologicalSort()
	fmt.Printf("Topological Sort with cycle (should be empty): %v\n", order2)
}
