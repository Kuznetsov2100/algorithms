package digraph

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolDigraph(t *testing.T) {
	routestring := "JFK MCO\n" +
		"ORD DEN\n" +
		"ORD HOU\n" +
		"DFW PHX\n" +
		"JFK ATL\n" +
		"ORD DFW\n" +
		"ORD PHX\n" +
		"ATL HOU\n" +
		"DEN PHX\n" +
		"PHX LAX\n" +
		"JFK ORD\n" +
		"DEN LAS\n" +
		"DFW HOU\n" +
		"ORD ATL\n" +
		"LAS LAX\n" +
		"ATL MCO\n" +
		"HOU MCO\n" +
		"LAS PHX\n"

	content := []byte(routestring)
	tmpfile, err := ioutil.TempFile("", "routes.*.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		t.Error(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Error(err)
	}
	sg := NewSymbolDigraph(tmpfile.Name(), " ")
	assert := assert.New(t)
	assert.True(sg.Contains("JFK"))

	s := sg.IndexOf("JFK")
	names := []string{"ORD", "ATL", "MCO"}

	g := sg.Graph()
	for i, v := range g.Adj(s) {
		assert.Equal(names[i], sg.NameOf(v))
	}

	assert.Equal(-1, sg.IndexOf("love"))
	assert.Panics(func() { sg.NameOf(20) })
}
