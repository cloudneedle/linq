package slice

import (
	"github.com/gocrud/linq"
	"math/rand"
)

// Distinct 函数返回一个新的切片，其中包含原始切片中的不重复元素。
func Distinct[T comparable](items []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(items))
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// Group 函数返回一个新的 map，其中包含原始切片中的元素按照键选择器函数的结果分组后的结果。
// 键选择器函数接收切片中的每个元素，返回一个键。
// 例如，如果要将切片中的元素按照奇偶性分组，可以这样调用：
//
//	result := GroupBy(items, func(item int) bool {
//	  return item%2 == 0
//	})
//
// 如果要将切片中的元素按照字符串长度分组，可以这样调用：
//
//	result := GroupBy(items, func(item string) int {
//	  return len(item)
//	})
func Group[T any, K comparable](items []T, keySelector func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range items {
		key := keySelector(item)
		result[key] = append(result[key], item)
	}
	return result
}

// Max 函数返回切片中的最大值。
// less 参数是一个函数，接收切片中的每个元素，返回一个整数。
// 如果切片为空，返回 nil。
func Max[T any](items []T, less func(T) int) *T {
	if len(items) == 0 {
		return nil
	}
	max := items[0]
	for _, item := range items {
		if less(item) > less(max) {
			max = item
		}
	}
	return &max
}

// Min 函数返回切片中的最小值。
// less 参数是一个函数，接收切片中的每个元素，返回一个整数。
// 如果切片为空，返回 nil。
func Min[T any](items []T, less func(T) int) *T {
	if len(items) == 0 {
		return nil
	}
	min := items[0]
	for _, item := range items {
		if less(item) < less(min) {
			min = item
		}
	}
	return &min
}

// Reduce 函数将泛型切片中的所有元素累加到一个值中。
// seed 参数是初始值，accumulator 参数是累加器函数。
// 累加器函数接收两个参数，第一个参数是累加器的当前值，第二个参数是切片中的当前元素。
// 累加器函数返回的值将作为下一次调用累加器函数时的第一个参数。
// 例如，如果要将切片中的所有元素相加，可以这样调用：
//
//	result := Reduce(items, 0, func(a, b int) int {
//	  return a + b
//	})
//
// 如果要将切片中的所有元素相乘，可以这样调用：
//
//	result := Reduce(items, 1, func(a, b int) int {
//	  return a * b
//	})
func Reduce[T any, U any](items []T, seed U, accumulator func(U, T) U) U {
	result := seed
	for _, item := range items {
		result = accumulator(result, item)
	}
	return result
}

// Reverse 函数将切片中的元素顺序反转。
func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i, j := 0, len(items)-1; i < len(items); i, j = i+1, j-1 {
		result[i] = items[j]
	}
	return result
}

// Map 用于将切片中的每个元素映射成一个新的元素，返回一个包含新元素的切片
// selector 参数是映射函数，接收切片中的每个元素，返回一个新元素。
// 例如，如果要将切片中的每个元素转换成字符串，可以这样调用：
//
//	result := Map(items, func(item int) string {
//	  return strconv.Itoa(item)
//	})
//
// 如果要将切片中的每个元素转换成另一个类型，可以这样调用：
//
//	result := Map(items, func(item int) MyType {
//	  return MyType(item)
//	})
func Map[T, R any](items []T, selector func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = selector(item)
	}
	return result
}

// Single 函数返回切片中满足条件的唯一元素。
// 如果切片中没有元素满足条件，返回 ErrorNoItemsFound。
// 如果切片中有多个元素满足条件，返回 ErrorMoreThanOneItemFound。
func Single[T any](items []T, predicate func(T) bool) (*T, error) {
	var result *T
	for _, item := range items {
		if predicate(item) {
			if result != nil {
				return nil, linq.ErrorMoreThanOneItemFound
			}
			result = new(T)
			*result = item
		}
	}
	if result == nil {
		return nil, linq.ErrorNoItemsFound
	}
	return result, nil
}

// Skip 函数返回切片中从第 n 个元素开始的所有元素。
func Skip[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return nil
	}
	return slice[n:]
}

// Sort 函数使用快速排序算法对切片进行排序。
func Sort[T any](items []T, less func(T, T) bool) []T {
	result := make([]T, len(items))
	copy(result, items)
	quickSort[T](result, less)
	return result
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

// Sum 函数用于计算切片中所有元素的总和。
// selector 参数是选择器函数，接收切片中的每个元素，返回一个整数。
// 例如，如果要计算切片中所有元素的和，可以这样调用：
//
//	result := Sum(items, func(item int) int {
//	  return item
//	})
//
// 如果要计算切片中所有元素的平方和，可以这样调用：
//
//	result := Sum(items, func(item int) int {
//	  return item * item
//	})
func Sum[T any](slice []T, selector func(T) int) int {
	sum := 0
	for _, item := range slice {
		sum += selector(item)
	}
	return sum
}

// Take 函数返回切片中的前 n 个元素。
func Take[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice
	}
	return slice[:n]
}

// Where 函数用于根据指定的条件筛选切片中的元素，返回一个新切片。
// predicate 参数是谓词函数，接收切片中的每个元素，返回一个布尔值。
// 例如，如果要筛选出切片中的偶数，可以这样调用：
//
//	result := Where(items, func(item int) bool {
//	  return item % 2 == 0
//	})
//
// 如果要筛选出切片中的奇数，可以这样调用：
//
//	result := Where(items, func(item int) bool {
//	  return item % 2 != 0
//	})
func Where[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
