package searching

import (
	"testing"
)

func TestSequentialSearchST_Put(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSequentialSearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if bst.Size() != 10 {
		t.Errorf("expect 10, got %d", bst.Size())
	}
	if bst.IsEmpty() {
		t.Error("should not be empty")
	}
	if err := bst.Put(nil, 2); err == nil {
		t.Error("should throw error: argument to Put() is nil")
	}
	if err := bst.Put(words("L"), nil); err != nil {
		t.Error(err)
	}
	if err := bst.Put(words("S"), 2); err != nil {
		t.Error(err)
	} else {
		if val, err := bst.Get(words("S")); err != nil {
			t.Error(err)
		} else {
			if val != 2 {
				t.Errorf("expect 2, got %d", val)
			}
		}
	}
}

func TestSequentialSearchST_Contains(t *testing.T) {
	bst := NewSequentialSearchST()
	if _, err := bst.Contains(nil); err == nil {
		t.Error("should throw error: argument to Contains() is nil")
	}
	if ok, err := bst.Contains(words("W")); err != nil {
		t.Error(err)
	} else {
		if ok {
			t.Errorf("expect false, got %t", ok)
		}
	}
}

func TestSequentialSearchST_Get(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSequentialSearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if _, err := bst.Get(nil); err == nil {
		t.Error("should throw error: argument to Get() is nil")
	}
	if val, _ := bst.Get(words("B")); val != nil {
		t.Errorf("expect nil, got %v", val)
	}

}

func TestSequentialSearchST_Delete(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSequentialSearchST()
	if err := bst.Delete(words("A")); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if err := bst.Delete(nil); err == nil {
		t.Error("should throw error: argument to Delete() is nil")
	}

	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Delete(tinyST[i])
		if ok, _ := bst.Contains(tinyST[i]); ok {
			t.Errorf("expect false, got %t", ok)
		}
	}
}

func TestSequentialSearchST_Keys(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSequentialSearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if length := len(bst.Keys()); length != 10 {
		t.Errorf("expect 10,got %d", length)
	}
}
