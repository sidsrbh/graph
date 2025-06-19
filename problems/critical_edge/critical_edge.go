// 6. Detect a Bridge (Critical Edge) in an Undirected Graph

// Problem Statement:
// Given an undirected graph, find all bridges (edges whose removal increases the number of connected components).

// Why important?

// Bridges represent single points of failure in networks.
// Common interview question, tests DFS, low-link values, and graph connectivity.
// Explanation:
// A bridge (or cut-edge) is an edge that, if removed, disconnects the graph.
// For example, in:

// A - B - C
//     |
//     D
// The edge B-C is a bridge. Remove it, and C becomes isolated.

package main

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

func (g *AdjListGraph[T]) CriticalEdge() [][2]T {
	visited := map[T]struct{}{}
	disc := map[T]int{}
	low := map[T]int{}
	bridges := [][2]T{}
	parent := map[T]T{}
	time := 1
	var dfs func(node T)

	dfs = func(node T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			disc[node] = time
			low[node] = time
			time++
			for i := range g.AdjList[node] {
				if _, done := visited[i]; !done {
					parent[i] = node
					dfs(i)
					low[node] = min(low[node], low[i])
					if low[i] > disc[node] {
						bridges = append(bridges, [2]T{node, i})
					}
				} else {
					if i != parent[node] {
						low[node] = min(low[node], disc[i])
					}
				}
			}
		}
	}
	for i := range g.Nodes {
		if _, ok := visited[i]; !ok {
			dfs(i)
		}
	}
	return bridges
}
