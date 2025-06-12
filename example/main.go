package main

import (
	"fmt"

	"github.com/sidsrbh/graph"
	// Replace this import path with the path where your graph package lives
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
