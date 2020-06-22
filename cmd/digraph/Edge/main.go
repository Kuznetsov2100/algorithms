package main

import (
	"fmt"

	"github.com/handane123/algorithms/digraph"
)

func main() {
	e := digraph.NewEdge(12, 34, 5.67)
	fmt.Println(e)
}
