package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/graph"
)

func main() {
	generator := graph.NewGraphGenerator()
	V, _ := strconv.Atoi(os.Args[1])
	E, _ := strconv.Atoi(os.Args[2])
	V1 := V / 2
	V2 := V - V1
	fmt.Println("complete graph")
	fmt.Println(generator.Complete(V))
	fmt.Println()

	fmt.Println("simple")
	fmt.Println(generator.Simple(V, E))
	fmt.Println()

	fmt.Println("Erdos-Renyi")
	p := float64(E) / float64((V * (V - 1) / 2.0))
	fmt.Println(generator.SimpleP(V, p))
	fmt.Println()

	fmt.Println("complete bipartite")
	fmt.Println(generator.CompleteBipartite(V1, V2))
	fmt.Println()

	fmt.Println("bipartite")
	fmt.Println(generator.BipartiteGraph(V1, V2, E))
	fmt.Println()

	fmt.Println("Erdos Renyi bipartite")
	q := float64(E) / float64((V1 * V2))
	fmt.Println(generator.BipartiteP(V1, V2, q))
	fmt.Println()

	fmt.Println("path")
	fmt.Println(generator.PathGraph(V))
	fmt.Println()

	fmt.Println("cycle")
	fmt.Println(generator.CycleGraph(V))
	fmt.Println()

	fmt.Println("binary tree")
	fmt.Println(generator.BinaryTree(V))
	fmt.Println()

	fmt.Println("tree")
	fmt.Println(generator.Tree(V))
	fmt.Println()

	fmt.Println("4-regular")
	fmt.Println(generator.Regular(V, 4))
	fmt.Println()

	fmt.Println("star")
	fmt.Println(generator.Star(V))
	fmt.Println()

	fmt.Println("wheel")
	fmt.Println(generator.Wheel(V))
	fmt.Println()

}
