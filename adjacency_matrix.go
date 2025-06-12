package graph

func (g *Graph[T]) AddNodeAdjMatrix(node T) {
	if _, exists := g.nodes[node]; exists {
		return
	}
	g.nodes[node] = struct{}{}
	g.indexToNodes = append(g.indexToNodes, node)
	g.nodesToIndex[node] = len(g.indexToNodes) - 1
	for i := 0; i < len(g.indexToNodes)-1; i++ {
		g.adjMatrix[i] = append(g.adjMatrix[i], false)
	}
	n := len(g.indexToNodes)
	newRow := make([]bool, n)
	g.adjMatrix = append(g.adjMatrix, newRow)
}

func (g *Graph[T]) RemoveNodeAdjMatrix(node T) {
	if _, exists := g.nodes[node]; !exists {
		return
	}
	length := len(g.indexToNodes)
	delete(g.nodes, node)
	index := g.nodesToIndex[node]
	delete(g.nodesToIndex, node)
	//remove from index to nodes
	for i := index; i < length-1; i++ {
		g.indexToNodes[i] = g.indexToNodes[i+1]
	}
	g.indexToNodes = g.indexToNodes[:length-1]
	for i, v := range g.indexToNodes {
		g.nodesToIndex[v] = i
	}
	//remove row and columns
	for k := 0; k < length; k++ {
		for i := index; i < length-1; i++ {
			g.adjMatrix[k][i] = g.adjMatrix[k][i+1]
		}
		g.adjMatrix[k] = g.adjMatrix[k][:length-1]
	}
	g.adjMatrix = g.adjMatrix[:length-1]

}

func (g *Graph[T]) AddEdgeAdjMatrix(from T, to T) {
	if _, ok := g.nodes[from]; !ok {
		g.AddNode(from)
	}
	if _, ok := g.nodes[to]; !ok {
		g.AddNode(to)
	}
	g.adjMatrix[g.nodesToIndex[from]][g.nodesToIndex[to]] = true
	if g.graphType == Undirected {
		g.adjMatrix[g.nodesToIndex[to]][g.nodesToIndex[from]] = true
	}

}

func (g *Graph[T]) RemoveEdgeAdjMatrix(from T, to T) {
	if _, ok := g.nodes[from]; !ok {
		return
	}
	if _, ok := g.nodes[to]; !ok {
		return
	}
	g.adjMatrix[g.nodesToIndex[from]][g.nodesToIndex[to]] = false
	if g.graphType == Undirected {
		g.adjMatrix[g.nodesToIndex[to]][g.nodesToIndex[from]] = false
	}
}
func (g *Graph[T]) HasEdgeAdjMatrix(from T, to T) bool {
	if _, ok := g.nodes[from]; !ok {
		return false
	}
	if _, ok := g.nodes[to]; !ok {
		return false
	}
	if g.adjMatrix[g.nodesToIndex[from]][g.nodesToIndex[to]] {
		if g.graphType == Directed {
			return true
		} else {
			if g.adjMatrix[g.nodesToIndex[from]][g.nodesToIndex[to]] {
				return true
			}
		}
	}
	return false
}

func (g *Graph[T]) NeighboursAdjMatrix(node T) []T {
	if !g.HasNode(node) {
		return make([]T, 0)
	}
	nodes := []T{}
	matrixLength := len(g.nodes)
	for i := 0; i < matrixLength; i++ {
		if g.adjMatrix[g.nodesToIndex[node]][i] {
			if g.graphType == Undirected {
				if g.adjMatrix[i][g.nodesToIndex[node]] {
					nodes = append(nodes, g.indexToNodes[i])
				}
			} else {
				nodes = append(nodes, g.indexToNodes[i])
			}
		}
	}
	return nodes
}

func (g *Graph[T]) EdgesAdjMatrix() [][2]T {
	edges := make([][2]T, 0)
	for i := 0; i < len(g.nodesToIndex); i++ {
		for k := 0; k < len(g.nodesToIndex); k++ {
			if g.adjMatrix[i][k] {
				if g.graphType == Directed {
					edges = append(edges, [2]T{g.indexToNodes[i], g.indexToNodes[k]})
				} else {
					if i < k {
						edges = append(edges, [2]T{g.indexToNodes[i], g.indexToNodes[k]})
					}
				}

			}
		}
	}
	return edges
}

func (g *Graph[T]) OutDegreeAdjMatrix(node T) int {
	count := 0
	index := g.nodesToIndex[node]
	for i := 0; i < len(g.indexToNodes); i++ {
		if g.adjMatrix[index][i] {
			count++
		}
	}
	return count
}

func (g *Graph[T]) InDegreeAdjMatrix(node T) int {
	count := 0
	index := g.nodesToIndex[node]
	for i := 0; i < len(g.indexToNodes); i++ {
		if g.adjMatrix[i][index] {
			count++
		}
	}
	return count
}

func (g *Graph[T]) BFSAdjMatrix(start T) []T {
	visited := map[T]struct{}{}
	queue := []T{start}
	order := []T{start}
	visited[start] = struct{}{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for i := 0; i < len(g.indexToNodes); i++ {
			if g.HasEdge(curr, g.indexToNodes[i]) {
				if _, ok := visited[g.indexToNodes[i]]; !ok {
					visited[g.indexToNodes[i]] = struct{}{}
					order = append(order, g.indexToNodes[i])
					queue = append(queue, g.indexToNodes[i])
				}
			}
		}
	}
	return order
}

func (g *Graph[T]) DFSIterativeAdjMatrix(start T) []T {
	order := []T{}
	visited := map[T]struct{}{}
	stack := []T{start}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[curr]; !ok {
			visited[curr] = struct{}{}
			order = append(order, curr)

		}
		for i := len(g.indexToNodes) - 1; i >= 0; i-- {
			if g.HasEdge(curr, g.indexToNodes[i]) {
				neighbour := g.indexToNodes[i]
				if _, done := visited[neighbour]; !done {
					stack = append(stack, neighbour)
				}

			}
		}
	}
	return order
}

func (g *Graph[T]) DFSRecursiveAdjMatrix(start T) []T {
	order := []T{}
	visited := map[T]struct{}{}

	var dfs func(node T)
	dfs = func(node T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			order = append(order, node)
			for i := 0; i < len(g.indexToNodes); i++ {
				if g.HasEdge(node, g.indexToNodes[i]) {
					nbr := g.indexToNodes[i]
					if nbr != node {
						if _, done := visited[nbr]; !done {
							dfs(nbr)
						}
					}
				}
			}
		}
	}
	return order
}

func (g *Graph[T]) RecursiveDFSAllPathFindingAdjMatrix(source T, target T) [][]T {
	orders := [][]T{}
	visited := make(map[T]struct{})

	var dfs func(node T, order []T)
	dfs = func(node T, order []T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			order = append(order, node)
			if node == target {
				temp := append([]T{}, order...)
				orders = append(orders, temp)
			} else {
				for i := 0; i < len(g.indexToNodes); i++ {
					if g.HasEdge(node, g.indexToNodes[i]) {
						nbr := g.indexToNodes[i]
						if _, done := visited[nbr]; !done {
							dfs(nbr, order)
						}

					}
				}
			}
		}
		delete(visited, node)
	}
	dfs(source, []T{})
	return orders
}

func (g *Graph[T]) RecursiveDFSAnyPathFindingAdjMatrix(source T, target T) []T {
	visited := make(map[T]struct{})
	returnOrder := []T{}

	var dfs func(node T, order []T) bool
	dfs = func(node T, order []T) bool {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			order = append(order, node)
			if node == target {
				returnOrder = append([]T{}, order...)
				return true
			} else {
				for i := 0; i < len(g.indexToNodes); i++ {
					if g.HasEdge(node, g.indexToNodes[i]) {
						nbr := g.indexToNodes[i]
						if _, done := visited[nbr]; !done {
							if dfs(nbr, order) {
								return true
							}
						}

					}
				}
			}
			delete(visited, node)
		}
		return false
	}
	if dfs(source, []T{}) {
		return returnOrder
	} else {
		return []T{}
	}

}

func (g *Graph[T]) DFSIterativeAllPathFindingAdjMatrix(source T, target T) [][]T {
	visiteds := []map[T]struct{}{map[T]struct{}{}}
	stack := []T{source}
	orders := [][]T{[]T{}}
	returningOrders := [][]T{}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		order := orders[len(stack)-1]
		visited := visiteds[len(stack)-1]
		stack = stack[:len(stack)-1]
		orders = orders[:len(orders)-1]
		visiteds = visiteds[:len(visiteds)-1]

		if _, ok := visited[curr]; !ok {
			if curr == target {
				clonedOrder := append([]T{}, order...)
				clonedOrder = append(clonedOrder, curr)
				returningOrders = append(returningOrders, clonedOrder)
			} else {
				visited[curr] = struct{}{}
				order = append(order, curr)
				for i := 0; i < len(g.indexToNodes); i++ {
					nbr := g.indexToNodes[i]
					if _, done := visited[nbr]; !done {
						stack = append(stack, nbr)
						clonedOrder := append([]T{}, order...)
						orders = append(orders, clonedOrder)
						clonedVisited := map[T]struct{}{}
						for k, v := range visited {
							clonedVisited[k] = v
						}

						visiteds = append(visiteds, clonedVisited)
					}

				}
			}

		}

	}
	return returningOrders

}

func (g *Graph[T]) DFSIterativeAnyPathFindingAdjMatrix(source T, target T) []T {
	visiteds := []map[T]struct{}{map[T]struct{}{}}
	stack := []T{source}
	orders := [][]T{[]T{}}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		order := orders[len(stack)-1]
		visited := visiteds[len(stack)-1]
		stack = stack[:len(stack)-1]
		orders = orders[:len(orders)-1]
		visiteds = visiteds[:len(visiteds)-1]

		if _, ok := visited[curr]; !ok {
			if curr == target {
				clonedOrder := append([]T{}, order...)
				clonedOrder = append(clonedOrder, curr)
				return clonedOrder
			} else {
				visited[curr] = struct{}{}
				order = append(order, curr)
				for i := 0; i < len(g.indexToNodes); i++ {
					nbr := g.indexToNodes[i]
					if _, done := visited[nbr]; !done {
						stack = append(stack, nbr)
						clonedOrder := append([]T{}, order...)
						orders = append(orders, clonedOrder)
						clonedVisited := map[T]struct{}{}
						for k, v := range visited {
							clonedVisited[k] = v
						}

						visiteds = append(visiteds, clonedVisited)
					}

				}
			}

		}

	}
	return []T{}

}
func (g *Graph[T]) BFSShortestPathAdjMatrix(source T, target T) []T {
	queue := []T{source}
	visited := map[T]struct{}{}
	parents := map[T]T{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if _, ok := visited[curr]; !ok {
			visited[curr] = struct{}{}
			if curr == target {
				path := []T{}
				for node := target; ; node = parents[node] {
					path = append(path, node)
					if node == source {
						return path
					}
				}
			} else {
				for i := 0; i < len(g.indexToNodes); i++ {
					if g.HasEdge(curr, g.indexToNodes[i]) {
						if _, done := visited[g.indexToNodes[i]]; !done {
							queue = append(queue, g.indexToNodes[i])
							parents[g.indexToNodes[i]] = curr
						}
					}
				}
			}
		}
	}
	return []T{}
}

func (g *Graph[T]) HasCycleDirectedAdjMatrix() bool {
	for i := 0; i < len(g.indexToNodes); i++ {
		curr := g.indexToNodes[i]
		visited := map[T]struct{}{}
		recStack := map[T]struct{}{}
		var dfs func(source T) bool
		dfs = func(source T) bool {
			if _, ok := visited[source]; !ok {
				visited[source] = struct{}{}
				recStack[source] = struct{}{}
				for i := 0; i < len(g.indexToNodes); i++ {
					if g.HasEdge(source, g.indexToNodes[i]) {
						if _, done := visited[g.indexToNodes[i]]; !done {
							if dfs(g.indexToNodes[i]) {
								return true
							}
						} else if _, inStack := recStack[g.indexToNodes[i]]; inStack {
							return true
						}
					}
				}

			}
			delete(recStack, source)
			return false
		}
		if dfs(curr) {
			return true
		}
	}
	return false

}

func (g *Graph[T]) HasCycleUndirectedAdjMatrix() bool {
	visited := map[T]struct{}{}
	var dfs func(source T, parent T) bool
	dfs = func(source T, parent T) bool {
		if _, ok := visited[source]; !ok {
			visited[source] = struct{}{}
			for i := 0; i < len(g.indexToNodes); i++ {
				if g.HasEdge(source, g.indexToNodes[i]) {
					if _, done := visited[g.indexToNodes[i]]; !done {
						if dfs(g.indexToNodes[i], source) {
							return true
						}
					} else if g.graphType == Undirected {
						if g.indexToNodes[i] != parent {
							return true
						}
					}
				}
			}

		}
		return false
	}
	for i := 0; i < len(g.indexToNodes); i++ {
		curr := g.indexToNodes[i]
		var zero T
		if dfs(curr, zero) {
			return true
		}
	}
	return false

}
