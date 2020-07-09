package graph

import (
	"fmt"
	"strings"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/searching"
)

// SymbolGraph struct represents an undirected graph,
// where the vertex names are arbitrary strings.
// By providing mappings between string vertex names and integers,
// it serves as a wrapper around the Graph data type,
// which assumes the vertex names are integers between 0 and V - 1.
// It also supports initializing a symbol graph from a file.
// This implementation uses an ST to map from strings to integers,
// an array to map from integers to strings, and a Graph to store the underlying graph.
// The indexOf and contains operations take time proportional to log V,
// where V is the number of vertices. The nameOf operation takes constant time.
type SymbolGraph struct {
	st    *searching.ST // string -> index
	keys  []key         // index  -> string
	graph *Graph        // the underlying graph
}

type key string

func (s key) CompareTo(k searching.Key) int {
	t := k.(key)
	if s < t {
		return -1
	} else if s > t {
		return 1
	} else {
		return 0
	}
}

// NewSymbolGraph initializes a graph from a file using the specified delimiter.
func NewSymbolGraph(filename, delimiter string) *SymbolGraph {
	sg := &SymbolGraph{st: searching.NewST(func(a, b interface{}) int {
		a1, b1 := a.(key), b.(key)
		return a1.CompareTo(b1)
	})}

	// First pass builds the index by reading strings to associate
	// distinct strings with an index
	in := stdin.NewInFileLine(filename)
	for !in.IsEmpty() {
		a := strings.Split(in.ReadString(), delimiter)
		for index := range a {
			if ok, _ := sg.st.Contains(key(a[index])); !ok {
				//nolint:errcheck
				sg.st.Put(key(a[index]), sg.st.Size())
			}
		}
	}
	// inverted index to get string keys in an array
	sg.keys = make([]key, sg.st.Size())
	for _, name := range sg.st.Keys() {
		val, _ := sg.st.Get(name)
		sg.keys[val.(int)] = name.(key)
	}

	// second pass builds the graph by connecting first vertex on each
	// line to all others
	sg.graph = NewGraph(sg.st.Size())
	in = stdin.NewInFileLine(filename)
	for !in.IsEmpty() {
		a := strings.Split(in.ReadString(), delimiter)
		val, _ := sg.st.Get(key(a[0]))
		v := val.(int)
		for i := 1; i < len(a); i++ {
			wval, _ := sg.st.Get(key(a[i]))
			w := wval.(int)
			sg.graph.AddEdge(v, w)
		}
	}
	return sg
}

// Contains returns ture if graph contain the vertex named s
func (sg *SymbolGraph) Contains(s string) bool {
	ok, _ := sg.st.Contains(key(s))
	return ok
}

// IndexOf returns the integer associated with the vertex named s.
func (sg *SymbolGraph) IndexOf(s string) int {
	val, _ := sg.st.Get(key(s))
	if val != nil {
		return val.(int)
	}
	return -1
}

// NameOf returns the name of the vertex associated with the integer v.
func (sg *SymbolGraph) NameOf(v int) string {
	sg.validateVertex(v)
	return string(sg.keys[v])
}

// Graph returns the graph associated with the symbol graph.
func (sg *SymbolGraph) Graph() *Graph {
	return sg.graph
}

func (sg *SymbolGraph) validateVertex(v int) {
	V := sg.graph.V()
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
