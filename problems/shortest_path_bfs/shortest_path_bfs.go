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

func (g *AdjListGraph[T]) ShortestPath(source T, target T) []T {
	visited := map[T]struct{}{}
	queue := []T{source}
	parents := map[T]T{} //child => parent
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if _, done := visited[curr]; !done {
			visited[curr] = struct{}{}
			if curr == target {
				path := []T{}
				for i := target; ; i = parents[i] {
					path = append(path, i)
					if i == source {
						clonedPath := []T{}
						for k := len(path) - 1; k >= 0; k-- {
							clonedPath = append(clonedPath, path[k])
						}
						return clonedPath
					}
				}
			} else {
				for k := range g.AdjList[curr] {
					if _, done := visited[k]; !done {
						parents[k] = curr
						queue = append(queue, k)
					}
				}
			}

		}
	}
	return []T{}
}

func main() {
	g := &AdjListGraph[string]{
		Nodes:   make(map[string]struct{}),
		AdjList: make(map[string]map[string]struct{}),
	}

	// Build the graph:
	// A - B - C - D
	//           |
	//           E
	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")

	// Test 1: Shortest path from A to E
	path := g.ShortestPath("A", "E")
	fmt.Printf("Shortest path from A to E: %v\n", path)

	// Test 2: Shortest path from A to D
	path = g.ShortestPath("A", "D")
	fmt.Printf("Shortest path from A to D: %v\n", path)

	// Test 3: Shortest path from E to A (should work in undirected)
	path = g.ShortestPath("E", "A")
	fmt.Printf("Shortest path from E to A: %v\n", path)

	// Test 4: Path that does not exist
	path = g.ShortestPath("A", "Z")
	fmt.Printf("Shortest path from A to Z: %v\n", path)
}
