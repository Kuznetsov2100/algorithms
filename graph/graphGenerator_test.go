package graph

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGraphGenerator_Simple(t *testing.T) {
	assert := assert.New(t)
	g, err := Simple(3, 4)
	assert.Nil(g)
	assert.Error(err)

	g1, err1 := Simple(3, -1)
	assert.Nil(g1)
	assert.Error(err1)

	rand.Seed(time.Now().Unix())
	V := rand.Intn(10)
	E := rand.Intn(V*(V-1)/2 + 1)
	g2, err2 := Simple(V, E)
	assert.Nil(err2)

	for i := 0; i < V; i++ {
		for _, x := range g2.Adj(i) {
			if i == x {
				assert.Fail("not a simple graph")
			}
		}
	}
}

func TestGraphGenerator_SimpleP(t *testing.T) {
	assert := assert.New(t)
	g, err := SimpleP(3, 1.2)
	assert.Nil(g)
	assert.Error(err)

	rand.Seed(time.Now().Unix())
	V := rand.Intn(10)
	p := rand.Float64()
	g1, err1 := SimpleP(V, p)
	assert.Nil(err1)

	for i := 0; i < V; i++ {
		for _, x := range g1.Adj(i) {
			if i == x {
				assert.Fail("not a simple graph")
			}
		}
	}
}

func TestGraphGenerator_Complete(t *testing.T) {
	assert := assert.New(t)
	rand.Seed(time.Now().Unix())
	V := rand.Intn(10)
	g := Complete(V)
	for i := 0; i < V; i++ {
		assert.Equal(V-1, g.Degree(i))
	}
}

func TestGraphGenerator_Bipartite(t *testing.T) {

}
