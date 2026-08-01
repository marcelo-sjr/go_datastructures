package linkedlist

import "errors"

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Append(val ...T) {
	if len(val) == 0 {
		return
	}

	ll.size += len(val)

	if ll.head == nil {
		node := Node[T]{
			data: val[0],
			next: nil,
		}

		ll.head = &node
		ll.tail = &node
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

func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

func (ll *LinkedList[T]) Tail() *Node[T] {
	return ll.tail
}

func (ll *LinkedList[T]) Len() int {
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}

func (ll *LinkedList[T]) Present(val T) bool {
	n := ll.Find(val)
	if n == nil {
		return false
	}

	return true
}