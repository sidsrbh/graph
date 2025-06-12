package graph

func (g *Graph[T]) AddNodeAdjList(node T) {
	if _, exists := g.nodes[node]; exists {
		return
	}
	g.nodes[node] = struct{}{}
	g.adjList[node] = make(map[T]struct{})
}

func (g *Graph[T]) RemoveNodeAdjList(node T) {
	if _, exists := g.nodes[node]; !exists {
		return
	}
	delete(g.nodes, node)
	delete(g.adjList, node)
	for key, _ := range g.nodes {
		delete(g.adjList[key], node)
	}
}

func (g *Graph[T]) AddEdgeAdjList(from T, to T) {
	g.AddNode(from)
	g.AddNode(to)
	g.adjList[from][to] = struct{}{}
	if g.graphType == Undirected {
		g.adjList[to][from] = struct{}{}
	}
}

func (g *Graph[T]) RemoveEdgeAdjList(from T, to T) {
	delete(g.adjList[from], to)
	if g.graphType == Undirected {
		delete(g.adjList[to], from)
	}
}

func (g *Graph[T]) HasEdgeAdjList(from T, to T) bool {
	possibleTo := g.adjList[from]
	_, fte := possibleTo[to]

	possibleFrom := g.adjList[to]
	if g.graphType == Undirected {
		_, tfe := possibleFrom[from]
		if fte && tfe {
			return true
		}
	} else {
		if fte {
			return true
		}
	}
	return false
}

func (g *Graph[T]) NeighboursAdjList(node T) []T {
	if !g.HasNode(node) {
		return make([]T, 0)
	}
	if _, exists := g.adjList[node]; !exists {
		return make([]T, 0)
	} else {
		elem := make([]T, 0, len(g.adjList[node]))
		for key, _ := range g.adjList[node] {
			elem = append(elem, key)
		}
		return elem
	}
}

func (g *Graph[T]) EdgesAdjList() [][2]T {
	edges := make([][2]T, 0)
	seen := make(map[[2]T]struct{})
	for from, _ := range g.adjList {
		for to, _ := range g.adjList[from] {
			elem := [2]T{from, to}
			if g.graphType == Undirected {
				rev := [2]T{to, from}
				if _, ok := seen[rev]; ok {
					continue
				}
			}
			edges = append(edges, elem)
			seen[elem] = struct{}{}
		}
	}
	return edges
}

func (g *Graph[T]) OutDegreeAdjList(node T) int {
	return len(g.adjList[node])
}

func (g *Graph[T]) InDegreeAdjList(node T) int {
	count := 0
	for key, _ := range g.adjList {
		if _, exists := g.adjList[key][node]; exists {
			count++
		}
	}
	return count
}

func (g *Graph[T]) BFSAdjList(start T) []T {
	order := []T{}
	visited := make(map[T]struct{})
	visited[start] = struct{}{}
	queue := []T{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		order = append(order, curr)
		for nbr, _ := range g.adjList[curr] {
			if _, ok := visited[nbr]; !ok {
				queue = append(queue, nbr)
				visited[nbr] = struct{}{}
			}
		}
	}
	return order
}

func (g *Graph[T]) DFSRecursiveAdjList(start T) []T {
	order := []T{}
	visited := make(map[T]struct{})
	var dfs func(curr T)
	dfs = func(curr T) {
		visited[curr] = struct{}{}
		order = append(order, curr)
		for key := range g.adjList[curr] {
			if _, exists := visited[key]; !exists {
				dfs(key)
			}
		}
	}
	dfs(start)
	return order
}

func (g *Graph[T]) DFSIterativeAdjList(start T) []T {
	visited := make(map[T]struct{})
	stack := []T{start}
	order := []T{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		//Upto but not including index len(stack -1)
		stack = stack[:len(stack)-1]
		if _, done := visited[curr]; !done {
			visited[curr] = struct{}{}
			order = append(order, curr)
			for key, _ := range g.adjList[curr] {
				stack = append(stack, key)
			}
		}

	}

	return order
}

func (g *Graph[T]) RecursiveDFSAllPathFindingAdjList(source T, target T) [][]T {
	orders := [][]T{}
	visited := make(map[T]struct{})

	var dfs func(node T, order []T)

	dfs = func(node T, order []T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			order = append(order, node)
			if node == target {
				clonedOrder := make([]T, len(order))
				copy(clonedOrder, order)
				orders = append(orders, clonedOrder)
			} else {
				for curr := range g.adjList[node] {
					dfs(curr, order)
				}
			}

		}
		delete(visited, node)
	}
	dfs(source, []T{})
	return orders
}

func (g *Graph[T]) RecursiveDFSAnyPathFindingAdjList(source T, target T) []T {
	order := []T{}
	visited := make(map[T]struct{})

	var dfs func(node T) bool

	dfs = func(node T) bool {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			order = append(order, node)
			if node == target {
				return true
			} else {
				for curr := range g.adjList[node] {
					if dfs(curr) {
						return true
					}
				}
			}

		} else {
			return false
		}
		order = order[:len(order)-1]
		return false
	}
	if dfs(source) {
		return order
	}
	return []T{}
}

func (g *Graph[T]) DFSIterativeAllPathFindingAdjList(source T, target T) [][]T {
	orders := [][]T{}
	stack := []T{source}
	visiteds := []map[T]struct{}{map[T]struct{}{source: {}}}
	visitingOrders := [][]T{[]T{source}}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currOrder := visitingOrders[len(visitingOrders)-1]
		visitingOrders = visitingOrders[:len(visitingOrders)-1]
		currVisited := visiteds[len(visiteds)-1]
		visiteds = visiteds[:len(visiteds)-1]
		if _, ok := currVisited[curr]; !ok {
			if curr == target {
				temp := make([]T, len(currOrder))
				copy(temp, currOrder)
				orders = append(orders, temp)
			} else {
				for k := range g.adjList[curr] {
					if _, done := currVisited[k]; !done {
						stack = append(stack, k)
						newVisited := map[T]struct{}{}
						for k := range currVisited {
							newVisited[k] = struct{}{}
						}
						newVisited[k] = struct{}{}
						newOrder := make([]T, len(currOrder)+1)
						copy(newOrder, currOrder)
						newOrder[len(currOrder)] = k
						visiteds = append(visiteds, newVisited)
						visitingOrders = append(visitingOrders, newOrder)
					}

				}
			}
		}

	}
	return orders

}

func (g *Graph[T]) DFSIterativeAnyPathFindingAdjList(source T, target T) []T {
	stack := []T{source}
	visiteds := []map[T]struct{}{{source: struct{}{}}}
	visitingOrders := [][]T{[]T{source}}
	order := []T{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currVisited := visiteds[len(visiteds)-1]
		visiteds = visiteds[:len(visiteds)-1]
		currVisitingOrder := visitingOrders[len(visitingOrders)-1]

		if curr == target {
			return currVisitingOrder
		} else {
			for k := range g.adjList[curr] {
				stack = append(stack, k)
				newVisitingOrder := make([]T, len(currVisitingOrder)+1)
				copy(newVisitingOrder, currVisitingOrder)
				newVisitingOrder[len(currVisitingOrder)] = k
				visitingOrders = append(visitingOrders, newVisitingOrder)
				newVisited := make(map[T]struct{}, len(currVisited))
				for l := range currVisited {
					newVisited[l] = struct{}{}
				}
				newVisited[k] = struct{}{}
				visiteds = append(visiteds, newVisited)
			}
		}
	}
	return order
}

func (g *Graph[T]) BFSShortestPathAdjList(source T, target T) []T {
	queue := []T{source}
	visited := make(map[T]struct{})
	parent := map[T]T{} //child -> Parent
	visited[source] = struct{}{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if target == curr {
			path := []T{}
			for node := target; ; node = parent[node] {
				path = append([]T{node}, path...)
				if node == source {
					return path
				}
			}

		}
		for k := range g.adjList[curr] {
			if _, ok := visited[k]; !ok {
				parent[k] = curr
				visited[k] = struct{}{}
				queue = append(queue, k)
			}

		}
	}
	return []T{}
}

func (g *Graph[T]) HasCycleDirectedAdjList() bool {
	visited := make(map[T]struct{})
	recStack := make(map[T]struct{})
	var dfs func(source T) bool
	dfs = func(source T) bool {
		visited[source] = struct{}{}
		recStack[source] = struct{}{}
		for k := range g.adjList[source] {
			if _, ok := visited[k]; !ok {
				if dfs(k) {
					return true
				}
			} else if _, done := recStack[k]; done {
				return true
			}
		}
		delete(recStack, source)
		return false
	}
	for i := range g.nodes {
		if _, ok := visited[i]; !ok {
			if dfs(i) {
				return true
			}
		}

	}
	return false
}

func (g *Graph[T]) HasCycleUndirectedAdjList() bool {
	visited := make(map[T]struct{})
	var dfs func(source T, parent T) bool
	dfs = func(source T, parent T) bool {
		visited[source] = struct{}{}
		for k := range g.adjList[source] {
			if _, ok := visited[k]; !ok {
				if dfs(k, source) {
					return true
				}
			} else {
				if k != parent {
					return true
				}
			}
		}
		return false
	}
	for i := range g.nodes {
		if _, ok := visited[i]; !ok {
			if dfs(i, i) {
				return true
			}
		}

	}
	return false
}
