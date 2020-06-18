package main

//   %  go run main.go routes.txt " "
//   JFK
//      MCO
//      ATL
//      ORD
//   LAX
//      PHX
//      LAS

//   % go run main.go movies.txt "/"
//   Tin Men (1987)
//      Hershey, Barbara
//      Geppi, Cindy
//      Jones, Kathy (II)
//      Herr, Marcia
//      ...
//      Blumenfeld, Alan
//      DeBoy, David
//   Bacon, Kevin
//      Woodsman, The (2004)
//      Wild Things (1998)
//      Where the Truth Lies (2005)
//      Tremors (1990)
//      ...
//      Apollo 13 (1995)
//      Animal House (1978)

//   Assumes that input file is encoded using UTF-8.
//   % iconv -f ISO-8859-1 -t UTF-8 movies-iso8859.txt > movies.txt

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
	fmt.Println("Please input the source, you will get the adjcent vertexs of the source, use CTRL+D to cancel.")
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
