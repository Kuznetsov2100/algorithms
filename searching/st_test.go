package searching

import (
	"testing"
)

func wordsComparator(a, b interface{}) int {
	a1 := a.(words)
	b1 := b.(words)
	if a1 < b1 {
		return -1
	} else if a1 > b1 {
		return 1
	} else {
		return 0
	}
}

func TestST_Put(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	st := NewST(wordsComparator)
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if st.Size() != 10 {
		t.Errorf("expect 10, got %d", st.Size())
	}
	if err := st.Put(nil, 2); err == nil {
		t.Error("should throw error: argument to Put() is nil")
	}
	if err := st.Put(words("L"), nil); err != nil {
		t.Errorf("%+v\n", err)
	}
	if err := st.Put(words("S"), 2); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if val, err := st.Get(words("S")); err != nil {
			t.Errorf("%+v\n", err)
		} else {
			if val != 2 {
				t.Errorf("expect 2, got %d", val)
			}
		}
	}
}

func TestST_Contains(t *testing.T) {
	st := NewST(wordsComparator)
	if _, err := st.Contains(nil); err == nil {
		t.Error("should throw error: argument to Contains() is nil")
	}
	if ok, err := st.Contains(words("W")); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if ok {
			t.Errorf("expect false, got %t", ok)
		}
	}
}

func TestST_Get(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	st := NewST(wordsComparator)
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if _, err := st.Get(nil); err == nil {
		t.Error("should throw error: argument to Get() is nil")
	}
	if val, _ := st.Get(words("B")); val != nil {
		t.Errorf("expect nil, got %v", val)
	}

}

func TestST_Remove(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	st := NewST(wordsComparator)
	if err := st.Remove(words("A")); err != nil {
		t.Errorf("%+v\n", err)
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if err := st.Remove(nil); err == nil {
		t.Error("should throw error: argument to Delete() is nil")
	}

	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Remove(tinyST[i])
		if ok, _ := st.Contains(tinyST[i]); ok {
			t.Errorf("expect false, got %t", ok)
		}
	}
}

func TestST_Min(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	st := NewST(wordsComparator)
	if _, err := st.Min(); err == nil {
		t.Error("should throw error: called Min() with empty symbol table")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if k, _ := st.Min(); k.CompareTo(words("A")) != 0 {
		t.Errorf("expect A, got %s", k.(words))
	}

}

func TestST_Max(t *testing.T) {
	tinyST := []words{"S", "E", "A", "R", "C", "H", "E", "X", "A", "M", "P", "L", "E"}
	st := NewST(wordsComparator)
	if _, err := st.Max(); err == nil {
		t.Error("should throw error: called Max() with empty symbol table")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if k, _ := st.Max(); k.CompareTo(words("X")) != 0 {
		t.Errorf("expect A, got %s", k.(words))
	}
}

func TestST_Floor(t *testing.T) {
	tinyST := []words{"R", "C", "H"}
	st := NewST(wordsComparator)
	if _, err := st.Floor(words("R")); err == nil {
		t.Error("should throw error: calls Floor() with empty symbol table")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if _, err := st.Floor(nil); err == nil {
		t.Error("should throw error: argument to Floor() is nil")
	}
	if k, _ := st.Floor(words("C")); k.CompareTo(words("C")) != 0 {
		t.Errorf("expect key:\"C\", got %s", k.(words))
	}
	if k, _ := st.Floor(words("B")); k != nil {
		t.Errorf("expect nil, got %s", k.(words))
	}
	if k, _ := st.Floor(words("D")); k.CompareTo(words("C")) != 0 {
		t.Errorf("expect key:\"C\", got %s", k.(words))
	}
	if k, _ := st.Floor(words("H")); k.CompareTo(words("H")) != 0 {
		t.Errorf("expect key:\"H\", got %s", k.(words))
	}
}

func TestST_Ceiling(t *testing.T) {
	tinyST := []words{"R", "C", "H"}
	st := NewST(wordsComparator)
	if _, err := st.Ceiling(words("M")); err == nil {
		t.Error("should throw error: calls Ceiling() with empty symbol table")
	}
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}
	if _, err := st.Ceiling(nil); err == nil {
		t.Error("should throw error: argument to Ceiling() is nil key")
	}
	if _, err := st.Ceiling(words("W")); err == nil {
		t.Error("should throw error: argument to Ceiling() is too large")
	}
	if k, _ := st.Ceiling(words("B")); k.CompareTo(words("C")) != 0 {
		t.Errorf("expect key:\"C\", got %s", k.(words))
	}
	if k, _ := st.Ceiling(words("R")); k.CompareTo(words("R")) != 0 {
		t.Errorf("expect key:\"R\", got %s", k.(words))
	}
}

func TestST_Keys(t *testing.T) {
	tinyST := []words{"A", "B", "C"}
	st := NewST(wordsComparator)
	for i := 0; i < len(tinyST); i++ {
		//nolint:errcheck
		st.Put(tinyST[i], i)
	}

	keys := st.Keys()
	for index, k := range keys {
		if k.(words) != tinyST[index] {
			t.Errorf("expect %s, got %s", tinyST[index], k.(words))
		}
	}
}
