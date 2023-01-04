package datastructure

type Set[T comparable] map[T]bool

func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T])
	set.Add(values...)
	return set
}

func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = true
	}
}

// Contain 包含
func (s Set[T]) Contain(value T) bool {
	_, ok := s[value]
	return ok
}

// ContainAll 包含多个
func (s Set[T]) ContainAll(other Set[T]) bool {
	for k := range other {
		_, ok := s[k]
		if !ok {
			return false
		}
	}
	return true
}

// Clone 返回 set 的副本
func (s Set[T]) Clone() Set[T] {
	set := NewSet[T]()
	set.Add(s.Values()...)
	return set
}

// Delete 删除多个值
func (s Set[T]) Delete(values ...T) {
	for _, v := range values {
		delete(s, v)
	}
}

// Equal 检查两个集合是否具有相同的元素
func (s Set[T]) Equal(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	return s.ContainAll(other) && other.ContainAll(s)
}

// Iterate 通过集合的每个元素调用函数
func (s Set[T]) Iterate(fn func(value T)) {
	for v := range s {
		fn(v)
	}
}

// IsEmpty 检查集合是否为空
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Size 获取集合中的元素个数
func (s Set[T]) Size() int {
	return len(s)
}

// Values 返回集合的所有值
func (s Set[T]) Values() []T {
	values := make([]T, 0)

	s.Iterate(func(value T) {
		values = append(values, value)
	})

	return values
}

// Union 创建一个新集合，包含集合的所有元素和其他
func (s Set[T]) Union(other Set[T]) Set[T] {
	set := s.Clone()
	set.Add(other.Values()...)
	return set
}

// Intersection 创建一个新集合，其元素都包含在 set s 和 other
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	set := NewSet[T]()
	s.Iterate(func(value T) {
		if other.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

// SymmetricDifference 创建一个新集合，其元素在 set1 或 set2 中，但不在两个集合中
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	set := NewSet[T]()
	s.Iterate(func(value T) {
		if !other.Contain(value) {
			set.Add(value)
		}
	})

	other.Iterate(func(value T) {
		if !s.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

// Minus 创建一个集合，其元素在原始集中但不在比较集中
func (s Set[T]) Minus(comparedSet Set[T]) Set[T] {
	set := NewSet[T]()

	s.Iterate(func(value T) {
		if !comparedSet.Contain(value) {
			set.Add(value)
		}
	})

	return set
}
