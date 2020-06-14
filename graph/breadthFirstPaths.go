package graph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

const infinity = math.MaxInt64

type BreadFirstPaths struct {
	marked []bool
	edgeTo []int
	distTo []int
}

func NewBreadFirstPaths(G *Graph, s int) *BreadFirstPaths {
	bfp := &BreadFirstPaths{
		marked: make([]bool, G.V()),
		edgeTo: make([]int, G.V()),
		distTo: make([]int, G.V()),
	}
	bfp.validateVertex(s)
	bfp.bfs(G, s)
	return bfp
}

func (bfp *BreadFirstPaths) bfs(G *Graph, s int) {
	q := arrayqueue.New()
	for v := 0; v < G.V(); v++ {
		bfp.distTo[v] = infinity
	}
	bfp.distTo[s] = 0
	bfp.marked[s] = true
	q.Enqueue(s)

	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		v := val.(int)
		for _, w := range G.Adj(v) {
			if !bfp.marked[w] {
				bfp.edgeTo[w] = v
				bfp.distTo[w] = bfp.distTo[v] + 1
				bfp.marked[w] = true
				q.Enqueue(w)
			}
		}
	}
}

func (bfp *BreadFirstPaths) HasPathTo(v int) bool {
	bfp.validateVertex(v)
	return bfp.marked[v]
}

func (bfp *BreadFirstPaths) DistTo(v int) int {
	bfp.validateVertex(v)
	return bfp.distTo[v]
}

func (bfp *BreadFirstPaths) PathTo(v int) (p []int) {
	bfp.validateVertex(v)
	if !bfp.HasPathTo(v) {
		return nil
	}
	path := arraystack.New()
	x := v
	for ; bfp.distTo[x] != 0; x = bfp.edgeTo[x] {
		path.Push(x)
	}
	path.Push(x)
	for _, val := range path.Values() {
		p = append(p, val.(int))
	}
	return p
}

func (bfp *BreadFirstPaths) validateVertex(v int) {
	length := len(bfp.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}
