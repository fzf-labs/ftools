package datastructure

import (
	"fmt"
	"github.com/pkg/errors"
)

// DoublyLink  双向链表 Pre 指针指向链接的前一个节点，Next 指针指向链接的下一个节点。
type DoublyLink[T any] struct {
	Head   *LinkNode[T]
	length int
}

func NewDoublyLink[T any]() *DoublyLink[T] {
	return &DoublyLink[T]{Head: nil}
}

// InsertAtHead 在头索引处将值插入双向链表
func (l *DoublyLink[T]) InsertAtHead(value T) {
	newNode := NewLinkNode(value)
	size := l.Size()

	if size == 0 {
		l.Head = newNode
		l.length++
		return
	}

	newNode.Next = l.Head
	newNode.Pre = nil

	l.Head.Pre = newNode
	l.Head = newNode

	l.length++
}

// InsertAtTail insert value into doubly linklist at tail index
func (l *DoublyLink[T]) InsertAtTail(value T) {
	current := l.Head
	if current == nil {
		l.InsertAtHead(value)
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	newNode := NewLinkNode(value)
	newNode.Next = nil
	newNode.Pre = current
	current.Next = newNode

	l.length++
}

// InsertAt insert value into doubly linklist at index
func (l *DoublyLink[T]) InsertAt(index int, value T) error {
	size := l.length
	if index < 0 || index > size {
		return errors.New("param index should between 0 and the length of doubly l.")
	}

	if index == 0 {
		l.InsertAtHead(value)
		return nil
	}

	if index == size {
		l.InsertAtTail(value)
		return nil
	}

	i := 0
	current := l.Head

	for current != nil {
		if i == index-1 {
			newNode := NewLinkNode(value)
			newNode.Next = current.Next
			newNode.Pre = current

			current.Next = newNode
			l.length++

			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("doubly l list no exist")
}

// DeleteAtHead delete value in doubly linklist at head index
func (l *DoublyLink[T]) DeleteAtHead() error {
	if l.Head == nil {
		return errors.New("doubly l list no exist")
	}
	current := l.Head
	l.Head = current.Next
	l.Head.Pre = nil
	l.length--

	return nil
}

// DeleteAtTail delete value in doubly linklist at tail index
func (l *DoublyLink[T]) DeleteAtTail() error {
	if l.Head == nil {
		return errors.New("doubly l list no exist")
	}
	current := l.Head
	if current.Next == nil {
		return l.DeleteAtHead()
	}

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	l.length--
	return nil
}

// DeleteAt delete value in doubly linklist at index
func (l *DoublyLink[T]) DeleteAt(index int) error {
	if l.Head == nil {
		return errors.New("doubly l list no exist")
	}
	current := l.Head
	if current.Next == nil || index == 0 {
		return l.DeleteAtHead()
	}

	if index == l.length-1 {
		return l.DeleteAtTail()
	}

	if index < 0 || index > l.length-1 {
		return errors.New("param index should between 0 and l size -1.")
	}

	i := 0
	for current != nil {
		if i == index-1 {
			current.Next = current.Next.Next
			l.length--
			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("delete error")
}

// Reverse the linked list
func (l *DoublyLink[T]) Reverse() {
	current := l.Head
	var temp *LinkNode[T]

	for current != nil {
		temp = current.Pre
		current.Pre = current.Next
		current.Next = temp
		current = current.Pre
	}

	if temp != nil {
		l.Head = temp.Pre
	}
}

// GetMiddleNode return node at middle index of linked list
func (l *DoublyLink[T]) GetMiddleNode() *LinkNode[T] {
	if l.Head == nil {
		return nil
	}
	if l.Head.Next == nil {
		return l.Head
	}
	fast := l.Head
	slow := l.Head

	for fast != nil {
		fast = fast.Next

		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		} else {
			return slow
		}
	}
	return slow
}

// Size return the count of doubly linked list
func (l *DoublyLink[T]) Size() int {
	return l.length
}

// Values return slice of all doubly linklist node value
func (l *DoublyLink[T]) Values() []T {
	var result []T
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Print all nodes info of a linked list
func (l *DoublyLink[T]) Print() {
	current := l.Head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

// IsEmpty checks if link is empty or not
func (l *DoublyLink[T]) IsEmpty() bool {
	return l.length == 0
}

// Clear all nodes in doubly linklist
func (l *DoublyLink[T]) Clear() {
	l.Head = nil
	l.length = 0
}
