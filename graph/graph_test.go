package graph

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
)

func testPanic(t *testing.T, msg string) {
	if r := recover(); r == nil {
		t.Errorf("%s did not panic", msg)
	}
}
func TestGraph_NewGraph(t *testing.T) {
	defer testPanic(t, "negative vertices")
	g := NewGraph(-1)
	if d := g.Degree(1); d > 0 {
		t.Error("should be zero")
	}
}

func TestGraph_NewGraphIn(t *testing.T) {
	buf := strings.NewReader("2\n" + "1\n" + "0 1\n")
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: s}
	g, err := NewGraphIn(in)
	if err != nil {
		t.Error(err)
	}
	if e := g.E(); e != 1 {
		t.Errorf("expect 1, got %d", e)
	}
	if v := g.V(); v != 2 {
		t.Errorf("expect 2, got %d", v)
	}
	if _, err := NewGraphIn(nil); err == nil {
		t.Error("should throw error: argument is nil")
	}
	buf1 := strings.NewReader("-1\n" + "1\n" + "0 1\n")
	s1 := bufio.NewScanner(buf1)
	s1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: s1}
	if _, err := NewGraphIn(in1); err == nil {
		t.Error("should throw error: number of vertices in a Graph must be non negative")
	}
	buf2 := strings.NewReader("2\n" + "-2\n" + "0 1\n")
	s2 := bufio.NewScanner(buf2)
	s2.Split(bufio.ScanWords)
	in2 := &stdin.In{Scanner: s2}
	if _, err := NewGraphIn(in2); err == nil {
		t.Error("should throw error: number of edges in a Graph must be non negative")
	}
}

func TestGraph_Degree(t *testing.T) {
	defer testPanic(t, "invalid v")
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	if d := g.Degree(0); d != 2 {
		t.Errorf("expect 2, got %d", d)
	}
	if e := g.E(); e != 2 {
		t.Errorf("expect 2, got %d", e)
	}
	if v := g.V(); v != 3 {
		t.Errorf("expect 3, got %d", v)
	}
	g.AddEdge(0, 4)
}

func TestGraph_Adj(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	if vertices := g.Adj(0); !reflect.DeepEqual(vertices, []int{2, 1}) {
		t.Errorf("expect int[]{1,2}, got %+v", vertices)
	}
}

func TestGraph_String(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	if s := g.String(); s != "3 vertices, 2 edges \n"+"0:  2 1\n"+"1:  0\n"+"2:  0\n" {
		t.Errorf("expect fuck, got %s", s)
	}
}
