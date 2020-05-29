package sort

import (
	"math/rand"
	"testing"
)

// type class implement Len(), Less(), Swap() for Comparable interface
type class []student

type student struct {
	//nolint:structcheck
	name  string
	score int
}

func (s class) Len() int {
	return len(s)
}

func (s class) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s class) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// type IntSlice implement Len(), Less(), Swap() for Comparable interface
type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Test_InsertionSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	InsertionSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

}

func Test_SelectionSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	SelectionSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

}

func Test_ShellSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	ShellSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

}

func Benchmark_InsertionSort(b *testing.B) {
	b.StopTimer()
	var numbers IntSlice
	for i := 0; i < 10000; i++ {
		numbers = append(numbers, rand.Int())
	}
	b.StartTimer()
	InsertionSort(numbers)
	b.StopTimer()
}
func Benchmark_SelectionSort(b *testing.B) {
	b.StopTimer()
	var numbers IntSlice
	for i := 0; i < 10000; i++ {
		numbers = append(numbers, rand.Int())
	}
	b.StartTimer()
	SelectionSort(numbers)
	b.StopTimer()
}

func Benchmark_ShellSort(b *testing.B) {
	b.StopTimer()
	var numbers IntSlice
	for i := 0; i < 10000; i++ {
		numbers = append(numbers, rand.Int())
	}
	b.StartTimer()
	ShellSort(numbers)
	b.StopTimer()
}
