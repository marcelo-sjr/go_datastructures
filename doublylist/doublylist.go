package dlist

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

