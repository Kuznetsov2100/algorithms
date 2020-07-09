package main

//   Identifies a cycle.
//   Runs in O(E + V) time.

//   % go run main.go tinyG.txt
//   3 4 5 3

//   % go run main.go mediumG.txt
//   15 0 225 15

//   % go run main.go largeG.txt
//   996673 762 840164 4619 785187 194717 996673

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := graph.NewGraphIn(in)
	finder := graph.NewCycle(G)

	if finder.HasCycle() {
		for _, v := range finder.GetCycle() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Graph is acyclic")
	}
}
