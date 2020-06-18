package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sg := graph.NewSymbolGraph(filename, delimiter)
	g := sg.Graph()
	stdin := stdin.NewStdInLine()
	fmt.Println("Please input the source, use CTRL+D to cancel.")
	for !stdin.IsEmpty() {
		source := strings.Trim(stdin.ReadString(), " ")
		if sg.Contains(source) {
			s := sg.IndexOf(source)
			for _, v := range g.Adj(s) {
				fmt.Println(" ", sg.NameOf(v))
			}
			fmt.Println()
		} else {
			fmt.Printf("input not contain '%s'\n", source)
		}
	}
}
