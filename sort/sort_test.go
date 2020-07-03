package sort

import (
	"math/rand"
	"testing"
)

// type class implement Len(), Less(), Swap(), Shuffle() for Comparable interface
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
	rand.Seed(23)
	rand.Shuffle(s.Len(), s.Swap)
}

// type intSlice implement Len(), Less(), Swap() Shuffle() for Comparable interface
type intSlice []int

func (p intSlice) Len() int {
	return len(p)
}

func (p intSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p intSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p intSlice) Shuffle() {
	rand.Seed(23)
	rand.Shuffle(p.Len(), p.Swap)
}

func createNumbers(size int) intSlice {
	numbers := make(intSlice, size)
	rand.Seed(23)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Int()
	}
	return numbers
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
	sortedNumbers := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
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

func Test_QuickSort2Way(t *testing.T) {
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
	QuickSort2Way(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
	}
}

func Test_QuickSort3Way(t *testing.T) {
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
	QuickSort3Way(numbers)
	if !IsSorted(numbers) {
		t.Errorf("expect true, got %t", IsSorted(numbers))
	}

	numbers1 := intSlice{10, 4, 3, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 5, 5, 5, 6, 7, 9, 8}
	QuickSort3Way(numbers1)
	if !IsSorted(numbers1) {
		t.Errorf("expect true, got %t", IsSorted(numbers1))
	}
}

func Test_HeapSort(t *testing.T) {
	s := class{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	if IsSorted(s) {
		t.Errorf("expect false, got %t", IsSorted(s))
	}
	HeapSort(s)
	if !IsSorted(s) {
		t.Errorf("expect true, got %t", IsSorted(s))
	}

	numbers := createNumbers(100)
	HeapSort(numbers)
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

func Benchmark_QuickSort2Way_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		QuickSort2Way(numbers)
		b.StopTimer()
	}
}

func Benchmark_QuickSort3Way_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		QuickSort3Way(numbers)
		b.StopTimer()
	}
}

func Benchmark_HeapSort_10k(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		numbers := createNumbers(10000)
		b.StartTimer()
		HeapSort(numbers)
		b.StopTimer()
	}
}
