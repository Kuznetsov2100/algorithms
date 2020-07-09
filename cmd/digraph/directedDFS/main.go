package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/handane123/algorithms/digraph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	in := stdin.NewInFileWords(os.Args[1])
	G, _ := digraph.NewDigraphIn(in)
	var sources []int
	for i := 2; i < len(os.Args); i++ {
		s, _ := strconv.Atoi(os.Args[i])
		sources = append(sources, s)
	}

	dfs := digraph.NewDirectedDFSources(G, sources)

	for v := 0; v < G.V(); v++ {
		if dfs.IsMarked(v) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

}
