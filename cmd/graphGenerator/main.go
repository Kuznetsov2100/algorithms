package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/graph"
)

func main() {
	V, _ := strconv.Atoi(os.Args[1])
	E, _ := strconv.Atoi(os.Args[2])
	V1 := V / 2
	V2 := V - V1
	fmt.Println("complete graph")
	fmt.Println(graph.Complete(V))
	fmt.Println()

	fmt.Println("simple")
	fmt.Println(graph.Simple(V, E))
	fmt.Println()

	fmt.Println("Erdos-Renyi")
	p := float64(E) / float64((V * (V - 1) / 2.0))
	fmt.Println(graph.SimpleP(V, p))
	fmt.Println()

	fmt.Println("complete bipartite")
	fmt.Println(graph.CompleteBipartite(V1, V2))
	fmt.Println()

	fmt.Println("bipartite")
	fmt.Println(graph.BipartiteGraph(V1, V2, E))
	fmt.Println()

	fmt.Println("Erdos Renyi bipartite")
	q := float64(E) / float64((V1 * V2))
	fmt.Println(graph.BipartiteP(V1, V2, q))
	fmt.Println()

	fmt.Println("path")
	fmt.Println(graph.PathGraph(V))
	fmt.Println()

	fmt.Println("cycle")
	fmt.Println(graph.CycleGraph(V))
	fmt.Println()

	fmt.Println("binary tree")
	fmt.Println(graph.BinaryTree(V))
	fmt.Println()

	fmt.Println("tree")
	fmt.Println(graph.Tree(V))
	fmt.Println()

	fmt.Println("4-regular")
	fmt.Println(graph.Regular(V, 4))
	fmt.Println()

	fmt.Println("star")
	fmt.Println(graph.Star(V))
	fmt.Println()

	fmt.Println("wheel")
	fmt.Println(graph.Wheel(V))
	fmt.Println()

}
