package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/digraph"
)

func main() {
	V, _ := strconv.Atoi(os.Args[1])
	E, _ := strconv.Atoi(os.Args[2])

	generator := digraph.NewDigraphGenerator()
	fmt.Println("complete digraph")
	fmt.Println(generator.Complete(V))
	fmt.Println()

	fmt.Println("simple")
	fmt.Println(generator.Simple(V, E))
	fmt.Println()

	fmt.Println("path")
	fmt.Println(generator.PathDigraph(V))
	fmt.Println()

	fmt.Println("cycle")
	fmt.Println(generator.CycleDigraph(V))
	fmt.Println()

	fmt.Println("Eulerian path")
	fmt.Println(generator.EulerianPathDigraph(V, E))
	fmt.Println()

	fmt.Println("Eulerian cycle")
	fmt.Println(generator.EulerianCycleDigraph(V, E))
	fmt.Println()

	fmt.Println("binary tree")
	fmt.Println(generator.BinaryTree(V))
	fmt.Println()

	fmt.Println("tournament")
	fmt.Println(generator.Tournament(V))
	fmt.Println()

	fmt.Println("DAG")
	fmt.Println(generator.Dag(V, E))
	fmt.Println()

	fmt.Println("rooted-in DAG")
	fmt.Println(generator.RootedInDAG(V, E))
	fmt.Println()

	fmt.Println("rooted-out DAG")
	fmt.Println(generator.RootedOutDAG(V, E))
	fmt.Println()

	fmt.Println("rooted-in tree")
	fmt.Println(generator.RootedInTree(V))
	fmt.Println()

	fmt.Println("rooted-out tree")
	fmt.Println(generator.RootedOutTree(V))
	fmt.Println()
}
