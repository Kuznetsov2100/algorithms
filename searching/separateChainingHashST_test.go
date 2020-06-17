package searching

import (
	"testing"
)

func testPanic(t *testing.T, msg string) {
	if r := recover(); r == nil {
		t.Errorf("%s did not panic", msg)
	}
}

func TestSeparateChainingHashST_NewSeparateChainingHashST(t *testing.T) {
	defer testPanic(t, "negative capacity")
	bst := NewSeparateChainingHashST(-1)
	if !bst.IsEmpty() {
		t.Error("should be empty")
	}
}

func TestSeparateChainingHashST_Put(t *testing.T) {
	tinyST := []StringHashKey{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSeparateChainingHashST(1)
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
	if err := bst.Put(StringHashKey("L"), nil); err != nil {
		t.Errorf("%+v\n", err)
	}
	if err := bst.Put(StringHashKey("S"), 2); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if val, err := bst.Get(StringHashKey("S")); err != nil {
			t.Errorf("%+v\n", err)
		} else {
			if val != 2 {
				t.Errorf("expect 2, got %d", val)
			}
		}
	}
}

func TestSeparateChainingHashST_Contains(t *testing.T) {
	bst := NewSeparateChainingHashST(5)
	if _, err := bst.Contains(nil); err == nil {
		t.Error("should throw error: argument to Contains() is nil")
	}
	if ok, err := bst.Contains(StringHashKey("W")); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if ok {
			t.Errorf("expect false, got %t", ok)
		}
	}
}

func TestSeparateChainingHashST_Get(t *testing.T) {
	tinyST := []StringHashKey{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSeparateChainingHashST(5)
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if _, err := bst.Get(nil); err == nil {
		t.Error("should throw error: argument to Get() is nil")
	}
	if val, _ := bst.Get(StringHashKey("B")); val != nil {
		t.Errorf("expect nil, got %v", val)
	}

}

func TestSeparateChainingHashST_Delete(t *testing.T) {
	tinyST := []StringHashKey{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSeparateChainingHashST(5)
	if err := bst.Delete(StringHashKey("A")); err != nil {
		t.Errorf("%+v\n", err)
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

func TestSeparateChainingHashST_Keys(t *testing.T) {
	tinyST := []StringHashKey{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewSeparateChainingHashST(0)
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if length := len(bst.Keys()); length != 10 {
		t.Errorf("expect 10,got %d", length)
	}
}
