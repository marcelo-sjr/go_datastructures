package linkedlist

// Node represents a single element in the linked list.
// Each node stores a value and a reference to the next node.
type Node[T comparable] struct {
	data T
	next *Node[T]
}

// Next returns the next node in the list.
// It returns nil if this is the last node.
func (node *Node[T]) Next() *Node[T] {
	return node.next
}

// Value returns the value stored in the node.
func (n *Node[T]) Value() T {
	return n.data
}