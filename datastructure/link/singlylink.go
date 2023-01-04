package datastructure

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

// SinglyLink 单向链表
type SinglyLink[T any] struct {
	Head   *LinkNode[T]
	length int
}

func NewSinglyLink[T any]() *SinglyLink[T] {
	return &SinglyLink[T]{
		Head: nil,
	}
}

// InsertAtHead 插入节点到头部
func (l *SinglyLink[T]) InsertAtHead(value T) {
	//创建新节点
	newNode := NewLinkNode(value).Next
	//新链表的下一个节点为原链表的头节点
	newNode.Next = l.Head
	//链表的头节点为新节点
	l.Head = newNode
	//链表长度加1
	l.length++
}

// InsertAtTail 插入节点到尾部
func (l *SinglyLink[T]) InsertAtTail(value T) {
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
	current.Next = newNode
	//链表长度加1
	l.length++
}

// InsertAt 插入节点到指定位置
func (l *SinglyLink[T]) InsertAt(index int, value T) error {
	size := l.length
	if index < 0 || index > size {
		return errors.New("param index should between 0 and the length of singly link.")
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
			current.Next = newNode
			l.length++

			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("singly link list no exist")
}

// DeleteAtHead 在头索引处删除单链表中的值
func (l *SinglyLink[T]) DeleteAtHead() error {
	if l.Head == nil {
		return errors.New("singly l list no exist")
	}
	current := l.Head
	l.Head = current.Next
	l.length--

	return nil
}

// DeleteAtTail 在尾索引处删除单链表中的值
func (l *SinglyLink[T]) DeleteAtTail() error {
	if l.Head == nil {
		return errors.New("singly l list no exist")
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

// DeleteAt delete value in singly linklist at index
func (l *SinglyLink[T]) DeleteAt(index int) error {
	if l.Head == nil {
		return errors.New("singly l list no exist")
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

// DeleteValue 删除单链表中的值
func (l *SinglyLink[T]) DeleteValue(value T) {
	if l.Head == nil {
		return
	}
	dummyHead := NewLinkNode(value)
	dummyHead.Next = l.Head
	current := dummyHead

	for current.Next != nil {
		if reflect.DeepEqual(current.Next.Value, value) {
			current.Next = current.Next.Next
			l.length--
		} else {
			current = current.Next
		}
	}

	l.Head = dummyHead.Next
}

// Reverse 反转链表
func (l *SinglyLink[T]) Reverse() {
	var pre, next *LinkNode[T]

	current := l.Head

	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}

	l.Head = pre
}

// GetMiddleNode 返回链表中间索引处的节点
func (l *SinglyLink[T]) GetMiddleNode() *LinkNode[T] {
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

// Size 返回单链表的计数
func (l *SinglyLink[T]) Size() int {
	return l.length
}

// Values 返回所有单链表节点值的切片
func (l *SinglyLink[T]) Values() []T {
	var result []T
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// IsEmpty 检查链接是否为空
func (l *SinglyLink[T]) IsEmpty() bool {
	return l.length == 0
}

// Clear 清空链表
func (l *SinglyLink[T]) Clear() {
	l.Head = nil
	l.length = 0
}

// Print 打印链表的所有节点信息
func (l *SinglyLink[T]) Print() {
	current := l.Head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}
