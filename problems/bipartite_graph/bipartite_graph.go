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

func (g *AdjListGraph[T]) IsBipartite() bool {
	visited := map[T]struct{}{}
	colors := map[T]bool{}
	var dfs func(node T, red bool)
	isBip := true
	dfs = func(node T, red bool) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			colors[node] = red

			for i := range g.AdjList[node] {
				if _, done := visited[i]; !done {
					dfs(i, !red)
				} else {
					if colors[i] == red {
						isBip = false
					}

				}
			}

		}
	}
	for i := range g.Nodes {
		if _, ok := visited[i]; !ok {
			dfs(i, true)
		}
	}
	return isBip
}
