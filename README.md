# Go Generic Graph Library

A flexible, generic, and extensible Graph library in Go (Golang) supporting both **Adjacency List** and **Adjacency Matrix** representations. Works for **Directed** and **Undirected** graphs, and supports all comparable types as nodes.

---

## Features

- **Generic Graphs:** Any comparable type as node (`int`, `string`, custom structs, etc.)
- **Dual Representation:** Switch between Adjacency List and Adjacency Matrix at creation.
- **Directed & Undirected:** Choose your graph type.
- **Core Operations:** Add/Remove nodes and edges, check connectivity, degrees, neighbor listings.
- **Traversals:** BFS, DFS (iterative & recursive), and all/any path finding.
- **Cycle Detection:** For both directed and undirected graphs, with robust algorithms for each representation.
- **Shortest Path (BFS):** Get shortest path between nodes.

---

## Getting Started

### 1. **Installation**

Simply copy the `graph` package folder into your project.

### 2. **Usage Example**

See [`main.go`](main.go) for a complete demo.

---

## Example

```go
package main

import (
	"fmt"
	"githib.com/sidsrbh/graph" // update import as per your project
)

func main() {
	// Create an undirected graph with adjacency list
	g := graph.NewGraph[string](graph.Undirected, graph.AdjacencyList)

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")

	fmt.Println("Nodes:", g.Nodes())
	fmt.Println("Edges:", g.Edges())
	fmt.Println("BFS from A:", g.BFS("A"))
	fmt.Println("DFS (iterative) from A:", g.DFSIterative("A"))
	fmt.Println("Cycle present?", g.HasCycleUndirected())

	// Switch to a directed graph with adjacency matrix
	dg := graph.NewGraph[int](graph.Directed, graph.AdjacencyMatrix)
	dg.AddEdge(1, 2)
	dg.AddEdge(2, 3)
	dg.AddEdge(3, 1)

	fmt.Println("\nDirected Matrix Graph")
	fmt.Println("Nodes:", dg.Nodes())
	fmt.Println("Edges:", dg.Edges())
	fmt.Println("BFS from 1:", dg.BFS(1))
	fmt.Println("Cycle present?", dg.HasCycleDirected())
}
````

---

## API Overview

* **Graph Construction:**

  * `NewGraph[T comparable](graphType GraphType, repType RepresentationType) *Graph[T]`

* **Core Methods:**

  * `AddNode(node T)`
  * `RemoveNode(node T)`
  * `AddEdge(from, to T)`
  * `RemoveEdge(from, to T)`
  * `HasNode(node T)`
  * `HasEdge(from, to T)`
  * `Neighbours(node T) []T`
  * `Nodes() []T`
  * `Edges() [][2]T`
  * `OutDegree(node T) int`
  * `InDegree(node T) int`
  * `Degree(node T) int`

* **Traversals:**

  * `BFS(start T) []T`
  * `DFSRecursive(start T) []T`
  * `DFSIterative(start T) []T`
  * `RecursiveDFSAllPathFinding(source, target T) [][]T`
  * `RecursiveDFSAnyPathFinding(source, target T) []T`
  * `DFSIterativeAllPathFinding(source, target T) [][]T`
  * `DFSIterativeAnyPathFinding(source, target T) []T`
  * `BFSShortestPath(source, target T) []T`

* **Cycle Detection:**

  * `HasCycleDirected() bool`
  * `HasCycleUndirected() bool`

---

## **How to Use with Your Project**

1. Place your `graph` package directory inside your project folder.
2. Import and use as shown above.
3. Build your algorithms and apps using the power of generic, flexible graph representations!

---

## **License**

MIT License

````
```
---

## **Example main.go**

```go
package main

import (
	"fmt"
	// Replace this import path with the path where your graph package lives
	"githib.com/sidsrbh/graph"
)

func main() {
	// --------- Adjacency List Undirected Example -----------
	fmt.Println("Undirected Graph (Adjacency List):")
	g := graph.NewGraph[string](graph.Undirected, graph.AdjacencyList)
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")

	fmt.Println("Nodes:", g.Nodes())
	fmt.Println("Edges:", g.Edges())
	fmt.Println("BFS from A:", g.BFS("A"))
	fmt.Println("DFSIterative from A:", g.DFSIterative("A"))
	fmt.Println("Cycle present?:", g.HasCycleUndirected())

	// --------- Adjacency Matrix Directed Example -----------
	fmt.Println("\nDirected Graph (Adjacency Matrix):")
	dg := graph.NewGraph[int](graph.Directed, graph.AdjacencyMatrix)
	dg.AddEdge(1, 2)
	dg.AddEdge(2, 3)
	dg.AddEdge(3, 4)
	dg.AddEdge(4, 2) // Creates a cycle

	fmt.Println("Nodes:", dg.Nodes())
	fmt.Println("Edges:", dg.Edges())
	fmt.Println("BFS from 1:", dg.BFS(1))
	fmt.Println("DFSRecursive from 1:", dg.DFSRecursive(1))
	fmt.Println("Cycle present?:", dg.HasCycleDirected())

	// --------- Shortest Path Example ----------
	fmt.Println("\nShortest path from 1 to 4:", dg.BFSShortestPath(1, 4))
}
````

---

## **How to Run**

1. Place your `graph` package directory and `main.go` in your project.
2. Update the import path as per your project structure.
3. Run:

   ```
   go run main.go
   ```
4. You should see the graph nodes, edges, traversals, and cycle detection outputs!

---

```
## Weighted Graphs

Your package also supports **weighted graphs**—for both adjacency list and adjacency matrix representations, in directed or undirected forms, with all core and traversal functions as for unweighted graphs.

### **When to Use**

Use `WeightedGraph` when your edges have associated weights (distances, costs, capacities, etc.), as in road networks, routing, flow, or shortest path problems.

---

### **Creating a Weighted Graph**

```go
// Undirected Weighted Graph (Adjacency List)
wg := graph.NewWeightedGraph[string](graph.Undirected, graph.AdjacencyList)

// Add weighted edges
wg.AddEdge("A", "B", 4)
wg.AddEdge("A", "C", 2)
wg.AddEdge("B", "C", 5)
wg.AddEdge("B", "D", 10)
wg.AddEdge("C", "D", 3)

fmt.Println("Nodes:", wg.Nodes())
fmt.Println("Edges:", wg.Edges()) // WeightedEdge{Edge: [2]string, Weight: int}

// Traversal works the same way
fmt.Println("BFS from A:", wg.BFS("A"))
fmt.Println("DFSIterative from A:", wg.DFSIterative("A"))
```

#### **Directed Weighted Graph (Adjacency Matrix) Example**

```go
// Directed Weighted Graph with integers
wdg := graph.NewWeightedGraph[int](graph.Directed, graph.AdjacencyMatrix)
wdg.AddEdge(1, 2, 7)
wdg.AddEdge(2, 3, 1)
wdg.AddEdge(3, 1, 2)

fmt.Println("Nodes:", wdg.Nodes())
fmt.Println("Edges:", wdg.Edges()) // WeightedEdge{Edge: [2]int, Weight: int}
fmt.Println("Cycle present?:", wdg.HasCycleDirected())
```

---

### **WeightedGraph API**

All APIs parallel the unweighted version, but with weights:

* `AddEdge(from, to T, weight int)` — add a weighted edge
* `RemoveEdge(from, to T)`
* `HasEdge(from, to T) bool`
* `Neighbours(node T) []T`
* `Edges() []WeightedEdge[T]` — returns slice of `{Edge: [2]T, Weight: int}`
* `Nodes() []T`
* `OutDegree(node T) int`
* `InDegree(node T) int`
* `BFS(start T) []T`, `DFSRecursive(start T) []T`, `DFSIterative(start T) []T`
* Pathfinding and cycle detection as before

#### **WeightedEdge Type**

```go
type WeightedEdge[T comparable] struct {
    Edge   [2]T  // from, to
    Weight int   // weight of the edge
}
```

---

### **Example: WeightedGraph with Shortest Path (Coming Soon)**

*Build Dijkstra’s, Bellman-Ford, Floyd-Warshall etc. on top of this interface!*

```go
// Use Edges() for custom algorithms (e.g. Dijkstra, Floyd-Warshall)
for _, edge := range wg.Edges() {
    fmt.Printf("Edge %v -> %v, weight=%d\n", edge.Edge[0], edge.Edge[1], edge.Weight)
}
```

---

### **Switching Between List and Matrix**

Just change the constructor’s `repType` argument. Everything else remains identical:

```go
// Adjacency Matrix version
wg := graph.NewWeightedGraph[string](graph.Undirected, graph.AdjacencyMatrix)
```

---

### **Quick Reference Table**

| Feature                            | Unweighted Graph | WeightedGraph  |
| ---------------------------------- | ---------------- | -------------- |
| Add/Remove nodes                   | ✅                | ✅              |
| Add/Remove edges                   | ✅                | ✅              |
| Edge weights                       | ❌                | ✅              |
| Directed/Undirected                | ✅                | ✅              |
| Adjacency List/Matrix              | ✅                | ✅              |
| Traversals (BFS/DFS)               | ✅                | ✅              |
| Degree, neighbors, edges           | ✅                | ✅              |
| Cycle detection                    | ✅                | ✅              |
| Custom algorithms (Dijkstra, etc.) | (not built-in)   | (not built-in) |

---

### **Summary**

* Use `Graph[T]` for simple (unweighted) graphs.
* Use `WeightedGraph[T]` for graphs with edge weights.
* Both APIs are parallel and easy to swap.
* Extend with custom algorithms as needed.

---

```

