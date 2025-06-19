//2. Detect Cycle in an Undirected Graph

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

func (g *AdjListGraph[T]) UndirectedCycleDetection() bool {
	visited := map[T]struct{}{}

	var dfs func(node T, parent T) bool
	dfs = func(node T, parent T) bool {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			for k, _ := range g.AdjList[node] {
				if _, done := visited[k]; !done {
					if dfs(k, node) {
						return true
					}
				} else if k != parent {
					return true
				}
			}
		}
		return false
	}
	for k, _ := range g.Nodes {
		var zero T
		if dfs(k, zero) {
			return true
		}
	}
	return false

}

func main() {
	// Graph 1: Contains a cycle (A-B-C-A)
	g := &AdjListGraph[string]{
		Nodes:   map[string]struct{}{},
		AdjList: map[string]map[string]struct{}{},
	}
	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A")
	fmt.Println("Graph 1 has cycle:", g.UndirectedCycleDetection()) // Output: true

	// Graph 2: No cycle (X-Y-Z)
	g2 := &AdjListGraph[string]{
		Nodes:   map[string]struct{}{},
		AdjList: map[string]map[string]struct{}{},
	}
	g2.AddEdge("X", "Y")
	g2.AddEdge("Y", "Z")
	fmt.Println("Graph 2 has cycle:", g2.UndirectedCycleDetection()) // Output: false

	// Graph 3: Disconnected with and without cycle
	g3 := &AdjListGraph[string]{
		Nodes:   map[string]struct{}{},
		AdjList: map[string]map[string]struct{}{},
	}
	g3.AddEdge("M", "N")
	g3.AddEdge("N", "O")
	g3.AddEdge("O", "M")                                             // cycle
	g3.AddEdge("P", "Q")                                             // acyclic component
	fmt.Println("Graph 3 has cycle:", g3.UndirectedCycleDetection()) // Output: true
}
