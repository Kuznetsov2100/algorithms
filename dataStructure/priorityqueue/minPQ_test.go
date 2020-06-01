package priorityqueue

import (
	"math/rand"
	"testing"

	"github.com/handane123/algorithms/sort"
)

type Class []Student

type Student struct {
	//nolint:structcheck
	name  string
	score int
}

func (s Class) Len() int {
	return len(s)
}

func (s Class) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s Class) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Class) Shuffle() {
	rand.Shuffle(s.Len(), s.Swap)
}

func (s Student) CompareTo(k Key) int {
	t := k.(Student)
	if s.score < t.score {
		return -1
	} else if s.score > t.score {
		return 1
	} else {
		return 0
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func benchmark_Insert(b *testing.B, pq *MinPQ, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			pq.Insert(Student{
				name:  randStringRunes(10),
				score: rand.Intn(1000),
			})
		}
	}
}

func Benchmark_Insert10k(b *testing.B) {
	b.StopTimer()
	size := 10000
	pq := NewMinPQ()
	b.StartTimer()
	benchmark_Insert(b, pq, size)
}

func Test_minPQ(t *testing.T) {
	studentlist := Class{
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

	pq := NewMinPQ()
	if _, err := pq.DelMin(); err == nil {
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

	min, _ := pq.DelMin()
	stu := min.(Student)
	sort.QuickSort(studentlist)
	if stu.score != studentlist[0].score {
		t.Errorf("expect %d, got %d", studentlist[0].score, stu.score)
	}
}
