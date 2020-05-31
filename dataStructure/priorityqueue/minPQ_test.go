package priorityqueue

import "testing"

type student struct {
	//nolint:structcheck
	name  string
	score int
}

func (s student) CompareTo(k Key) int {
	t := k.(student)
	if s.score < t.score {
		return -1
	} else if s.score > t.score {
		return 1
	} else {
		return 0
	}
}

func Test(t *testing.T) {
	studentlist := []student{
		{name: "xiaoming", score: 120},
		{name: "meimei", score: 110},
		{name: "daniu", score: 115},
		{name: "lili", score: 80},
		{name: "fafa", score: 90},
		{name: "hehe", score: 91},
		{name: "ganggang", score: 93},
		{name: "jiji", score: 102},
		{name: "jiajia", score: 103},
		{name: "mimi", score: 81},
		{name: "weiwei", score: 85},
	}

	pq := NewPQ()
	if _, err := pq.DelMin(); err == nil {
		t.Error("should throw error when delete element from an empty minPQ")
	}

	for _, x := range studentlist {
		pq.Insert(x)
	}
	if pq.IsEmpty() {
		t.Errorf("expect false, got %t", pq.IsEmpty())
	}
	if pq.Size() != 11 {
		t.Errorf("expect 11, got %d", pq.Size())
	}

	min, _ := pq.DelMin()
	fuck := min.(student)
	if fuck.score != 80 {
		t.Errorf("expect 80, got %d", fuck.score)
	}
}
