package digraph

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

// DepthFirstOrder struct represents a data type for determining depth-first search
// ordering of the vertices in a digraph, including preorder, postorder, and reverse postorder.
// This implementation uses depth-first search. Each constructor takes O(V + E) time,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the digraph).
type DepthFirstOrder struct {
	marked      []bool
	pre         []int
	post        []int
	preorder    *arrayqueue.Queue
	postorder   *arrayqueue.Queue
	preCounter  int
	postCounter int
}

// NewDepthFirstOrder determines a depth-first order for the digraph G.
func NewDepthFirstOrder(G *Digraph) *DepthFirstOrder {
	dfo := &DepthFirstOrder{
		marked:    make([]bool, G.V()),
		pre:       make([]int, G.V()),
		post:      make([]int, G.V()),
		preorder:  arrayqueue.New(),
		postorder: arrayqueue.New(),
	}
	for v := 0; v < G.V(); v++ {
		if !dfo.marked[v] {
			dfo.dfs(G, v)
		}
	}
	return dfo
}

func (dfo *DepthFirstOrder) dfs(G *Digraph, v int) {
	dfo.marked[v] = true
	dfo.pre[v] = dfo.preCounter
	dfo.preCounter++
	dfo.preorder.Enqueue(v)
	for _, w := range G.Adj(v) {
		if !dfo.marked[w] {
			dfo.dfs(G, w)
		}
	}
	dfo.post[v] = dfo.postCounter
	dfo.postCounter++
	dfo.postorder.Enqueue(v)
}

// Pre returns the preorder number of vertex v.
func (dfo *DepthFirstOrder) Pre(v int) int {
	dfo.validateVertex(v)
	return dfo.pre[v]
}

// Post returns the postorder number of vertex v.
func (dfo *DepthFirstOrder) Post(v int) int {
	dfo.validateVertex(v)
	return dfo.post[v]
}

// PreOrder returns the vertices in preorder.
func (dfo *DepthFirstOrder) PreOrder() (pre []int) {
	for _, x := range dfo.preorder.Values() {
		pre = append(pre, x.(int))
	}
	return pre
}

// PostOrder returns the vertices in postorder.
func (dfo *DepthFirstOrder) PostOrder() (post []int) {
	for _, x := range dfo.postorder.Values() {
		post = append(post, x.(int))
	}
	return post
}

// ReversePost returns the vertices in reverse postorder.
func (dfo *DepthFirstOrder) ReversePost() (reverse []int) {
	n := dfo.postorder.Size()
	reverse = make([]int, n)
	for i, x := range dfo.postorder.Values() {
		reverse[n-i-1] = x.(int)
	}
	return reverse
}

func (dfo *DepthFirstOrder) validateVertex(v int) {
	V := len(dfo.marked)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
