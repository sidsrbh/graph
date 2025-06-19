// 6. Articulation Points (Cut Vertices) in an Undirected Graph

// Problem Statement:
// Given an undirected graph, find all the articulation points.
// An articulation point is a node that, if removed (along with its edges), increases the number of connected components in the graph (i.e., disconnects the graph).

// Example
// Graph:

// A - B - C
//     |
//     D
// Articulation Point: B (removing B disconnects A from C and D)
// Explanation
// Why is this important?
// Finding articulation points helps in network reliability, identifying single points of failure, etc.

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

func (g *AdjListGraph[T]) ArticulationPoints() []T {
	visited := map[T]struct{}{}
	desc := map[T]int{}
	low := map[T]int{}
	points := []T{}
	parents := map[T]T{}
	added := map[T]struct{}{}
	time := 1 //child => parent

	var dfs func(node T)
	dfs = func(node T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			desc[node] = time
			low[node] = time
			time++
			children := 0
			for i := range g.AdjList[node] {
				if _, done := visited[i]; !done {
					parents[i] = node
					children++
					dfs(i)
					low[node] = min(low[node], low[i])
					if _, ok := parents[node]; ok {
						if low[i] >= desc[node] {
							if _, done := added[node]; !done {
								points = append(points, node)
								added[node] = struct{}{}
							}
						}
					}
				} else {
					if _, ok := parents[node]; !ok || i != parents[node] {
						low[node] = min(low[node], desc[i])
					}

				}
			}
			if _, ok := parents[node]; !ok && children > 1 {
				if _, done := added[node]; !done {
					points = append(points, node)
					added[node] = struct{}{}
				}
			}
		}
	}
	for i := range g.Nodes {
		if _, ok := visited[i]; !ok {
			dfs(i)
		}
	}
	return points
}

func main() {
	// Create the graph
	g := &AdjListGraph[string]{
		Nodes:   make(map[string]struct{}),
		AdjList: make(map[string]map[string]struct{}),
	}

	// Build a test graph (mix of line, star, and cycle):
	// Line:   A - B - C
	// Star:         |
	//             (D)
	// Cycle:        E - F - G - E

	edges := [][2]string{
		{"A", "B"},
		{"B", "C"},
		{"C", "D"},
		{"E", "F"},
		{"F", "G"},
		{"G", "E"},
	}

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	// You can add a disconnected node for further testing:
	g.Nodes["X"] = struct{}{}
	g.AdjList["X"] = map[string]struct{}{}

	// Print the articulation points
	points := g.ArticulationPoints()
	fmt.Println("Articulation Points:", points)
}
