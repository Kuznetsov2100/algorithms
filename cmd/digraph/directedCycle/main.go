package main

//   Finds a directed cycle in a digraph.

//    % go run main.go tinyDG.txt
//    Directed cycle: 3 5 4 3

//    % go run main.go tinyDAG.txt
//    No directed cycle

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := digraph.NewDigraphIn(in)
	finder := digraph.NewDirectedCycle(G)

	if finder.HasCycle() {
		fmt.Print("Directed cycle: ")
		for _, v := range finder.GetCycle() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("No directed cycle")
	}
	fmt.Println()
}
