package arrayqueue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Errorf("expect true, got %t", q.IsEmpty())
	}

	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}

	if q.Size() != 10 {
		t.Errorf("expect 10, got %d", q.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := New()
	for i := 1; i <= 10; i++ {
		q.Enqueue(i)
	}
	for i := 1; i <= 10; i++ {
		if element, err := q.Dequeue(); err != nil {
			t.Error(err)
		} else {
			if element != i {
				t.Errorf("expect %d, got %d", i, element)
			}
		}
	}

	if !q.IsEmpty() {
		t.Errorf("expect true, got %t", q.IsEmpty())
	}

	if q.Size() != 0 {
		t.Errorf("expect 0, got %d", q.Size())
	}

	if _, err := q.Dequeue(); err == nil {
		t.Error("should throw error when dequeue an empty queue!")
	}

}

func TestQueue_Peek(t *testing.T) {
	q := New()
	q.Enqueue(1)
	if element, err := q.Peek(); err != nil {
		t.Error(err)
	} else {
		if element != 1 {
			t.Errorf("expect 1, got %d", element)
		}
	}
	//nolint:errcheck
	q.Dequeue()
	if _, err := q.Peek(); err == nil {
		t.Error("should throw error when peek an empty queue!")
	}
}

func benchmarkEnqueue(b *testing.B, q *Queue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			q.Enqueue(n)
		}
	}
}

func benchmarkDequeue(b *testing.B, q *Queue, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			//nolint:errcheck
			q.Dequeue()
		}
	}
}

func BenchmarkStackEnqueue100(b *testing.B) {
	b.StopTimer()
	size := 100
	q := New()
	b.StartTimer()
	benchmarkEnqueue(b, q, size)
}

func BenchmarkStackDequeue100(b *testing.B) {
	b.StopTimer()
	size := 100
	q := New()
	for i := 0; i < size; i++ {
		q.Enqueue(i)
	}
	b.StartTimer()
	benchmarkDequeue(b, q, size)
}
