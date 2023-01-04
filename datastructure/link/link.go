package datastructure

// LinkNode 是一个链表节点，它的 Value 和 Pre 指向前一个节点，Next 指向链接的下一个节点。
type LinkNode[T any] struct {
	Value T
	Pre   *LinkNode[T]
	Next  *LinkNode[T]
}

func NewLinkNode[T any](value T) *LinkNode[T] {
	return &LinkNode[T]{Value: value}
}
