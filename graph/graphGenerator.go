package graph

import (
	"errors"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/handane123/algorithms/dataStructure/priorityqueue"
)

type edge struct {
	v int
	w int
}

func newEdge(v, w int) *edge {
	if v < w {
		return &edge{v: v, w: w}
	}
	return &edge{v: w, w: v}

}

func comparator(a, b interface{}) int {
	a1 := a.(*edge)
	b1 := b.(*edge)
	if a1.v < b1.v {
		return -1
	}
	if a1.v > b1.v {
		return 1
	}
	if a1.w < b1.w {
		return -1
	}
	if a1.w > b1.w {
		return 1
	}
	return 0
}

// Simple returns a random simple graph containing V vertices and E edges.
// A simple graph is a graph with no self-loops and parallel edges.
func Simple(V, E int) (*Graph, error) {
	if E > V*(V-1)/2 {
		return nil, errors.New("too many edges")
	}
	if E < 0 {
		return nil, errors.New("too few edges")
	}
	rand.Seed(time.Now().Unix())
	G := NewGraph(V)
	set := treeset.NewWith(comparator)
	for G.E() < E {
		v := rand.Intn(V)
		w := rand.Intn(V)
		e := newEdge(v, w)
		if v != w && !set.Contains(e) {
			set.Add(e)
			G.AddEdge(v, w)
		}
	}
	return G, nil
}

// SimpleP returns a random simple graph on V vertices, with an edge between any two vertices with probability p.
func SimpleP(V int, p float64) (*Graph, error) {
	if p < 0.0 || p > 1.0 {
		return nil, errors.New("probability must be between 0 and 1")
	}
	rand.Seed(time.Now().Unix())
	G := NewGraph(V)
	for v := 0; v < V; v++ {
		for w := v + 1; w < V; w++ {
			if rand.Float64() < p {
				G.AddEdge(v, w)
			}
		}
	}
	return G, nil
}

// Complete returns the complete graph on V vertices.
// A complete graph is a simple undirected graph in which every pair of distinct vertices is connected by a unique edge.
func Complete(V int) *Graph {
	g, _ := SimpleP(V, 1.0)
	return g
}

// CompleteBipartite returns a complete bipartite graph on V1 and V2 vertices.
func CompleteBipartite(V1, V2 int) *Graph {
	g, _ := Bipartite(V1, V2, V1*V2)
	return g
}

// Bipartite returns a random simple bipartite graph on V1 and V2 vertices with E edges.
func Bipartite(V1, V2, E int) (*Graph, error) {
	if E > V1*V2 {
		return nil, errors.New("too many edges")
	}
	if E < 0 {
		return nil, errors.New("too few edges")
	}
	rand.Seed(time.Now().Unix())
	G := NewGraph(V1 + V2)
	vertices := createVertices(V1 + V2)
	set := treeset.NewWith(comparator)
	for G.E() < E {
		i := rand.Intn(V1)
		j := V1 + rand.Intn(V2)
		e := newEdge(vertices[i], vertices[j])
		if !set.Contains(e) {
			set.Add(e)
			G.AddEdge(vertices[i], vertices[j])
		}
	}
	return G, nil
}

// BipartiteP returns a random simple bipartite graph on V1 and V2 vertices, containing each possible edge with probability p.
func BipartiteP(V1, V2 int, p float64) (*Graph, error) {
	if p < 0.0 || p > 1.0 {
		return nil, errors.New("probability must be between 0 and 1")
	}
	rand.Seed(time.Now().Unix())
	vertices := createVertices(V1 + V2)
	G := NewGraph(V1 + V2)
	for i := 0; i < V1; i++ {
		for j := 0; j < V2; j++ {
			if rand.Float64() < p {
				G.AddEdge(vertices[i], vertices[j])
			}
		}
	}
	return G, nil
}

// Path returns a path graph on V vertices.
func Path(V int) *Graph {
	G := NewGraph(V)
	vertices := createVertices(V)
	for i := 0; i < V-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	return G
}

// BinaryTree returns a complete binary tree graph on V vertices.
func BinaryTree(V int) *Graph {
	G := NewGraph(V)
	vertices := createVertices(V)
	for i := 1; i < V; i++ {
		G.AddEdge(vertices[i], vertices[(i-1)/2])
	}
	return G
}

// Cycle returns a cycle graph on V vertices.
func Cycle(V int) *Graph {
	G := NewGraph(V)
	vertices := createVertices(V)
	for i := 0; i < V-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	G.AddEdge(vertices[V-1], vertices[0])
	return G
}

// EulerianCycle returns an Eulerian cycle graph on V vertices.
func EulerianCycle(V, E int) (*Graph, error) {
	if E <= 0 {
		return nil, errors.New("an Eulerian cycle must have at least one edge")
	}
	if V <= 0 {
		return nil, errors.New("an Eulerian cycle must have at least one vertex")
	}
	rand.Seed(time.Now().Unix())
	G := NewGraph(V)
	vertices := make([]int, E)
	for i := 0; i < E; i++ {
		vertices[i] = rand.Intn(V)
	}
	for i := 0; i < E-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	G.AddEdge(vertices[E-1], vertices[0])
	return G, nil
}

// EulerianPath returns an Eulerian path graph on V vertices.
func EulerianPath(V, E int) (*Graph, error) {
	if E < 0 {
		return nil, errors.New("negative number of edges")
	}
	if V <= 0 {
		return nil, errors.New("an Eulerian path must have at least one vertex")
	}
	rand.Seed(time.Now().Unix())
	G := NewGraph(V)
	vertices := make([]int, E+1)
	for i := 0; i < E+1; i++ {
		vertices[i] = rand.Intn(V)
	}
	for i := 0; i < E; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	return G, nil
}

// Wheel returns a wheel graph on V vertices.
func Wheel(V int) (*Graph, error) {
	if V <= 1 {
		return nil, errors.New("number of vertices must be at least 2")
	}
	G := NewGraph(V)
	vertices := createVertices(V)

	// simple cycle on V-1 vertices
	for i := 1; i < V-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	G.AddEdge(vertices[V-1], vertices[1])

	// connect vertices[0] to every vertex on cycle
	for i := 1; i < V; i++ {
		G.AddEdge(vertices[0], vertices[i])
	}
	return G, nil
}

// Star returns a star graph on V vertices.
func Star(V int) (*Graph, error) {
	if V <= 0 {
		return nil, errors.New("number of vertices must be at least 1")
	}
	G := NewGraph(V)
	vertices := createVertices(V)

	// connect vertices[0] to every other vertex
	for i := 1; i < V; i++ {
		G.AddEdge(vertices[0], vertices[i])
	}
	return G, nil
}

// Regular returns a uniformly random k-regular graph on V vertices (not necessarily simple).
func Regular(V, k int) (*Graph, error) {
	if V*k%2 != 0 {
		return nil, errors.New("number of vertices * k must be even")
	}
	rand.Seed(time.Now().Unix())
	G := NewGraph(V)
	vertices := make([]int, V*k)
	for v := 0; v < V; v++ {
		for j := 0; j < k; j++ {
			vertices[v+V*j] = v
		}
	}
	rand.Shuffle(V*k, func(i, j int) {
		vertices[i], vertices[j] = vertices[j], vertices[i]
	})
	for i := 0; i < V*k/2; i++ {
		G.AddEdge(vertices[2*i], vertices[2*i+1])
	}
	return G, nil
}

type vkey int

func (v vkey) CompareTo(k priorityqueue.Key) int {
	that := k.(vkey)
	if v < that {
		return -1
	} else if v > that {
		return 1
	} else {
		return 0
	}

}

// Tree returns a uniformly random tree on V vertices.
func Tree(V int) *Graph {
	G := NewGraph(V)
	if V == 1 {
		return G
	}
	// Cayley's theorem: there are V^(V-2) labeled trees on V vertices
	// Prufer sequence: sequence of V-2 values between 0 and V-1
	// Prufer's proof of Cayley's theorem: Prufer sequences are in 1-1
	// with labeled trees on V vertices
	rand.Seed(time.Now().Unix())
	prufer := make([]int, V-2)
	for i := 0; i < V-2; i++ {
		prufer[i] = rand.Intn(V)
	}
	// degree of vertex v = 1 + number of times it appers in Prufer sequence
	degree := make([]int, V)
	for v := 0; v < V; v++ {
		degree[v] = 1
	}
	for i := 0; i < V-2; i++ {
		degree[prufer[i]]++
	}

	pq := priorityqueue.NewMinPQ()
	for v := 0; v < V; v++ {
		if degree[v] == 1 {
			pq.Insert(vkey(v))
		}
	}
	// repeatedly delMin() degree 1 vertex that has the minimum index
	for i := 0; i < V-2; i++ {
		v, _ := pq.DelMin()
		G.AddEdge(int(v.(vkey)), prufer[i])
		degree[int(v.(vkey))]--
		degree[prufer[i]]--
		if degree[prufer[i]] == 1 {
			pq.Insert(vkey(prufer[i]))
		}
	}
	a, _ := pq.DelMin()
	b, _ := pq.DelMin()
	G.AddEdge(int(a.(vkey)), int(b.(vkey)))
	return G
}

func createVertices(capacity int) []int {
	rand.Seed(time.Now().Unix())
	vertices := make([]int, capacity)
	for i := 0; i < capacity; i++ {
		vertices[i] = i
	}
	rand.Shuffle(capacity, func(i, j int) {
		vertices[i], vertices[j] = vertices[j], vertices[i]
	})
	return vertices
}
