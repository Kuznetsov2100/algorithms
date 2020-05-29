package linkedbag

// A generic bag, implemented using a singly linked list.
type Bag struct {
	first *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

// Initialize a Bag.
func New() *Bag {
	return &Bag{}
}

// Adds the item to this bag.
func (b *Bag) Add(element interface{}) {
	var oldfirst *node = b.first
	b.first = &node{}
	b.first.item = element
	b.first.next = oldfirst
	b.n++
}

// Is this bag empty?
func (b *Bag) IsEmpty() bool {
	return b.first == nil
}

// Returns the number of items in this bag.
func (b *Bag) Size() int {
	return b.n
}
