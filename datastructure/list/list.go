package datastructure

import "reflect"

// List 线性表
type List[T any] struct {
	data []T
}

func NewList[T any](data []T) *List[T] {
	return &List[T]{
		data: data,
	}
}

// Data 数据
func (l *List[T]) Data() []T {
	return l.data
}

// ValueOf 返回指定索引的值
func (l *List[T]) ValueOf(index int) (*T, bool) {
	if index < 0 || index >= len(l.data) {
		return nil, false
	}
	return &l.data[index], true
}

// IndexOf 通过值来找索引
func (l *List[T]) IndexOf(value T) int {
	index := -1
	data := l.data
	for i, v := range data {
		if reflect.DeepEqual(v, value) {
			index = i
			break
		}
	}
	return index
}

// LastIndexOf 返回此列表中最后一次出现的值的索引。
// if not found return -1
func (l *List[T]) LastIndexOf(value T) int {
	index := -1
	data := l.data
	for i := len(data) - 1; i >= 0; i-- {
		if reflect.DeepEqual(data[i], value) {
			index = i
			break
		}
	}
	return index
}

// IndexOfFunc 返回满足 f(v) 的第一个索引
// if not found return -1
func (l *List[T]) IndexOfFunc(f func(T) bool) int {
	index := -1
	data := l.data
	for i, v := range data {
		if f(v) {
			index = i
			break
		}
	}
	return index
}

// LastIndexOfFunc 返回此列表中满足 f() 的值的最后一次出现的索引
// if not found return -1
func (l *List[T]) LastIndexOfFunc(f func(T) bool) int {
	index := -1
	data := l.data
	for i := len(data) - 1; i >= 0; i-- {
		if f(data[i]) {
			index = i
			break
		}
	}
	return index
}

// Contain 检查列表中的值是否
func (l *List[T]) Contain(value T) bool {
	data := l.data
	for _, v := range data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Push 将值附加到列表数据
func (l *List[T]) Push(value T) {
	l.data = append(l.data, value)
}

// InsertAtFirst 在第一个索引处将值插入列表
func (l *List[T]) InsertAtFirst(value T) {
	l.InsertAt(0, value)
}

// InsertAtLast 在最后一个索引处将值插入列表
func (l *List[T]) InsertAtLast(value T) {
	l.InsertAt(len(l.data), value)
}

// InsertAt 将值插入索引处的列表
func (l *List[T]) InsertAt(index int, value T) {
	data := l.data
	size := len(data)

	if index < 0 || index > size {
		return
	}
	l.data = append(data[:index], append([]T{value}, data[index:]...)...)
}

// PopFirst 删除列表的第一个值并返回
func (l *List[T]) PopFirst() (*T, bool) {
	if len(l.data) == 0 {
		return nil, false
	}

	v := l.data[0]
	l.DeleteAt(0)

	return &v, true
}

// PopLast 删除列表的最后一个值并返回
func (l *List[T]) PopLast() (*T, bool) {
	size := len(l.data)
	if size == 0 {
		return nil, false
	}

	v := l.data[size-1]
	l.DeleteAt(size - 1)

	return &v, true
}

// DeleteAt 删除索引处列表的值
func (l *List[T]) DeleteAt(index int) {
	data := l.data
	size := len(data)
	if index < 0 || index > size-1 {
		return
	}
	if index != size-1 {
		data = append(data[:index], data[index+1:]...)
	}
	l.data = data
}

// DeleteIf 删除所有满足 f()，返回已删除元素的计数
func (l *List[T]) DeleteIf(f func(T) bool) int {
	data := l.data
	size := len(data)

	var c int
	for index := 0; index < len(data); index++ {
		if !f(data[index]) {
			continue
		}
		if index != size-1 {
			data = append(data[:index], data[index+1:]...)
			index--
		}
		c++
	}

	if c > 0 {
		l.data = data
	}
	return c
}

// UpdateAt更新索引处列表的值，索引应该在 0 和列表大小 -1 之间
func (l *List[T]) UpdateAt(index int, value T) {
	data := l.data
	size := len(data)

	if index < 0 || index >= size {
		return
	}
	l.data = append(data[:index], append([]T{value}, data[index+1:]...)...)
}

// Equal 将列表与其他列表进行比较，使用 reflect.DeepEqual
func (l *List[T]) Equal(other *List[T]) bool {
	if len(l.data) != len(other.data) {
		return false
	}

	for i := 0; i < len(l.data); i++ {
		if !reflect.DeepEqual(l.data[i], other.data[i]) {
			return false
		}
	}

	return true
}

// IsEmpty 检查列表是否为空
func (l *List[T]) IsEmpty() bool {
	return len(l.data) == 0
}

// Clear 列表的数据
func (l *List[T]) Clear() {
	l.data = make([]T, 0)
}

// Clone 返回列表的副本
func (l *List[T]) Clone() *List[T] {
	cl := NewList(make([]T, len(l.data)))
	copy(cl.data, l.data)

	return cl
}

// Merge 合并两个列表，返回新列表，不改变原始列表
func (l *List[T]) Merge(other *List[T]) *List[T] {
	l1, l2 := len(l.data), len(other.data)
	ml := NewList(make([]T, l1+l2))

	data := append([]T{}, append(l.data, other.data...)...)
	ml.data = data

	return ml
}

// Size 返回列表数据项的数量
func (l *List[T]) Size() int {
	return len(l.data)
}

// Cap 返回内部数据的上限
func (l *List[T]) Cap() int {
	return cap(l.data)
}

// Swap 列表中索引 i 和 j 的值
func (l *List[T]) Swap(i, j int) {
	size := len(l.data)
	if i < 0 || i >= size || j < 0 || j >= size {
		return
	}
	l.data[i], l.data[j] = l.data[j], l.data[i]
}

// Reverse 反转列表的项目顺序
func (l *List[T]) Reverse() {
	for i, j := 0, len(l.data)-1; i < j; i, j = i+1, j-1 {
		l.data[i], l.data[j] = l.data[j], l.data[i]
	}
}

// Unique 删除列表中的重复项
func (l *List[T]) Unique() {
	data := l.data
	size := len(data)

	uniqueData := make([]T, 0)
	for i := 0; i < size; i++ {
		value := data[i]
		skip := true
		for _, v := range uniqueData {
			if reflect.DeepEqual(value, v) {
				skip = false
				break
			}
		}
		if skip {
			uniqueData = append(uniqueData, value)
		}
	}

	l.data = uniqueData
}

// Union 创建一个新列表，包含列表 l 中的所有元素和其他元素，删除重复元素。
func (l *List[T]) Union(other *List[T]) *List[T] {
	result := NewList([]T{})

	result.data = append(result.data, l.data...)
	result.data = append(result.data, other.data...)
	result.Unique()

	return result
}

// Intersection 创建一个新列表，其元素都包含在列表 l 和其他
func (l *List[T]) Intersection(other *List[T]) *List[T] {
	result := NewList(make([]T, 0))

	for _, v := range l.data {
		if other.Contain(v) {
			result.data = append(result.data, v)
		}
	}

	return result
}

// SubList 返回指定 fromIndex（包括）和 toIndex（不包括）之间的原始列表的子列表。
func (l *List[T]) SubList(fromIndex, toIndex int) *List[T] {
	data := l.data[fromIndex:toIndex]
	subList := make([]T, len(data))
	copy(subList, data)
	return NewList(subList)
}
