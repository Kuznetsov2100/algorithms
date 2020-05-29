package linkedbag

type Bag struct {
	first *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

func New() *Bag {
	return &Bag{}
}

func (b *Bag) Add(element interface{}) {
	var oldfirst *node = b.first
	b.first = &node{}
	b.first.item = element
	b.first.next = oldfirst
	b.n++
}

func (b *Bag) IsEmpty() bool {
	return b.first == nil
}

func (b *Bag) Size() int {
	return b.n
}
