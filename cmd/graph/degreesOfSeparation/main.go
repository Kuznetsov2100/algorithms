package main

//   %  go run main.go routes.txt " " "JFK"
//   LAS
//      JFK
//      ORD
//      DEN
//      LAS
//   DFW
//      JFK
//      ORD
//      DFW
//   EWR
//      Not in database.

//   % go run main.go movies.txt "/" "Bacon, Kevin"
//   Kidman, Nicole
//      Bacon, Kevin
//      Woodsman, The (2004)
//      Grier, David Alan
//      Bewitched (2005)
//      Kidman, Nicole
//   Grant, Cary
//      Bacon, Kevin
//      Planes, Trains & Automobiles (1987)
//      Martin, Steve (I)
//      Dead Men Don't Wear Plaid (1982)
//      Grant, Cary

//   % go run main.go movies.txt "/" "Animal House (1978)"
//   Titanic (1997)
//      Animal House (1978)
//      Allen, Karen (I)
//      Raiders of the Lost Ark (1981)
//      Taylor, Rocky (I)
//      Titanic (1997)
//   To Catch a Thief (1955)
//      Animal House (1978)
//      Vernon, John (I)
//      Topaz (1969)
//      Hitchcock, Alfred (I)
//      To Catch a Thief (1955)

import (
	"fmt"
	"os"
	"strings"

	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/io/stdin"
)

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	source := os.Args[3]

	sg := graph.NewSymbolGraph(filename, delimiter)
	G := sg.Graph()
	if !sg.Contains(source) {
		fmt.Println(source, " not in database.")
		return
	}

	s := sg.IndexOf(source)
	bfs := graph.NewBreadthFirstPaths(G, s)

	stdin := stdin.NewStdInLine()
	fmt.Println("Please input a vertex, you will get the path from source to your vertex , use CTRL+D to cancel.")
	for !stdin.IsEmpty() {
		sink := strings.Trim(stdin.ReadString(), " ")
		if sg.Contains(sink) {
			t := sg.IndexOf(sink)
			if bfs.HasPathTo(t) {
				for _, v := range bfs.PathTo(t) {
					fmt.Println("   ", sg.NameOf(v))
				}
			} else {
				fmt.Println("Not connected.")
			}
		} else {
			fmt.Printf("\"%s\" is not in database.\n", sink)
		}
	}
}
