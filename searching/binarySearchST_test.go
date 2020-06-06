package searching

import (
	"testing"
)

type words string

func (w words) CompareTo(k Key) int {
	t := k.(words)
	if w < t {
		return -1
	} else if w > t {
		return 1
	} else {
		return 0
	}
}
func TestBinarySearchST_Put(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if bst.Size() != 10 {
		t.Errorf("expect 10, got %d", bst.Size())
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

func TestBinarySearchST_Contains(t *testing.T) {
	bst := NewBinarySearchST()
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

func TestBinarySearchST_Get(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
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

func TestBinarySearchST_Rank(t *testing.T) {
	bst := NewBinarySearchST()
	if _, err := bst.Rank(nil); err == nil {
		t.Error("should throw error: argument to Rank() is nil")
	}
}

func TestBinarySearchST_Delete(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
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
	if err := bst.Delete(words("W")); err != nil {
		t.Error(err)
	}
}

func TestBinarySearchST_DeleteMin(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
	if err := bst.DeleteMin(); err == nil {
		t.Error("should throw error: symbol table underflow")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if err := bst.DeleteMin(); err != nil {
		t.Error(err)
	}
}

func TestBinarySearchST_DeleteMax(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
	if err := bst.DeleteMax(); err == nil {
		t.Error("should throw error: symbol table underflow")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if err := bst.DeleteMax(); err != nil {
		t.Error(err)
	}
}

func TestBinarySearchST_Min(t *testing.T) {
	bst := NewBinarySearchST()
	if _, err := bst.Min(); err == nil {
		t.Error("should throw error: called Min() with empty symbol table")
	}
}

func TestBinarySearchST_Max(t *testing.T) {
	bst := NewBinarySearchST()
	if _, err := bst.Max(); err == nil {
		t.Error("should throw error: called Max() with empty symbol table")
	}
}

func TestBinarySearchST_Select(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if _, err := bst.Select(-1); err == nil {
		t.Errorf("should throw error: called Select() with invalid argument")
	}
	// 0-indexed
	if k, err := bst.Select(1); err != nil {
		t.Error(err)
	} else {
		if k.CompareTo(words("C")) != 0 {
			t.Errorf("expect key:\"C\", got %s", k.(words))
		}
	}
}

func TestBinarySearchST_Floor(t *testing.T) {
	tinyST := []words{"R", "C", "H"}
	bst := NewBinarySearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if _, err := bst.Floor(nil); err == nil {
		t.Error("should throw error: argument to Floor() is nil")
	}
	if k, _ := bst.Floor(words("C")); k.CompareTo(words("C")) != 0 {
		t.Errorf("expect key:\"C\", got %s", k.(words))
	}
	if k, _ := bst.Floor(words("B")); k != nil {
		t.Errorf("expect nil, got %s", k.(words))
	}
	if k, _ := bst.Floor(words("D")); k.CompareTo(words("C")) != 0 {
		t.Errorf("expect key:\"C\", got %s", k.(words))
	}
}

func TestBinarySearchST_Ceiling(t *testing.T) {
	tinyST := []words{"R", "C", "H"}
	bst := NewBinarySearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if _, err := bst.Ceiling(nil); err == nil {
		t.Error("should throw error: argument to Ceiling() is nil")
	}
	if k, _ := bst.Ceiling(words("W")); k != nil {
		t.Errorf("expect nil, got %s", k.(words))
	}
	if k, _ := bst.Ceiling(words("B")); k.CompareTo(words("C")) != 0 {
		t.Errorf("expect key:\"C\", got %s", k.(words))
	}
}

func TestBinarySearchST_SizeOf(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	bst := NewBinarySearchST()
	if _, err := bst.SizeOf(nil, words("F")); err == nil {
		t.Error("should throw error: first argument to SizeOf() is nil")
	}
	if _, err := bst.SizeOf(words("F"), nil); err == nil {
		t.Error("should throw error: second argument to SizeOf() is nil")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}
	if i, _ := bst.SizeOf(words("R"), words("A")); i != 0 {
		t.Errorf("expect 0, got %d", i)
	}
	if i, _ := bst.SizeOf(words("A"), words("C")); i != 2 {
		t.Errorf("expect 2, got %d", i)
	}
	if i, _ := bst.SizeOf(words("X"), words("Z")); i != 1 {
		t.Errorf("expect 1, got %d", i)
	}
}

func TestBinarySearchST_Keys(t *testing.T) {
	tinyST := []words{"A", "B", "C"}
	bst := NewBinarySearchST()
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		bst.Put(tinyST[i], i)
	}

	keys := bst.Keys()
	for index, k := range keys {
		if k.(words) != tinyST[index] {
			t.Errorf("expect %s, got %s", tinyST[index], k.(words))
		}
	}
}
