package digraph

import (
	"math/rand"
	"time"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/pkg/errors"
)

// DigraphGenerator struct provides static methods for creating various digraphs,
// including Erdos-Renyi random digraphs, random DAGs, random rooted trees, random rooted DAGs,
// random tournaments, path digraphs, cycle digraphs, and the complete digraph.
type DigraphGenerator struct {
}

// NewDigraphGenerator constructs the DigraphGenerator struct
func NewDigraphGenerator() *DigraphGenerator {
	return &DigraphGenerator{}
}

type privateEdge struct {
	v int
	w int
}

func newprivateEdge(v, w int) *privateEdge {
	return &privateEdge{v: v, w: w}
}

func comparator(a, b interface{}) int {
	a1 := a.(*privateEdge)
	b1 := b.(*privateEdge)
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

// Simple returns a random simple digraph containing V vertices and E edges.
func (generator *DigraphGenerator) Simple(V, E int) (*Digraph, error) {
	if E > V*(V-1) {
		return nil, errors.New("too many edges")
	}
	if E < 0 {
		return nil, errors.New("too few edges")
	}

	rand.Seed(time.Now().Unix())
	G := NewDigraph(V)
	set := treeset.NewWith(comparator)
	for G.E() < E {
		v := rand.Intn(V)
		w := rand.Intn(V)
		e := newprivateEdge(v, w)
		if v != w && !set.Contains(e) {
			set.Add(e)
			G.AddEdge(v, w)
		}
	}
	return G, nil
}

// SimpleP returns a random simple digraph on V vertices, with an edge between any two vertices with probability p.
func (generator *DigraphGenerator) SimpleP(V int, p float64) (*Digraph, error) {
	if p < 0.0 || p > 1.0 {
		return nil, errors.New("probability must be between 0 and 1")
	}
	rand.Seed(time.Now().Unix())
	G := NewDigraph(V)
	for v := 0; v < V; v++ {
		for w := 0; w < V; w++ {
			if v != w {
				if rand.Float64() < p {
					G.AddEdge(v, w)
				}
			}
		}
	}
	return G, nil
}

// Complete returns the complete digraph on V vertices.
func (generator *DigraphGenerator) Complete(V int) *Digraph {
	G := NewDigraph(V)
	for v := 0; v < V; v++ {
		for w := 0; w < V; w++ {
			if v != w {
				G.AddEdge(v, w)
			}
		}
	}
	return G
}

// Dag returns a random simple DAG containing V vertices and E edges.
func (generator *DigraphGenerator) Dag(V, E int) (*Digraph, error) {
	if E > V*(V-1)/2 {
		return nil, errors.New("too many edges")
	}
	if E < 0 {
		return nil, errors.New("too few edges")
	}
	G := NewDigraph(V)
	set := treeset.NewWith(comparator)
	vertices := createVertices(V)

	for G.E() < E {
		v := rand.Intn(V)
		w := rand.Intn(V)
		e := newprivateEdge(v, w)
		if v < w && !set.Contains(e) {
			set.Add(e)
			G.AddEdge(vertices[v], vertices[w])
		}
	}
	return G, nil
}

// Tournament returns a random tournament digraph on V vertices.
func (generator *DigraphGenerator) Tournament(V int) *Digraph {
	G := NewDigraph(V)
	rand.Seed(time.Now().Unix())
	for v := 0; v < G.V(); v++ {
		for w := v + 1; w < G.V(); w++ {
			if rand.Float64() < 0.5 {
				G.AddEdge(v, w)
			} else {
				G.AddEdge(w, v)
			}
		}
	}
	return G
}

// CompleteRootedInDAG returns a complete rooted-in DAG on V vertices.
func (generator *DigraphGenerator) CompleteRootedInDAG(V int) *Digraph {
	G := NewDigraph(V)
	vertices := createVertices(V)
	for i := 0; i < V; i++ {
		for j := i + 1; j < V; j++ {
			G.AddEdge(vertices[i], vertices[j])
		}
	}
	return G
}

// RootedInDAG returns a random rooted-in DAG on V vertices and E edges.
func (generator *DigraphGenerator) RootedInDAG(V, E int) (*Digraph, error) {
	if E > V*(V-1)/2 {
		return nil, errors.New("too many edges")
	}
	if E < V-1 {
		return nil, errors.New("too few edges")
	}
	G := NewDigraph(V)
	set := treeset.NewWith(comparator)
	vertices := createVertices(V)
	rand.Seed(time.Now().Unix())
	for v := 0; v < V-1; v++ {
		w := rand.Intn(V-v-1) + v + 1
		e := newprivateEdge(v, w)
		set.Add(e)
		G.AddEdge(vertices[v], vertices[w])
	}
	for G.E() < E {
		v := rand.Intn(V)
		w := rand.Intn(V)
		e := newprivateEdge(v, w)
		if v < w && !set.Contains(e) {
			set.Add(e)
			G.AddEdge(vertices[v], vertices[w])
		}
	}
	return G, nil
}

// CompleteRootedOutDAG returns a complete rooted-out DAG on V vertices.
func (generator *DigraphGenerator) CompleteRootedOutDAG(V int) *Digraph {
	G := NewDigraph(V)
	vertices := createVertices(V)
	for i := 0; i < V; i++ {
		for j := i + 1; j < V; j++ {
			G.AddEdge(vertices[j], vertices[i])
		}
	}
	return G
}

// RootedOutDAG returns a random rooted-out DAG on V vertices and E edges.
func (generator *DigraphGenerator) RootedOutDAG(V, E int) (*Digraph, error) {
	if E > V*(V-1)/2 {
		return nil, errors.New("too many edges")
	}
	if E < V-1 {
		return nil, errors.New("too few edges")
	}
	G := NewDigraph(V)
	set := treeset.NewWith(comparator)
	// fix a topological order
	vertices := createVertices(V)
	rand.Seed(time.Now().Unix())

	// one edge pointing from each vertex, other than the root = vertices[V-1]
	for v := 0; v < V-1; v++ {
		w := rand.Intn(V-v-1) + v + 1
		e := newprivateEdge(w, v)
		set.Add(e)
		G.AddEdge(vertices[w], vertices[v])
	}
	for G.E() < E {
		v := rand.Intn(V)
		w := rand.Intn(V)
		e := newprivateEdge(w, v)
		if v < w && !set.Contains(e) {
			set.Add(e)
			G.AddEdge(vertices[w], vertices[v])
		}
	}
	return G, nil
}

// RootedInTree returns a random rooted-in tree on V vertices.
func (generator *DigraphGenerator) RootedInTree(V int) (*Digraph, error) {
	return generator.RootedInDAG(V, V-1)
}

// RootedOutTree returns a random rooted-out tree on V vertices.
func (generator *DigraphGenerator) RootedOutTree(V int) (*Digraph, error) {
	return generator.RootedOutDAG(V, V-1)
}

// PathDigraph returns a path digraph on V vertices.
func (generator *DigraphGenerator) PathDigraph(V int) *Digraph {
	G := NewDigraph(V)
	vertices := createVertices(V)
	for i := 0; i < V-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	return G
}

// BinaryTree returns a complete binary tree digraph on V vertices.
func (generator *DigraphGenerator) BinaryTree(V int) *Digraph {
	G := NewDigraph(V)
	vertices := createVertices(V)
	for i := 1; i < V; i++ {
		G.AddEdge(vertices[i], vertices[(i-1)/2])
	}
	return G
}

// CycleDigraph returns a cycle digraph on V vertices.
func (generator *DigraphGenerator) CycleDigraph(V int) *Digraph {
	G := NewDigraph(V)
	vertices := createVertices(V)
	for i := 0; i < V-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	G.AddEdge(vertices[V-1], vertices[0])
	return G
}

// EulerianCycleDigraph returns an Eulerian cycle digraph on V vertices.
func (generator *DigraphGenerator) EulerianCycleDigraph(V, E int) (*Digraph, error) {
	if E <= 0 {
		return nil, errors.New("an Eulerian cycle must at least one edge")
	}
	if V <= 0 {
		return nil, errors.New("an Eulerian cycle must at least one vertex")
	}
	G := NewDigraph(V)
	rand.Seed(time.Now().Unix())
	vertices := make([]int, E)
	for i := range vertices {
		vertices[i] = rand.Intn(V)
	}
	for i := 0; i < E-1; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	G.AddEdge(vertices[E-1], vertices[0])
	return G, nil
}

// EulerianPathDigraph returns an Eulerian path digraph on V vertices.
func (generator *DigraphGenerator) EulerianPathDigraph(V, E int) (*Digraph, error) {
	if E < 0 {
		return nil, errors.New("negative number of edges")
	}
	if V <= 0 {
		return nil, errors.New("an Eulerian path must have at least one vertex")
	}
	G := NewDigraph(V)
	rand.Seed(time.Now().Unix())
	vertices := make([]int, E+1)
	for i := range vertices {
		vertices[i] = rand.Intn(V)
	}
	for i := 0; i < E; i++ {
		G.AddEdge(vertices[i], vertices[i+1])
	}
	return G, nil
}

// StrongDigraph returns a random simple digraph on V vertices, E edges and (at most) c strong components.
func (generator *DigraphGenerator) StrongDigraph(V, E, c int) (*Digraph, error) {
	if c >= V || c <= 0 {
		return nil, errors.New("number of components must be between 1 and V")
	}
	if E <= 2*(V-c) {
		return nil, errors.New("number of edges must be at least 2*(V-c)")
	}
	if E > V*(V-1)/2 {
		return nil, errors.New("too many edges")
	}
	rand.Seed(time.Now().Unix())
	// the digraph
	G := NewDigraph(V)
	// edges added to G (to avoid duplicate edges)
	set := treeset.NewWith(comparator)
	label := make([]int, V)
	for i := range label {
		label[i] = rand.Intn(c)
	}

	// make all vertices with label c a strong component by
	// combining a rooted in-tree and a rooted out-tree
	for i := 0; i < c; i++ {
		count := 0
		for v := 0; v < G.V(); v++ {
			if label[v] == i {
				count++
			}
		}

		vertices := make([]int, count)
		j := 0
		for v := 0; v < V; v++ {
			if label[v] == i {
				vertices[j] = v
				j++
			}
		}
		rand.Shuffle(len(vertices), func(i, j int) {
			vertices[i], vertices[j] = vertices[j], vertices[i]
		})

		// rooted-in tree with root = vertices[count-1]
		for v := 0; v < count-1; v++ {
			w := rand.Intn(count-v-1) + v + 1
			e := newprivateEdge(w, v)
			set.Add(e)
			G.AddEdge(vertices[w], vertices[v])
		}

		// rooted-out tree with root = vertices[count-1]
		for v := 0; v < count-1; v++ {
			w := rand.Intn(count-v-1) + v + 1
			e := newprivateEdge(v, w)
			set.Add(e)
			G.AddEdge(vertices[v], vertices[w])
		}
	}

	for G.E() < E {
		v := rand.Intn(V)
		w := rand.Intn(V)
		e := newprivateEdge(v, w)
		if !set.Contains(e) && v != w && label[v] <= label[w] {
			set.Add(e)
			G.AddEdge(v, w)
		}
	}
	return G, nil
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
