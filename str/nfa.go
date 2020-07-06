package str

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
	"github.com/handane123/algorithms/digraph"
)

type NFA struct {
	graph  *digraph.Digraph
	regexp string
	m      int
}

func NewNFA(regexp string) *NFA {
	nfa := &NFA{regexp: regexp, m: len(regexp), graph: digraph.NewDigraph(len(regexp) + 1)}
	ops := arraystack.New()
	for i := 0; i < nfa.m; i++ {
		lp := i
		if string(regexp[i]) == "(" || string(regexp[i]) == "|" {
			ops.Push(i)
		} else if string(regexp[i]) == ")" {
			val, _ := ops.Pop()
			or := val.(int)

			if string(regexp[or]) == "|" {
				val, _ := ops.Pop()
				lp = val.(int)
				nfa.graph.AddEdge(lp, or+1)
				nfa.graph.AddEdge(or, i)
			} else if string(regexp[or]) == "(" {
				lp = or
			}
		}

		if i < nfa.m-1 && string(regexp[i+1]) == "*" {
			nfa.graph.AddEdge(lp, i+1)
			nfa.graph.AddEdge(i+1, lp)
		}
		if string(regexp[i]) == "(" || string(regexp[i]) == "*" || string(regexp[i]) == ")" {
			nfa.graph.AddEdge(i, i+1)
		}
	}
	if ops.Size() != 0 {
		panic("invalid regular expression")
	}
	return nfa
}

func (nfa *NFA) Recognizes(txt string) bool {
	dfs := digraph.NewDirectedDFS(nfa.graph, 0)
	var pc []int
	for v := 0; v < nfa.graph.V(); v++ {
		if dfs.IsMarked(v) {
			pc = append(pc, v)
		}
	}

	for i := 0; i < len(txt); i++ {
		if string(txt[i]) == "*" || string(txt[i]) == "|" || string(txt[i]) == "(" || string(txt[i]) == ")" {
			panic(fmt.Sprintf("text contains the metacharacter '%c'\n", txt[i]))
		}
		match := make([]int, 0)
		for _, v := range pc {
			if v == nfa.m {
				continue
			}
			if nfa.regexp[v] == txt[i] || string(nfa.regexp[v]) == "." {
				match = append(match, v+1)
			}
		}
		dfs = digraph.NewDirectedDFSources(nfa.graph, match)
		pc = make([]int, 0)
		for v := 0; v < nfa.graph.V(); v++ {
			if dfs.IsMarked(v) {
				pc = append(pc, v)
			}
		}
		if len(pc) == 0 {
			return false
		}
	}

	for _, v := range pc {
		if v == nfa.m {
			return true
		}
	}
	return false
}
