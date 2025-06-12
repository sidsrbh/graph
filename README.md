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

MIT License (or your choice)

````

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
