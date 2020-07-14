package main

import (
	"fmt"

	"github.com/handane123/algorithms/digraph"
)


func main() {
	e := digraph.NewFlowEdge(12, 23, 4.56)
	fmt.Println(e)
}