package linq

import "math/rand"

// Linq[T] 是一个泛型接口，定义了一组用于查询和操作元素的方法。
type Linq[T any] interface {
	// Count 方法返回满足指定条件的元素的数量。
	Count(predicate func(T) bool) int
	// First 方法返回满足指定条件的第一个元素，如果找不到元素，则返回错误:ErrMoreThanOneItemFound。
	First(predicate func(T) bool) (*T, error)
	// FirstOrDefault 方法返回满足指定条件的第一个元素，如果找不到元素，则返回 nil。
	FirstOrDefault(predicate func(T) bool) *T
	// ForEach 方法对 Linq[T] 中的每个元素执行指定的操作。
	ForEach(action func(*T)) Linq[T]
	// Max 方法返回 Linq[T] 中的最大值。
	Max(less func(T) int) *T
	// Min 方法返回 Linq[T] 中的最小值。
	Min(less func(T) int) *T
	// Sort 方法返回一个新的 Linq[T]，排序元素。
	Sort(less func(T, T) bool) Linq[T]
	// Single 方法返回满足指定条件的唯一元素，如果找不到元素,
	// 如果找到多个元素，则返回错误:ErrMoreThanOneItemFound。
	// 如果找不到元素，则返回错误:ErrNoItemsFound。
	Single(predicate func(T) bool) (*T, error)
	// SingleOrDefault 方法返回满足指定条件的唯一元素，如果找不到元素，则返回 nil。
	// 如果找到多个元素，则返回错误:ErrMoreThanOneItemFound。
	SingleOrDefault(predicate func(T) bool) (*T, error)
	// Skip 方法返回一个新的 Linq[T]，其中包含从指定位置开始的所有元素。
	Skip(n int) Linq[T]
	// Sum 方法计算 Linq[T] 中所有元素的总和。
	Sum(selector func(T) int) int
	// Take 方法返回一个新的 Linq[T]，其中包含前 n 个元素。
	Take(n int) Linq[T]
	// ToSlice 方法返回一个包含 Linq[T] 中所有元素的切片。
	ToSlice() []T
	// Where 方法返回一个新的 Linq[T]，其中包含满足指定条件的元素。
	Where(predicate func(T) bool) Linq[T]
}

// slice[T] 是一个泛型类型，实现了 Linq[T] 接口。
type slice[T any] struct {
	items []T
}

// Slice 函数返回一个新的 Linq[T]，其中包含指定的元素。
func Slice[T any](items []T) Linq[T] {
	return &slice[T]{
		items: items,
	}
}

// Max 方法返回 Linq[T] 中的最大值。
func (s *slice[T]) Max(less func(T) int) *T {
	if len(s.items) == 0 {
		return nil
	}
	max := s.items[0]
	for _, item := range s.items {
		if less(item) > less(max) {
			max = item
		}
	}
	return &max
}

// Min 方法返回 Linq[T] 中的最小值。
func (s *slice[T]) Min(less func(T) int) *T {
	if len(s.items) == 0 {
		return nil
	}
	min := s.items[0]
	for _, item := range s.items {
		if less(item) < less(min) {
			min = item
		}
	}
	return &min
}

// Sum 方法计算 Linq[T] 中所有元素的总和。
func (s *slice[T]) Sum(selector func(T) int) int {
	sum := 0
	for _, item := range s.items {
		sum += selector(item)
	}
	return sum
}

// OrderByAsc 方法返回一个新的 Linq[T]，其中包含按升序排序的元素。
func (s *slice[T]) Sort(less func(T, T) bool) Linq[T] {
	result := make([]T, len(s.items))
	copy(result, s.items)
	quickSort[T](result, less)
	return &slice[T]{items: result}
}

// quickSort 方法使用快速排序算法对切片进行排序。
func quickSort[T any](items []T, less func(T, T) bool) {
	if len(items) < 2 {
		return
	}
	left, right := 0, len(items)-1
	pivot := rand.Int() % len(items)
	items[pivot], items[right] = items[right], items[pivot]
	for i := range items {
		if less(items[i], items[right]) {
			items[left], items[i] = items[i], items[left]
			left++
		}
	}
	items[left], items[right] = items[right], items[left]
	quickSort[T](items[:left], less)
	quickSort[T](items[left+1:], less)
}

// ForEach 方法对 Linq[T] 中的每个元素执行指定的操作。
func (s *slice[T]) ForEach(action func(*T)) Linq[T] {
	for i := range s.items {
		action(&s.items[i])
	}
	return s
}

// Where 方法返回一个新的 Linq[T]，其中包含满足指定条件的元素。
func (s *slice[T]) Where(predicate func(T) bool) Linq[T] {
	var result []T
	for _, item := range s.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return &slice[T]{items: result}
}

// ToSlice 方法返回一个包含 Linq[T] 中所有元素的切片。
func (s *slice[T]) ToSlice() []T {
	return s.items
}

// Count 方法返回满足指定条件的元素的数量。
func (s *slice[T]) Count(predicate func(T) bool) int {
	count := 0
	for _, item := range s.items {
		if predicate(item) {
			count++
		}
	}
	return count
}

// Single 方法返回满足指定条件的唯一元素，
// 如果找到多个元素，则返回错误:ErrMoreThanOneItemFound。
// 如果找不到元素，则返回错误:ErrNoItemsFound。
func (s *slice[T]) Single(predicate func(T) bool) (*T, error) {
	var result *T
	for _, item := range s.items {
		if predicate(item) {
			if result != nil {
				return nil, ErrorMoreThanOneItemFound
			}
			result = &item
		}
	}
	if result == nil {
		return nil, ErrorNoItemsFound
	}
	return result, nil
}

// SingleOrDefault 方法返回满足指定条件的唯一元素，如果找不到元素，则返回 nil。
// 如果找到多个元素，则返回错误:ErrMoreThanOneItemFound。
func (s *slice[T]) SingleOrDefault(predicate func(T) bool) (*T, error) {
	result, err := s.Single(predicate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// First 方法返回满足指定条件的第一个元素，如果找不到元素，则返回错误:ErrMoreThanOneItemFound。
func (s *slice[T]) First(predicate func(T) bool) (*T, error) {
	for _, item := range s.items {
		if predicate(item) {
			return &item, nil
		}
	}
	return nil, ErrorMoreThanOneItemFound
}

// FirstOrDefault 方法返回满足指定条件的第一个元素，如果找不到元素，则返回 nil。
func (s *slice[T]) FirstOrDefault(predicate func(T) bool) *T {
	result, err := s.First(predicate)
	if err != nil {
		return nil
	}
	return result
}

// Skip 方法返回一个新的 Linq[T]，其中包含从指定位置开始的所有元素。
func (s *slice[T]) Skip(n int) Linq[T] {
	if n >= len(s.items) {
		return &slice[T]{items: []T{}}
	}
	return &slice[T]{items: s.items[n:]}
}

// Take 方法返回一个新的 Linq[T]，其中包含前 n 个元素。
func (s *slice[T]) Take(n int) Linq[T] {
	if n >= len(s.items) {
		return s
	}
	return &slice[T]{items: s.items[:n]}
}
