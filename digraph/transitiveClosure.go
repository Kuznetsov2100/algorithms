package digraph

import "fmt"

type TransitiveClosure struct {
	tc []*DirectedDFS
}

func NewTransitiveClosure(G *Digraph) *TransitiveClosure {
	tc := make([]*DirectedDFS, G.V())
	for v := 0; v < G.V(); v++ {
		tc[v] = NewDirectedDFS(G, v)
	}
	return &TransitiveClosure{tc: tc}
}

func (t *TransitiveClosure) Reachable(v, w int) bool {
	t.validateVertex(v)
	t.validateVertex(w)
	return t.tc[v].IsMarked(w)
}

func (t *TransitiveClosure) validateVertex(v int) {
	V := len(t.tc)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
