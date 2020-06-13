package graph

import "testing"

func TestGraphGenerator_Simple(t *testing.T) {
	if g, err := Simple(10, 5); err != nil {
		t.Error(err)
	} else {
		if v, e := g.V(), g.E(); v != 10 || e != 5 {
			t.Error("fuck")
		}
	}
}
