package linkedstack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := New()
	if !s.IsEmpty() {
		t.Errorf("s.IsEmpty() = %t; want true", s.IsEmpty())
	}

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	if s.Size() != 10 {
		t.Errorf("s.Size() = %d; want 10", s.Size())
	}

}

func TestStack_Pop(t *testing.T) {
	s := New()
	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	for i := 10; i >= 1; i-- {
		if element, err := s.Pop(); err != nil {
			t.Error(err)
		} else {
			if element != i {
				t.Errorf("element = %d; want %d", element, i)
			}
		}
	}
	if !s.IsEmpty() {
		t.Errorf("s.IsEmpty() = %t; want true", s.IsEmpty())
	}
	if s.Size() != 0 {
		t.Errorf("s.Size() = %d; want 0", s.Size())
	}

	if _, err := s.Pop(); err == nil {
		t.Error("should throw error when pop an empty stack!")
	}

}

func TestStack_Peek(t *testing.T) {

	s := New()
	s.Push(1)
	if element, err := s.Peek(); err != nil {
		t.Error(err)
	} else {
		if element != 1 {
			t.Errorf("s.Peek() = %d; want 1", element)
		}
	}

	//nolint:errcheck
	s.Pop()

	if _, err := s.Peek(); err == nil {
		t.Error("should throw error when peek an empty stack!")
	}

}

func TestStack_Values(t *testing.T) {
	s := New()
	for i := 1; i <= 3; i++ {
		s.Push(i)
	}
	got := s.Values()
	var expect []interface{}
	for !s.IsEmpty() {
		val, _ := s.Pop()
		expect = append(expect, val)
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expect: %v, Got: %v", expect, got)
	}
}

func benchmarkPush(b *testing.B, stack *Stack, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			stack.Push(n)
		}
	}
}

func benchmarkPop(b *testing.B, stack *Stack, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			//nolint:errcheck
			stack.Pop()
		}
	}
}

func BenchmarkStackPush100(b *testing.B) {
	b.StopTimer()
	size := 100
	s := New()
	b.StartTimer()
	benchmarkPush(b, s, size)
}

func BenchmarkStackPop100(b *testing.B) {
	b.StopTimer()
	size := 100
	s := New()
	for i := 0; i < size; i++ {
		s.Push(i)
	}
	b.StartTimer()
	benchmarkPop(b, s, size)
}
