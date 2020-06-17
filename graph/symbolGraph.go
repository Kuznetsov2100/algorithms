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

func keyComparator(a, b interface{}) int {
	a1 := a.(key)
	b1 := b.(key)
	if a1 < b1 {
		return -1
	} else if a1 > b1 {
		return 1
	} else {
		return 0
	}
}
func NewSymbolGraph(filename string, delimiter string) *SymbolGraph {
	sg := &SymbolGraph{st: searching.NewST(keyComparator)}
	in := stdin.NewInFileLine(filename)

	for !in.IsEmpty() {
		a := strings.Split(in.ReadString(), delimiter)
		for index := range a {
			if ok, err := sg.st.Contains(key(a[index])); err != nil {
				fmt.Println(err)
			} else if !ok {
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
			w, _ := sg.st.Get(key(a[i]))
			sg.graph.AddEdge(v, w.(int))
		}
	}
	return sg
}

func (sg *SymbolGraph) Contains(s string) (bool, error) {
	return sg.st.Contains(key(s))
}

func (sg *SymbolGraph) IndexOf(s string) (int, error) {
	if val, err := sg.st.Get(key(s)); err != nil {
		return -1, err
	} else {
		return val.(int), nil
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
