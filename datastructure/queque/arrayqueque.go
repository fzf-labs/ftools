package queque

import (
	"fmt"
	"reflect"
)

// ArrayQueue 用切片实现队列
type ArrayQueue[T any] struct {
	items    []T // 切片
	head     int // 入队位置
	tail     int // 出队位置
	capacity int // 容量
	size     int //	元素个数
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		items:    make([]T, 0, capacity),
		head:     0,
		tail:     0,
		capacity: capacity,
		size:     0,
	}
}

func (q *ArrayQueue[T]) Data() []T {
	var items []T
	for i := q.head; i < q.tail; i++ {
		items = append(items, q.items[i])
	}
	return items
}

func (q *ArrayQueue[T]) Size() int {
	return q.size
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *ArrayQueue[T]) Front() T {
	return q.items[q.head]
}

func (q *ArrayQueue[T]) Back() T {
	return q.items[q.tail-1]
}

// EnQueue put element into queue
func (q *ArrayQueue[T]) Enqueue(item T) bool {
	if q.head == 0 && q.tail == q.capacity {
		return false
	} else if q.head != 0 && q.tail == q.capacity {
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}
		q.tail = q.tail - q.head
		q.head = 0
	}

	q.items = append(q.items, item)
	q.tail++
	q.size++
	return true
}

// DeQueue 删除队列的头元素并返回它，如果队列为空，则返回 nil 和错误
func (q *ArrayQueue[T]) Dequeue() (T, bool) {
	var item T
	if q.head == q.tail {
		return item, false
	}
	item = q.items[q.head]
	q.head++
	q.size--
	return item, true
}

// Clear the queue data
func (q *ArrayQueue[T]) Clear() {
	capacity := q.capacity
	q.items = make([]T, 0, capacity)
	q.head = 0
	q.tail = 0
	q.size = 0
	q.capacity = capacity
}

// Contain checks if the value is in queue or not
func (q *ArrayQueue[T]) Contain(value T) bool {
	for _, v := range q.items {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Print queue data
func (q *ArrayQueue[T]) Print() {
	info := "["
	for i := q.head; i < q.tail; i++ {
		info += fmt.Sprintf("%+v, ", q.items[i])
	}
	info += "]"
	fmt.Println(info)
}
