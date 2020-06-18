package graph

import (
	"fmt"
	"strings"

	"github.com/handane123/algorithms/searching"
	"github.com/handane123/algorithms/stdin"
)

type SymbolGraph struct {
	st    *searching.ST
	keys  []key
	graph *Graph
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

func NewSymbolGraph(filename, delimiter string) *SymbolGraph {
	sg := &SymbolGraph{st: searching.NewST(func(a, b interface{}) int {
		a1, b1 := a.(key), b.(key)
		return a1.CompareTo(b1)
	})}
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
	sg.keys = make([]key, sg.st.Size())
	for _, name := range sg.st.Keys() {
		val, _ := sg.st.Get(name)
		sg.keys[val.(int)] = name.(key)
	}

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

func (sg *SymbolGraph) Contains(s string) bool {
	ok, _ := sg.st.Contains(key(s))
	return ok
}

func (sg *SymbolGraph) IndexOf(s string) int {
	if val, _ := sg.st.Get(key(s)); val != nil {
		return val.(int)
	} else {
		return -1
	}
}

func (sg *SymbolGraph) NameOf(v int) string {
	sg.validateVertex(v)
	return string(sg.keys[v])
}

func (sg *SymbolGraph) Graph() *Graph {
	return sg.graph
}

func (sg *SymbolGraph) validateVertex(v int) {
	V := sg.graph.V()
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
