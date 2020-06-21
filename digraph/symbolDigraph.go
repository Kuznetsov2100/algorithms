package digraph

import (
	"fmt"
	"strings"

	"github.com/handane123/algorithms/searching"
	"github.com/handane123/algorithms/stdin"
)

// SymbolDigraph struct represents an digraph,
// where the vertex names are arbitrary strings.
// By providing mappings between string vertex names and integers,
// it serves as a wrapper around the Graph data type,
// which assumes the vertex names are integers between 0 and V - 1.
// It also supports initializing a symbol digraph from a file.
// This implementation uses an ST to map from strings to integers,
// an array to map from integers to strings, and a Graph to store the underlying digraph.
// The indexOf and contains operations take time proportional to log V,
// where V is the number of vertices. The nameOf operation takes constant time.
type SymbolDigraph struct {
	st    *searching.ST // string -> index
	keys  []key         // index  -> string
	graph *Digraph      // the underlying digraph
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

// NewSymbolDigraph initializes a digraph from a file using the specified delimiter.
func NewSymbolDigraph(filename, delimiter string) *SymbolDigraph {
	sg := &SymbolDigraph{st: searching.NewST(func(a, b interface{}) int {
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
	sg.graph = NewDigraph(sg.st.Size())
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

// Contains returns ture if the digraph contain the vertex named s
func (sg *SymbolDigraph) Contains(s string) bool {
	ok, _ := sg.st.Contains(key(s))
	return ok
}

// IndexOf returns the integer associated with the vertex named s.
func (sg *SymbolDigraph) IndexOf(s string) int {
	if val, _ := sg.st.Get(key(s)); val != nil {
		return val.(int)
	} else {
		return -1
	}
}

// NameOf returns the name of the vertex associated with the integer v.
func (sg *SymbolDigraph) NameOf(v int) string {
	sg.validateVertex(v)
	return string(sg.keys[v])
}

// Digraph returns the digraph associated with the symbol digraph.
func (sg *SymbolDigraph) Digraph() *Digraph {
	return sg.graph
}

func (sg *SymbolDigraph) validateVertex(v int) {
	V := sg.graph.V()
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
