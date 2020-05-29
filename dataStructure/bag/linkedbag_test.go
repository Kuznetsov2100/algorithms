package linkedbag

import "testing"

func TestBag_Add(t *testing.T) {
	b := New()
	if !b.IsEmpty() {
		t.Errorf("expect true, got %t", b.IsEmpty())
	}
	for i := 1; i <= 10; i++ {
		b.Add(i)
	}

	if b.Size() != 10 {
		t.Errorf("expect 10, got %d", b.Size())
	}

}

func benchmarkAdd(b *testing.B, bag *Bag, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			bag.Add(n)
		}
	}
}

func BenchmarkBagAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	bag := New()
	b.StartTimer()
	benchmarkAdd(b, bag, size)
}
