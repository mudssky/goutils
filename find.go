package goutils

import (
	"fmt"
	"math/rand"

	c "github.com/mudssky/goutils/constraints"
)

// IndexOf returns the index at which the first occurrence of a value is found in an array or return -1
// if the value cannot be found.
//
// 类似js的indexOf
// 示例:
// IndexOf([]int{1, 2, 3}, 2) // 返回: 1
// IndexOf([]string{"a", "b", "c"}, "d") // 返回: -1
func IndexOf[T comparable](collection []T, element T) int {
	for i, item := range collection {
		if item == element {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index at which the last occurrence of a value is found in an array or return -1
// if the value cannot be found.
//
// 类似js的同名函数
// 示例:
// LastIndexOf([]int{1, 2, 2, 3}, 2) // 返回: 2
// LastIndexOf([]string{"a", "b", "a"}, "a") // 返回: 2
func LastIndexOf[T comparable](collection []T, element T) int {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if collection[i] == element {
			return i
		}
	}

	return -1
}

// Find search an element in a slice based on a predicate. It returns element and true if element was found.
//
// 因为元素为空的情况,此时predicate判断为空,此时返回值都是空,没法判断有没有找到,所以加了一个bool的返回值区分
// 示例:
// Find([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // 返回: 2, true
// Find([]string{}, func(v string) bool { return len(v) > 0 }) // 返回: "", false
func Find[T any](collection []T, predicate func(item T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

// FindIndexOf searches an element in a slice based on a predicate and returns the index and true.
// It returns -1 and false if the element is not found.
// 示例:
// FindIndexOf([]int{1, 2, 3}, func(v int) bool { return v == 2 }) // 返回: 2, 1, true
// FindIndexOf([]string{"a", "b"}, func(v string) bool { return v == "c" }) // 返回: "", -1, false
func FindIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	for i, item := range collection {
		if predicate(item) {
			return item, i, true
		}
	}

	var result T
	return result, -1, false
}

// FindLastIndexOf searches last element in a slice based on a predicate and returns the index and true.
// It returns -1 and false if the element is not found.
// 示例:
// FindLastIndexOf([]int{1, 2, 2, 3}, func(v int) bool { return v == 2 }) // 返回: 2, 2, true
// FindLastIndexOf([]string{"a", "b"}, func(v string) bool { return v == "c" }) // 返回: "", -1, false
func FindLastIndexOf[T any](collection []T, predicate func(item T) bool) (T, int, bool) {
	length := len(collection)

	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return collection[i], i, true
		}
	}

	var result T
	return result, -1, false
}

// Nth returns the element at index `nth` of collection. If `nth` is negative, the nth element
// from the end is returned. An error is returned when nth is out of slice bounds.
//
// 类似js的at
// 示例:
// Nth([]int{1, 2, 3}, 1) // 返回: 2, nil
// Nth([]string{"a", "b", "c"}, -1) // 返回: "c", nil
// Nth([]int{}, 0) // 返回: 0, error
func Nth[T any, N c.Integer](collection []T, nth N) (T, error) {
	n := int(nth)
	l := len(collection)
	if n >= l || -n > l {
		var t T
		return t, fmt.Errorf("nth: %d out of slice bounds", n)
	}

	if n >= 0 {
		return collection[n], nil
	}
	return collection[l+n], nil
}

// MaxOfCollection returns the maximum value in a slice of ordered elements.
// It compares elements using the '>' operator and returns the largest element found.
//
// If the collection is empty, it returns the zero value of type T.
//
// 返回一个切片中的最大值
// 示例:
// MaxOfCollection([]int{1, 3, 2}) // 返回: 3
// MaxOfCollection([]string{"a", "c", "b"}) // 返回: "c"
// MaxOfCollection([]int{}) // 返回: 0
func MaxOfCollection[T c.Ordered](collection []T) T {
	var max T

	if len(collection) == 0 {
		return max
	}

	max = collection[0]

	for i := 1; i < len(collection); i++ {
		item := collection[i]

		if item > max {
			max = item
		}
	}

	return max
}

// Without returns a new slice containing all elements from the original collection
// except those specified in the exclude parameter.
//
// The function creates a new slice and does not modify the original collection.
//
// 排除exclude中的值
// 示例:
// Without([]int{1, 2, 3}, 2) // 返回: [1, 3]
// Without([]string{"a", "b", "c"}, "a", "c") // 返回: ["b"]
func Without[T comparable](collection []T, exclude ...T) []T {

	result := make([]T, 0, len(collection))
	for _, e := range collection {
		if !Contains(exclude, e) {
			result = append(result, e)
		}
	}
	return result
}

// WithoutEmpty returns a new slice containing all non-zero elements from the original collection.
// It filters out any elements that are equal to the zero value of type T.
//
// The function creates a new slice and does not modify the original collection.
//
// 排除empty值
// 示例:
// WithoutEmpty([]int{0, 1, 2, 0}) // 返回: [1, 2]
// WithoutEmpty([]string{"", "a", "", "b"}) // 返回: ["a", "b"]
func WithoutEmpty[T comparable](collection []T) []T {
	var empty T

	result := make([]T, 0, len(collection))
	for _, e := range collection {
		if e != empty {
			result = append(result, e)
		}
	}

	return result
}

// Sample returns a random item from collection.
// It uses the provided random number generator to select a random element.
//
// If the collection is empty, it returns the zero value of type T and an error.
//
// 从列表中随机取一个值
// 示例:
// Sample([]int{1, 2, 3}) // 可能返回: 2
// Sample([]string{"a", "b", "c"}) // 可能返回: "b"
func Sample[T any](collection []T) T {
	size := len(collection)
	if size == 0 {
		return Empty[T]()
	}

	return collection[rand.Intn(size)]
}

// Samples returns N random unique items from collection.
// It uses the provided random number generator to select random elements without replacement.
// If count is greater than the collection size, it returns all elements in random order.
//
// 从列表中随机取n个值
// 示例:
// Samples([]int{1, 2, 3, 4}, 2) // 可能返回: [3, 1]
// Samples([]string{"a", "b", "c"}, 5) // 可能返回: ["c", "a", "b"]
func Samples[T any](collection []T, count int) []T {
	size := len(collection)

	copy := append([]T{}, collection...)

	// 预分配结果切片容量为count和size中的较小值（结果最大可能大小）
	resultSize := count
	if size < count {
		resultSize = size
	}
	results := make([]T, 0, resultSize)

	for i := 0; i < size && i < count; i++ {
		copyLength := size - i

		index := rand.Intn(size - i)
		results = append(results, copy[index])

		// Removes element.
		// It is faster to swap with last element and remove it.
		copy[index] = copy[copyLength-1]
		copy = copy[:copyLength-1]
	}

	return results
}

// HasDuplicates checks if a slice contains any duplicate elements.
// It returns true if any element appears more than once in the collection, false otherwise.
//
// The function uses a map to track seen elements, making it an O(n) operation.
//
// 检查一个列表中是否有重复项
// 示例:
// HasDuplicates([]int{1, 2, 2}) // 返回: true
// HasDuplicates([]string{"a", "b", "c"}) // 返回: false
func HasDuplicates[T comparable](collection []T) bool {
	length := len(collection)
	seen := make(map[T]bool, length)
	for _, v := range collection {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = true
	}
	return false
}
