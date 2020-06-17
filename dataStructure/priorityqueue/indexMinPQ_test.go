package priorityqueue

import (
	"math/rand"
	"testing"

	"github.com/handane123/algorithms/sort"
)

type words string
type wordslist []words

func (l wordslist) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l wordslist) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l wordslist) Len() int {
	return len(l)
}

func (l wordslist) Shuffle() {
	rand.Shuffle(l.Len(), l.Swap)
}

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

func testPanic(t *testing.T, msg string) {
	if r := recover(); r == nil {
		t.Errorf("%s did not panic", msg)
	}
}

func TestIndexMinPQ_NewIndexMinPQ(t *testing.T) {
	defer testPanic(t, "maxN < 0")
	pq := NewIndexMinPQ(-2)
	if !pq.IsEmpty() {
		t.Error("should be empty")
	}

}

func TestIndexMinPQ_Insert(t *testing.T) {
	defer testPanic(t, "index illegal")
	strings := []words{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings))
	if _, err := pq.MinKey(); err == nil {
		t.Error("should throw error:priority queue underflow")
	}
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}
	if pq.Size() != len(strings) {
		t.Errorf("expect %d, got %d", pq.Size(), len(strings))
	}
	if k, err := pq.MinKey(); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if key := k.(words); key != "best" {
			t.Errorf("expect \"best\", got %s", key)
		}
	}
	if err := pq.Insert(9, words("love")); err == nil {
		t.Error("should throw error: index is already in the priority queue")
	}
	//nolint:errcheck
	pq.Insert(10, words("family"))
}

func TestIndexMinPQ_DelMin(t *testing.T) {
	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings))
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}
	sortedStrings := make(wordslist, len(strings))
	copy(sortedStrings, strings)
	sort.QuickSort(sortedStrings)
	for j := 0; !pq.IsEmpty(); j++ {
		if i, err := pq.DelMin(); err != nil {
			t.Errorf("%+v\n", err)
		} else {
			if strings[i] != sortedStrings[j] {
				t.Errorf("expect %s, got %s", sortedStrings[j], strings[i])
			}
		}
	}
	if _, err := pq.DelMin(); err == nil {
		t.Error("should throw error: priority queue underflow")
	}
}

func TestIndexMinPQ_DecreaseKey(t *testing.T) {
	defer testPanic(t, "index illegal")

	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings) + 1)
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}

	if err := pq.DecreaseKey(10, words("trump")); err == nil {
		t.Error("should throw error: index is not in the priority queue")
	}

	if err := pq.DecreaseKey(9, words("worst")); err == nil {
		t.Error("should throw error:calling DecreaseKey() with a key equal to the key in the priority queue")
	}
	if err := pq.DecreaseKey(9, words("zoo")); err == nil {
		t.Error("should throw error:calling DecreaseKey() with a key strictly greater than the key in the priority queue")
	}
	//nolint:errcheck
	pq.DecreaseKey(9, words("selah"))

	//nolint:errcheck
	pq.DecreaseKey(11, words("trump"))
}

func TestIndexMinPQ_IncreaseKey(t *testing.T) {
	defer testPanic(t, "index illegal")

	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings) + 1)
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}

	if err := pq.IncreaseKey(10, words("trump")); err == nil {
		t.Error("should throw error: index is not in the priority queue")
	}

	if err := pq.IncreaseKey(9, words("worst")); err == nil {
		t.Error("should throw error:calling IncreaseKey() with a key equal to the key in the priority queue")
	}
	if err := pq.IncreaseKey(9, words("telepresence")); err == nil {
		t.Error("should throw error:calling IncreaseKey() with a key strictly less than the key in the priority queue")
	}
	//nolint:errcheck
	pq.IncreaseKey(9, words("zoo"))

	//nolint:errcheck
	pq.IncreaseKey(11, words("trump"))
}

func TestIndexMinPQ_ChangeKey(t *testing.T) {
	defer testPanic(t, "index illegal")

	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings) + 1)
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}

	if err := pq.ChangeKey(10, words("spacex")); err == nil {
		t.Error("should throw error: index is not in the priority queue")
	}

	//nolint:errcheck
	pq.ChangeKey(9, words("spacex"))

	//nolint:errcheck
	pq.ChangeKey(11, words("spacex"))
}

func TestIndexMinPQ_Delete(t *testing.T) {
	defer testPanic(t, "index illegal")

	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings) + 1)
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}

	if err := pq.Delete(10); err == nil {
		t.Error("should throw error: index is not in the priority queue")
	}
	//nolint:errcheck
	pq.Delete(8)

	//nolint:errcheck
	pq.Delete(11)
}

func TestIndexMinPQ_MinIndex(t *testing.T) {

	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings) + 1)
	if _, err := pq.MinIndex(); err == nil {
		t.Error("should throw error: priority queue underflow")
	}
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}
	if index, err := pq.MinIndex(); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if index != 3 {
			t.Errorf("expect 3, got %d", index)
		}
	}

}

func TestIndexMinPQ_KeyOf(t *testing.T) {
	defer testPanic(t, "index illegal")

	strings := wordslist{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	pq := NewIndexMinPQ(len(strings) + 1)
	for i := range strings {
		if err := pq.Insert(i, strings[i]); err != nil {
			t.Errorf("%+v\n", err)
		}
	}

	if _, err := pq.KeyOf(10); err == nil {
		t.Error("should throw error: index is not in the priority queue")
	}
	if k, err := pq.KeyOf(9); err != nil {
		t.Errorf("%+v\n", err)
	} else {
		if key := k.(words); key != "worst" {
			t.Errorf("expect \"worst\", got %s", key)
		}
	}
	//nolint:errcheck
	pq.KeyOf(11)
}
