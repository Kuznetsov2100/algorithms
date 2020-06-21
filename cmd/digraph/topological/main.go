package main

//   Compute topological ordering of a DAG or edge-weighted DAG.
//    Runs in O(E + V) time.

//    % go run main.go jobs.txt "/"
//    Calculus
//    Linear Algebra
//    Introduction to CS
//    Advanced Programming
//    Algorithms
//    Theoretical CS
//    Artificial Intelligence
//    Robotics
//    Machine Learning
//    Neural Networks
//    Databases
//    Scientific Computing
//    Computational Biology

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
