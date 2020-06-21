package main

import (
	"fmt"
	"os"

	"github.com/handane123/algorithms/digraph"
)

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sg := digraph.NewSymbolDigraph(filename, delimiter)
	topological := digraph.NewTopological(sg.Digraph())

	for _, v := range topological.Order() {
		fmt.Println(sg.NameOf(v))
	}
}
