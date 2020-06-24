package main

import (
	"fmt"

	"github.com/handane123/algorithms/digraph"
)

func main() {
	edge := digraph.NewDirectedEdge(1, 2, 3.29)
	fmt.Println(edge)
}
