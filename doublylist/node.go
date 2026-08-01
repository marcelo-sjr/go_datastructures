package doublylist

type Node[T comparable] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}

func (n *Node[T]) Value() T {
	return n.data
}