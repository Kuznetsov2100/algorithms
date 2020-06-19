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

	fmt.Println("complete digraph")
	fmt.Println(digraph.Complete(V))
	fmt.Println()

	fmt.Println("simple")
	fmt.Println(digraph.Simple(V, E))
	fmt.Println()

	fmt.Println("path")
	fmt.Println(digraph.PathDigraph(V))
	fmt.Println()

	fmt.Println("cycle")
	fmt.Println(digraph.CycleDigraph(V))
	fmt.Println()

	fmt.Println("Eulerian path")
	fmt.Println(digraph.EulerianPathDigraph(V, E))
	fmt.Println()

	fmt.Println("Eulerian cycle")
	fmt.Println(digraph.EulerianCycleDigraph(V, E))
	fmt.Println()

	fmt.Println("binary tree")
	fmt.Println(digraph.BinaryTree(V))
	fmt.Println()

	fmt.Println("tournament")
	fmt.Println(digraph.Tournament(V))
	fmt.Println()

	fmt.Println("DAG")
	fmt.Println(digraph.Dag(V, E))
	fmt.Println()

	fmt.Println("rooted-in DAG")
	fmt.Println(digraph.RootedInDAG(V, E))
	fmt.Println()

	fmt.Println("rooted-out DAG")
	fmt.Println(digraph.RootedOutDAG(V, E))
	fmt.Println()

	fmt.Println("rooted-in tree")
	fmt.Println(digraph.RootedInTree(V))
	fmt.Println()

	fmt.Println("rooted-out tree")
	fmt.Println(digraph.RootedOutTree(V))
	fmt.Println()
}
