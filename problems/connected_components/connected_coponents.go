// 1. Number of Connected Components (Undirected Graph)

// Problem Statement
// Given an undirected graph (nodes and edges), count how many connected components it contains.

// Definitions
// Undirected Graph: Edges have no direction (A-B is the same as B-A).
// Connected Component: A set of nodes such that every node in the set can reach every other node in the set via some path, and which is maximal (you can't add another node from the graph without breaking connectivity).
// Example
// Nodes:
// A, B, C, D, E, F

// Edges:
// (A-B), (B-C), (D-E)

// Visualization:

// A---B---C    D---E    F
// {A,B,C} are all connected to each other.
// {D,E} are connected to each other.
// {F} is isolated (not connected to anyone).
// So there are 3 components:

// {A,B,C}
// {D,E}
// {F}

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
	g.AdjList[b][a] = struct{}{}
}

func (g *AdjListGraph[T]) ConnectedComponents() int {
	visited := map[T]struct{}{}

	var dfs func(node T, order []T)
	dfs = func(node T, order []T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			order = append(order, node)
			for k, _ := range g.AdjList[node] {
				if _, done := visited[k]; !done {
					dfs(k, order)
				}
			}
		}
		//delete(visited, node)
	}
	count := 0
	for i := range g.Nodes {
		if _, ok := visited[i]; !ok {
			count++
			dfs(i, []T{})
		}

	}
	return count
}

func main() {
	// Example: Nodes: A, B, C, D, E, F; Edges: (A-B), (B-C), (D-E)
	g := &AdjListGraph[string]{
		Nodes:   map[string]struct{}{},
		AdjList: map[string]map[string]struct{}{},
	}
	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("D", "E")
	g.AddEdge("G", "A")
	g.AddEdge("D", "C")
	g.AddEdge("B", "G")
	g.Nodes["F"] = struct{}{} // F is an isolated node

	components := g.ConnectedComponents()
	fmt.Println("Number of connected components:", components) // Output: 3

	// You can test with other cases as well!
}
