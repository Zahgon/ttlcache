package ttlcache

import (
	"container/list"
)

// expirationQueue stores items that are ordered by their expiration
// timestamps. The 0th item is closest to its expiration.
type expirationQueue[K comparable, V any] []*list.Element

// newExpirationQueue creates and initializes a new expiration queue.
func newExpirationQueue[K comparable, V any]() expirationQueue[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// isEmpty checks if the queue is empty.
func (q expirationQueue[K, V]) isEmpty() bool { _ = "STUB: not implemented"; return false }

// update updates an existing item's value and position in the queue.
func (q *expirationQueue[K, V]) update(elem *list.Element) { _ = "STUB: not implemented"; return }

// push pushes a new item into the queue and updates the order of its
// elements.
func (q *expirationQueue[K, V]) push(elem *list.Element) {
	_ = "STUB: not implemented"

	// remove removes an item from the queue and updates the order of its
	// elements.
	return
}

func (q *expirationQueue[K, V]) remove(elem *list.Element) { _ = "STUB: not implemented"; return }

// Len returns the total number of items in the queue.
func (q expirationQueue[K, V]) Len() int {
	_ = "STUB: not implemented"

	// Less checks if the item at the i position expires sooner than
	// the one at the j position.
	return 0
}

func (q expirationQueue[K, V]) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap switches the places of two queue items.
func (q expirationQueue[K, V]) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Push appends a new item to the item slice.
func (q *expirationQueue[K, V]) Push(x interface{}) { _ = "STUB: not implemented"; return }

// Pop removes and returns the last item.
func (q *expirationQueue[K, V]) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// avoid memory leak
