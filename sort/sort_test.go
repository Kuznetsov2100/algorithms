package sort

import (
	"math/rand"
	"testing"
	"time"
)

// type class implement Len(), Less(), Swap() Shuffle() for Comparable interface
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

func (s class) Shuffle() {
	rand.Shuffle(s.Len(), s.Swap)
}

// type IntSlice implement Len(), Less(), Swap() Shuffle() for Comparable interface
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

func (p IntSlice) Shuffle() {
	rand.Shuffle(p.Len(), p.Swap)
}

func Test_BubbleSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	BubbleSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

	// to test ischanged = false
	sortedNumbers := IntSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSort(sortedNumbers)
	if !IsSorted(sortedNumbers) {
		t.Errorf("expect true, got %t", IsSorted(sortedNumbers))
	}
	numbers := createNumbers(100)
	QuickSort(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
	}
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

	numbers := createNumbers(100)
	InsertionSort(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
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

	numbers := createNumbers(100)
	SelectionSort(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
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

	numbers := createNumbers(100)
	ShellSort(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
	}

}

func Test_MergeSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	MergeSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

	numbers := createNumbers(100)
	MergeSort(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
	}

}
func createNumbers(size int) IntSlice {
	numbers := make(IntSlice, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
}
func Test_QuickSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	QuickSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

	numbers := createNumbers(100)
	QuickSort(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
	}
}

func Benchmark_BubbleSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		BubbleSort(numbers)
		b.StopTimer()
	}
}

func Benchmark_InsertionSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		InsertionSort(numbers)
		b.StopTimer()
	}
}
func Benchmark_SelectionSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		SelectionSort(numbers)
		b.StopTimer()
	}
}

func Benchmark_ShellSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		ShellSort(numbers)
		b.StopTimer()
	}
}

func Benchmark_MergeSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		MergeSort(numbers)
		b.StopTimer()
	}
}

func Benchmark_QuickSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		QuickSort(numbers)
		b.StopTimer()
	}
}
