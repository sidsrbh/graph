// Problem 6: Strongly Connected Components (Kosaraju’s Algorithm)

// Problem Statement
// Given a directed graph, find all strongly connected components (SCCs) — maximal groups of nodes where every node can reach every other node in the same group via a directed path.

// Example Input

// Nodes: A, B, C, D, E, F
// Edges: (A → B), (B → C), (C → A), (B → D), (D → E), (E → F), (F → D)

// Output:

// There are 2 SCCs:

// {A, B, C}
// {D, E, F}
// Explanation
// SCC 1: A, B, C (there are cycles among these nodes; you can reach any from any other following directed edges)
// SCC 2: D, E, F (these three form a cycle as well)

// / ===========>>>>>>>> UNFINISHED PROBLEM <<<<<<======================= ////
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

func (g *AdjListGraph[T]) StronglyConnectedComponents() [][]T {
	visited := map[T]struct{}{}
	recStack := map[T]struct{}{}
	orders := [][]T{}
	var dfs func(node T)
	dfs = func(node T) {
		if _, ok := visited[node]; !ok {
			visited[node] = struct{}{}
			recStack[node] = struct{}{}

			for i := range g.AdjList[node] {
				if _, ok := visited[i]; !ok {
					dfs(i)
				} else if _, ok := recStack[i]; ok {
					order := []T{}
					for j, _ := range recStack {
						order = append(order, j)
					}
					orders = append(orders, order)
				}
			}
		}
		delete(recStack, node)
	}

	for i := range g.Nodes {
		if _, ok := visited[i]; !ok {
			dfs(i)
		}
	}
	return orders
}

// If you start DFS at `A`:
// - You’ll traverse: A → B → C → D → A, which is a cycle (A, B, C, D).
// You might think: "Mark these as one SCC."
// - But when you continue DFS, you can reach G → E → F → B.

// But **F → B** connects back to the main cycle!
// So, the whole set A, B, C, D, E, F, G are in one SCC.
// If you only look for cycles and their recStack,
// you might incorrectly break them into smaller pieces or
//  miss the connections.

// #### What about backtracking to already found SCCs?

// - If you finish DFS for a node and see it’s already “in an SCC”,
// simply merging current recStack to that SCC works
// **if and only if the whole recStack is also reachable
// in both directions from the SCC**. This is hard to
// guarantee just by looking at recStack.

// ---

// ### 3. **Kosaraju’s/Tarjan’s: Why the Standard Algorithms Are Needed**

// These classic algorithms:
// - Properly handle all cases, including
// “long chains” connecting into cycles.
// - **Tarjan’s:** Assigns “low-link” values to correctly group nodes together.
// - **Kosaraju’s:** Uses finishing order in one DFS,
// then a reverse-graph DFS to find SCCs.

// ---

// ### **Summary**

// - **Every cycle is an SCC,** but not every SCC is just a
// single simple cycle.
// - Simply using cycles and recStack doesn’t always group all
// the right nodes together,
// especially in graphs with "tails" leading into cycles.
// - **Standard algorithms** guarantee you get the *largest* strongly
// connected set each time.

// ---
