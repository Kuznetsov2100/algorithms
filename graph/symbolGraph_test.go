package graph

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolGraph(t *testing.T) {
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
	sg := NewSymbolGraph(tmpfile.Name(), " ")
	assert := assert.New(t)
	ok, err1 := sg.Contains("JFK")
	assert.Nil(err1)
	assert.True(ok)

	s, err2 := sg.IndexOf("JFK")
	names := []string{"ORD", "ATL", "MCO"}
	assert.Nil(err2)

	g := sg.Graph()
	for i, v := range g.Adj(s) {
		assert.Equal(names[i], sg.NameOf(v))
	}

	index, _ := sg.IndexOf("love")
	assert.Equal(-1, index)

	assert.Panics(func() { sg.NameOf(20) })
}
