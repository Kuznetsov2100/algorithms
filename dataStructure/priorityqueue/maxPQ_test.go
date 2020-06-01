package priorityqueue

import (
	"testing"

	"github.com/handane123/algorithms/sort"
)

func Test_maxPQ(t *testing.T) {
	studentlist := Class{
		{name: "daniu", score: 115},
		{name: "lili", score: 80},
		{name: "fafa", score: 90},
		{name: "hehe", score: 91},
		{name: "ganggang", score: 93},
		{name: "jiji", score: 102},
		{name: "jiajia", score: 103},
		{name: "mimi", score: 81},
		{name: "meimei", score: 110},
		{name: "weiwei", score: 85},
		{name: "xiaoming", score: 120},
	}

	pq := NewMaxPQ()
	if _, err := pq.DelMax(); err == nil {
		t.Error("should throw error when delete element from an empty minPQ")
	}

	for _, x := range studentlist {
		pq.Insert(x)
	}
	if pq.IsEmpty() {
		t.Errorf("expect false, got %t", pq.IsEmpty())
	}
	if pq.Size() != len(studentlist) {
		t.Errorf("expect %d, got %d", len(studentlist), pq.Size())
	}

	max, _ := pq.DelMax()
	fuck := max.(Student)
	sort.QuickSort(studentlist)
	if fuck.score != studentlist[len(studentlist)-1].score {
		t.Errorf("expect %d, got %d", studentlist[len(studentlist)-1].score, fuck.score)
	}
}
