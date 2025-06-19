package graph

type GraphType int

const (
	Directed GraphType = iota
	Undirected
)

type RepresentationType int

const (
	AdjacencyList RepresentationType = iota
	AdjacencyMatrix
)

const INF = int(1e9)

type Graph[T comparable] struct {
	graphType GraphType
	repType   RepresentationType

	nodes map[T]struct{}

	//AdjacencyList
	adjList map[T]map[T]struct{}

	//Adjacency Matrix reated storage
	nodesToIndex map[T]int
	indexToNodes []T
	adjMatrix    [][]bool
}

type WeightedGraph[T comparable] struct {
	graphType GraphType
	repType   RepresentationType

	nodes map[T]struct{}

	//AdjacencyList
	adjList map[T]map[T]int

	//Adjacency Matrix reated storage
	nodesToIndex map[T]int
	indexToNodes []T
	adjMatrix    [][]int
}

func NewGraph[T comparable](graphType GraphType, repType RepresentationType) *Graph[T] {
	graph := &Graph[T]{
		graphType:    graphType,
		repType:      repType,
		nodes:        make(map[T]struct{}),
		adjList:      make(map[T]map[T]struct{}),
		nodesToIndex: make(map[T]int),
		indexToNodes: []T{},
		adjMatrix:    [][]bool{},
	}
	return graph
}

func (g *Graph[T]) AddNode(node T) {
	if g.repType == AdjacencyList {
		g.AddNodeAdjList(node)
	} else {
		g.AddNodeAdjMatrix(node)
	}
}

func (g *Graph[T]) RemoveNode(node T) {
	if g.repType == AdjacencyList {
		g.RemoveNodeAdjList(node)
	} else {
		g.RemoveNodeAdjMatrix(node)
	}
}

func (g *Graph[T]) AddEdge(from T, to T) {
	if g.repType == AdjacencyList {
		g.AddEdgeAdjList(from, to)
	} else {
		g.AddEdgeAdjMatrix(from, to)
	}
}

func (g *Graph[T]) RemoveEdge(from T, to T) {
	if g.repType == AdjacencyList {
		g.RemoveEdgeAdjList(from, to)
	} else {
		g.RemoveEdgeAdjMatrix(from, to)
	}
}

func (g *Graph[T]) HasNode(node T) bool {
	_, exists := g.nodes[node]
	return exists
}

func (g *Graph[T]) HasEdge(from T, to T) bool {
	if g.repType == AdjacencyList {
		return g.HasEdgeAdjList(from, to)
	} else {
		return g.HasEdgeAdjMatrix(from, to)
	}
}

func (g *Graph[T]) Neighbours(node T) []T {
	if g.repType == AdjacencyList {
		return g.NeighboursAdjList(node)
	} else {
		return g.NeighboursAdjMatrix(node)
	}
}

func (g *Graph[T]) Nodes() []T {
	elems := make([]T, 0, len(g.nodes))
	for key, _ := range g.nodes {
		elems = append(elems, key)
	}
	return elems
}

func (g *Graph[T]) Edges() [][2]T {
	if g.repType == AdjacencyList {
		return g.EdgesAdjList()
	} else {
		return g.EdgesAdjMatrix()
	}
}

func (g *Graph[T]) OutDegree(node T) int {
	if g.repType == AdjacencyList {
		return g.OutDegreeAdjList(node)
	} else {
		return g.OutDegreeAdjMatrix(node)
	}
}

func (g *Graph[T]) InDegree(node T) int {
	if g.repType == AdjacencyList {
		return g.InDegreeAdjList(node)
	} else {
		return g.InDegreeAdjMatrix(node)
	}
}

func (g *Graph[T]) Degree(node T) int {
	return g.OutDegree(node)
}

func (g *Graph[T]) BFS(start T) []T {
	if g.repType == AdjacencyList {
		return g.BFSAdjList(start)
	} else {
		return g.BFSAdjMatrix(start)
	}
}

func (g *Graph[T]) DFSRecursive(start T) []T {

	if g.repType == AdjacencyList {
		return g.DFSRecursiveAdjList(start)
	} else {
		return g.DFSRecursiveAdjMatrix(start)
	}
}

func (g *Graph[T]) DFSIterative(start T) []T {
	if g.repType == AdjacencyList {
		return g.DFSIterativeAdjList(start)
	} else {
		return g.DFSIterativeAdjMatrix(start)
	}
}

func (g *Graph[T]) RecursiveDFSAllPathFinding(source T, target T) [][]T {
	if g.repType == AdjacencyList {
		return g.RecursiveDFSAllPathFindingAdjList(source, target)
	} else {
		return g.RecursiveDFSAllPathFindingAdjMatrix(source, target)
	}
}

func (g *Graph[T]) RecursiveDFSAnyPathFinding(source T, target T) []T {
	if g.repType == AdjacencyList {
		return g.RecursiveDFSAnyPathFindingAdjList(source, target)
	} else {
		return g.RecursiveDFSAnyPathFindingAdjMatrix(source, target)
	}
}

func (g *Graph[T]) DFSIterativeAllPathFinding(source T, target T) [][]T {
	if g.repType == AdjacencyList {
		return g.DFSIterativeAllPathFindingAdjList(source, target)
	} else {
		return g.DFSIterativeAllPathFindingAdjMatrix(source, target)
	}

}

func (g *Graph[T]) DFSIterativeAnyPathFinding(source T, target T) []T {
	if g.repType == AdjacencyList {
		return g.DFSIterativeAnyPathFindingAdjList(source, target)
	} else {
		return g.DFSIterativeAnyPathFindingAdjMatrix(source, target)
	}
}

func (g *Graph[T]) BFSShortestPath(source T, target T) []T {
	if g.repType == AdjacencyList {
		return g.BFSShortestPathAdjList(source, target)
	} else {
		return g.BFSShortestPathAdjMatrix(source, target)
	}
}

func (g *Graph[T]) HasCycleDirected() bool {
	if g.repType == AdjacencyList {
		return g.HasCycleDirectedAdjList()
	} else {
		return g.HasCycleDirectedAdjMatrix()
	}
}

func (g *Graph[T]) HasCycleUndirected() bool {
	if g.repType == AdjacencyList {
		return g.HasCycleUndirectedAdjList()
	} else {
		return g.HasCycleUndirectedAdjMatrix()
	}
}

func NewWeightedGraph[T comparable](graphType GraphType, repType RepresentationType) *WeightedGraph[T] {
	graph := &WeightedGraph[T]{
		graphType:    graphType,
		repType:      repType,
		nodes:        make(map[T]struct{}),
		adjList:      make(map[T]map[T]int),
		nodesToIndex: make(map[T]int),
		indexToNodes: []T{},
		adjMatrix:    [][]int{},
	}
	return graph
}

func (g *WeightedGraph[T]) AddNode(node T) {
	if g.repType == AdjacencyList {
		g.AddNodeAdjList(node)
	} else {
		g.AddNodeAdjMatrix(node)
	}
}

func (g *WeightedGraph[T]) RemoveNode(node T) {
	if g.repType == AdjacencyList {
		g.RemoveNodeAdjList(node)
	} else {
		g.RemoveNodeAdjMatrix(node)
	}
}

func (g *WeightedGraph[T]) AddEdge(from T, to T, weight int) {
	if g.repType == AdjacencyList {
		g.AddEdgeAdjList(from, to, weight)
	} else {
		g.AddEdgeAdjMatrix(from, to, weight)
	}
}

func (g *WeightedGraph[T]) RemoveEdge(from T, to T) {
	if g.repType == AdjacencyList {
		g.RemoveEdgeAdjList(from, to)
	} else {
		g.RemoveEdgeAdjMatrix(from, to)
	}
}

func (g *WeightedGraph[T]) HasNode(node T) bool {
	_, exists := g.nodes[node]
	return exists
}

func (g *WeightedGraph[T]) HasEdge(from T, to T) bool {
	if g.repType == AdjacencyList {
		return g.HasEdgeAdjList(from, to)
	} else {
		return g.HasEdgeAdjMatrix(from, to)
	}
}

func (g *WeightedGraph[T]) Neighbours(node T) []T {
	if g.repType == AdjacencyList {
		return g.NeighboursAdjList(node)
	} else {
		return g.NeighboursAdjMatrix(node)
	}
}

func (g *WeightedGraph[T]) Nodes() []T {
	elems := make([]T, 0, len(g.nodes))
	for key, _ := range g.nodes {
		elems = append(elems, key)
	}
	return elems
}

func (g *WeightedGraph[T]) Edges() []WeightedEdge[T] {
	if g.repType == AdjacencyList {
		return g.EdgesAdjList()
	} else {
		return g.EdgesAdjMatrix()
	}
}

func (g *WeightedGraph[T]) OutDegree(node T) int {
	if g.repType == AdjacencyList {
		return g.OutDegreeAdjList(node)
	} else {
		return g.OutDegreeAdjMatrix(node)
	}
}

func (g *WeightedGraph[T]) InDegree(node T) int {
	if g.repType == AdjacencyList {
		return g.InDegreeAdjList(node)
	} else {
		return g.InDegreeAdjMatrix(node)
	}
}

func (g *WeightedGraph[T]) Degree(node T) int {
	return g.OutDegree(node)
}

func (g *WeightedGraph[T]) BFS(start T) []T {
	if g.repType == AdjacencyList {
		return g.BFSAdjList(start)
	} else {
		return g.BFSAdjMatrix(start)
	}
}

func (g *WeightedGraph[T]) DFSRecursive(start T) []T {

	if g.repType == AdjacencyList {
		return g.DFSRecursiveAdjList(start)
	} else {
		return g.DFSRecursiveAdjMatrix(start)
	}
}

func (g *WeightedGraph[T]) DFSIterative(start T) []T {
	if g.repType == AdjacencyList {
		return g.DFSIterativeAdjList(start)
	} else {
		return g.DFSIterativeAdjMatrix(start)
	}
}

func (g *WeightedGraph[T]) RecursiveDFSAllPathFinding(source T, target T) [][]T {
	if g.repType == AdjacencyList {
		return g.RecursiveDFSAllPathFindingAdjList(source, target)
	} else {
		return g.RecursiveDFSAllPathFindingAdjMatrix(source, target)
	}
}

func (g *WeightedGraph[T]) RecursiveDFSAnyPathFinding(source T, target T) []T {
	if g.repType == AdjacencyList {
		return g.RecursiveDFSAnyPathFindingAdjList(source, target)
	} else {
		return g.RecursiveDFSAnyPathFindingAdjMatrix(source, target)
	}
}

func (g *WeightedGraph[T]) DFSIterativeAllPathFinding(source T, target T) [][]T {
	if g.repType == AdjacencyList {
		return g.DFSIterativeAllPathFindingAdjList(source, target)
	} else {
		return g.DFSIterativeAllPathFindingAdjMatrix(source, target)
	}

}

func (g *WeightedGraph[T]) DFSIterativeAnyPathFinding(source T, target T) []T {
	if g.repType == AdjacencyList {
		return g.DFSIterativeAnyPathFindingAdjList(source, target)
	} else {
		return g.DFSIterativeAnyPathFindingAdjMatrix(source, target)
	}
}

func (g *WeightedGraph[T]) BFSShortestPath(source T, target T) []T {
	if g.repType == AdjacencyList {
		return g.BFSShortestPathAdjList(source, target)
	} else {
		return g.BFSShortestPathAdjMatrix(source, target)
	}
}

func (g *WeightedGraph[T]) HasCycleDirected() bool {
	if g.repType == AdjacencyList {
		return g.HasCycleDirectedAdjList()
	} else {
		return g.HasCycleDirectedAdjMatrix()
	}
}

func (g *WeightedGraph[T]) HasCycleUndirected() bool {
	if g.repType == AdjacencyList {
		return g.HasCycleUndirectedAdjList()
	} else {
		return g.HasCycleUndirectedAdjMatrix()
	}
}
