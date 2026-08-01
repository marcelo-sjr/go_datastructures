package linkedlist

type Node[T comparable] struct {
	data T
	next *Node[T]
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

func (n *Node[T]) Value() T {
	return n.data
}