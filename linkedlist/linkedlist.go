package linkedlist

import "errors"

// LinkedList represents a generic singly linked list.
// It stores references to the first and last nodes and keeps
// track of the current number of elements.
type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New creates and returns an empty linked list with the type inserted.
func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append appends one or more values to the end of the list.
// If no values are provided, the list remains unchanged.
func (ll *LinkedList[T]) Append(val ...T) {
	if len(val) == 0 {
		return
	}

	ll.size += len(val)

	if ll.head == nil {
		node := &Node[T]{
			data: val[0],
			next: nil,
		}

		ll.head = node
		ll.tail = node
		val = val[1:]
	}

	for _, v := range val {
		node := &Node[T]{
			data: v,
			next: nil,
		}

		ll.tail.next = node
		ll.tail = node
	}
}

// Values returns all elements in the list as a slice,
// preserving their insertion order.
func (ll *LinkedList[T]) Values() []T {
	values := make([]T, 0, ll.size)

	for curr := ll.head; curr != nil; curr = curr.next {
		values = append(values, curr.data)
	}

	return values
}

// Find searches for the first node containing val.
// It returns nil if the value is not present in the list.
func (ll *LinkedList[T]) Find(val T) *Node[T] {
	if ll == nil || ll.head == nil {
		return nil
	}

	for curr := ll.head; curr != nil; curr = curr.next {
		if curr.data == val {
			return curr
		}
	}

	return nil
}

// InsertAfter inserts a new node immediately after the target node.
// It returns an error if the target node is nil.
func (ll *LinkedList[T]) InsertAfter(target *Node[T], val T) error {
	if target == nil {
		return errors.New("target node cannot be nil")
	}

	target.next = &Node[T]{
		data: val,
		next: target.next,
	}

	if target == ll.tail {
		ll.tail = target.next
	}

	ll.size++

	return nil
}

// Delete removes the specified node from the list.
// It returns true if the node was found and removed,
// or false otherwise.
func (ll *LinkedList[T]) Delete(target *Node[T]) bool {
	if ll == nil || ll.head == nil || target == nil {
		return false
	}

	if target == ll.head {
		ll.head = target.next

		if ll.head == nil {
			ll.tail = nil
		}

		ll.size--
		return true
	}

	prev := ll.head
	for curr := ll.head.next; curr != nil; curr = curr.next {
		if target == curr {
			prev.next = curr.next

			if curr == ll.tail {
				ll.tail = prev
			}

			ll.size--
			return true
		}

		prev = curr
	}

	return false
}

// PushFront inserts a new value at the beginning of the list.
func (ll *LinkedList[T]) PushFront(val T) {
	node := &Node[T]{
		data: val,
		next: ll.head,
	}

	ll.head = node
	if ll.tail == nil {
		ll.tail = node
	}

	ll.size++
}

// PushBack inserts a new value at the end of the list.
func (ll *LinkedList[T]) PushBack(val T) {
	node := &Node[T]{
		data: val,
		next: nil,
	}

	if ll.tail == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	ll.size++
}

// Head returns the first node of the list.
// It returns nil if the list is empty.
func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

// Tail returns the last node of the list.
// It returns nil if the list is empty.
func (ll *LinkedList[T]) Tail() *Node[T] {
	return ll.tail
}

// Len returns the number of elements currently stored in the list.
func (ll *LinkedList[T]) Len() int {
	return ll.size
}

// IsEmpty reports whether the list contains no elements.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}

// Clear removes all elements from the list,
// leaving it in an empty state.
func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

// Present reports whether the specified value exists in the list.
func (ll *LinkedList[T]) Present(val T) bool {
	n := ll.Find(val)
	if n == nil {
		return false
	}

	return true
}