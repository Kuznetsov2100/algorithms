package priorityqueue

import (
	"math/rand"
	"testing"

	"github.com/handane123/algorithms/sort"
)

func TestMaxPQ_Insert(t *testing.T) {
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
	for _, x := range studentlist {
		pq.Insert(x)
	}
	if pq.IsEmpty() {
		t.Errorf("expect false, got %t", pq.IsEmpty())
	}
	if pq.Size() != len(studentlist) {
		t.Errorf("expect %d, got %d", len(studentlist), pq.Size())
	}
}

func TestMaxPQ_DelMax(t *testing.T) {
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
	for _, x := range studentlist {
		pq.Insert(x)
	}
	sort.QuickSort(studentlist)
	for i := pq.Size() - 1; !pq.IsEmpty(); i-- {
		if max, err := pq.DelMax(); err != nil {
			t.Fatal(err)
		} else {
			if stu, ok := max.(Student); !ok {
				t.Errorf("expect Student type, got %T", stu)
			} else {
				if stu.score != studentlist[i].score {
					t.Errorf("expect %d, got %d", studentlist[i].score, stu.score)
				}
			}
		}
	}
	if !pq.IsEmpty() {
		t.Errorf("expect ture, got %t", pq.IsEmpty())
	}
	if _, err := pq.DelMax(); err == nil {
		t.Error("should throw error when delete element from an empty max priority queue!")
	}
}

func benchmarkMaxPQInsert(b *testing.B, pq *MaxPQ, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			pq.Insert(Student{
				name:  randStringRunes(10),
				score: rand.Intn(10000),
			})
		}
	}
}

func benchmarkMaxPQDelMin(b *testing.B, pq *MaxPQ, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			//nolint:errcheck
			pq.DelMax()
		}
	}
}
func BenchmarkMaxPQDelMax10k(b *testing.B) {
	b.StopTimer()
	size := 10000
	pq := NewMaxPQ()
	for n := 0; n < size; n++ {
		pq.Insert(Student{
			name:  randStringRunes(10),
			score: rand.Intn(10000),
		})
	}
	b.StartTimer()
	benchmarkMaxPQDelMin(b, pq, size)
}

func BenchmarkMaxPQInsert10k(b *testing.B) {
	b.StopTimer()
	size := 10000
	pq := NewMaxPQ()
	b.StartTimer()
	benchmarkMaxPQInsert(b, pq, size)
}
