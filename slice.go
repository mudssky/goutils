package goutils

import (
	"errors"
	"math/rand"

	c "github.com/mudssky/goutils/constraints"
)

// Chunk returns an array of elements split into groups the length of size. If array can't be split evenly,
// the final chunk will be the remaining elements.
//
// 对切片按照大小拆分
//
// Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
// [][]int{{0, 1}, {2, 3}, {4, 5}}
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	chunksNum := len(collection) / size
	if len(collection)%size != 0 {
		chunksNum += 1
	}

	result := make([][]T, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(collection) {
			last = len(collection)
		}
		result = append(result, collection[i*size:last])
	}

	return result
}

// Compact returns a slice of all non-zero elements.
//
// 返回所有非零值元素的集合
func Compact[T comparable](collection []T) []T {
	var zero T

	// 预分配结果切片容量为原切片长度（最大可能大小）
	result := make([]T, 0, len(collection))

	for _, item := range collection {
		if item != zero {
			result = append(result, item)
		}
	}

	return result
}

// Map manipulates a slice and transforms it to a slice of another type.
// It applies the iteratee function to each element in the collection and returns a new slice
// containing the results of each iteratee invocation.
//
// The iteratee function receives each item and its index, and should return a value of the new type.
//
// 将切片中的每个元素通过iteratee函数转换为新类型
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

// Filter iterates over elements of collection, returning a slice of all elements for which the predicate returns true.
// The predicate function receives each item and its index, and should return a boolean value.
//
// The function creates a new slice and does not modify the original collection.
//
// 过滤切片，返回所有满足条件的元素
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))
	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}

// FilterMap returns a slice which obtained after both filtering and mapping using the given callback function.
// The callback function should return two values:
//   - the result of the mapping operation and
//   - whether the result element should be included or not.
//
// 这个是lodash里面没有的，因为callback是同时执行filter和map的逻辑的。lodash里面都是单一功能函数
func FilterMap[T any, R any](collection []T, callback func(item T, index int) (R, bool)) []R {
	result := []R{}
	for i, item := range collection {
		if r, ok := callback(item, i); ok {
			result = append(result, r)
		}
	}
	return result
}

// FlatMap transforms each element of a slice into a slice of another type and flattens the result.
// Unlike Map which returns a single value for each element, FlatMap returns a slice of values for each element
// and then concatenates all resulting slices into a single slice.
//
// The iteratee function receives each item and its index, and should return a slice of the new type.
//
// FlatMap和Map的区别是可以对切片的每一项执行一个类型转换，并将结果展平为一个切片
func FlatMap[T any, R any](collection []T, iteratee func(item T, index int) []R) []R {
	result := make([]R, 0, len(collection))

	for i, item := range collection {
		result = append(result, iteratee(item, i)...)
	}

	return result
}

// Drop drops n elements from the beginning of a slice or array.
//
// 移除前个元素
func Drop[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return []T{}
	}

	result := make([]T, 0, len(collection)-n)

	return append(result, collection[n:]...)
}

// DropRight drops n elements from the end of a slice or array.
//
// 移除后n个元素
func DropRight[T any](collection []T, n int) []T {
	if len(collection) <= n {
		return []T{}
	}

	result := make([]T, 0, len(collection)-n)
	return append(result, collection[:len(collection)-n]...)
}

// DropWhile drops elements from the beginning of a slice or array while the predicate returns true.
//
// 去除predicate判断的第一个假值位置开始后面的元素.
func DropWhile[T any](collection []T, predicate func(item T) bool) []T {
	i := 0
	for ; i < len(collection); i++ {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, 0, len(collection)-i)
	return append(result, collection[i:]...)
}

// DropRightWhile drops elements from the end of a slice or array while the predicate returns true.
//
// 如果predicate返回true，移除后n个元素
func DropRightWhile[T any](collection []T, predicate func(item T) bool) []T {
	i := len(collection) - 1
	for ; i >= 0; i-- {
		if !predicate(collection[i]) {
			break
		}
	}

	result := make([]T, 0, i+1)
	return append(result, collection[:i+1]...)
}

// Fill fills elements of array with `initial` value.
//
// 需要实现Clonable接口,对基本类型的切片来说很不方便
// 但是估计未来官方应该要内置一个类似Clonable的泛型,这样才比较好写
func Fill[T c.Clonable[T]](collection []T, initial T) []T {
	result := make([]T, 0, len(collection))

	for range collection {
		result = append(result, initial.Clone())
	}

	return result
}

// Flatten returns an array a single level deep.
//
// 展平二维数组
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make([]T, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
//
// 数组反向
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}

// Count counts the c.Number of elements in the collection that compare equal to value.
//
// 返回集合中的等于某个值元素个数
func Count[T comparable](collection []T, value T) (count int) {
	for _, item := range collection {
		if item == value {
			count++
		}
	}

	return count
}

// CountBy counts the c.Number of elements in the collection for which predicate is true.
//
// 同Count，但是用函数判断
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	for _, item := range collection {
		if predicate(item) {
			count++
		}
	}

	return count
}

// CountValues counts the c.Number of each element in the collection.
//
// 统计集合中的元素，返回一个Map
func CountValues[T comparable](collection []T) map[T]int {
	result := make(map[T]int)

	for _, item := range collection {
		result[item]++
	}

	return result
}

// CountValuesBy counts the c.Number of each element return from mapper function.
// Is equivalent to chaining Map and CountValues.
//
// 同CountValues，但是使用一个映射函数
func CountValuesBy[T any, U comparable](collection []T, mapper func(item T) U) map[U]int {
	result := make(map[U]int)

	for _, item := range collection {
		result[mapper(item)]++
	}

	return result
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
//
// 类似js的ForEach方法
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// GroupBy returns an object composed of keys generated from the results of running each element of collection through iteratee.
//
// 使用函数映射(传入值)来分组
//
//	groups := goutils.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
//	    return i%3
//	})
//
// map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}
func GroupBy[T any, U comparable](collection []T, iteratee func(item T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}

	return result
}

// KeyBy transforms a slice or an array of structs to a map based on a pivot callback.
//
// 和GroupBy类似，但是不分组，也就是最后得到的map，值不是切片
func KeyBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]V {
	result := make(map[K]V, len(collection))

	for _, v := range collection {
		k := iteratee(v)
		result[k] = v
	}

	return result
}

// PartitionBy returns an array of elements split into groups. The order of grouped values is
// determined by the order they occur in collection. The grouping is generated from the results
// of running each element of collection through iteratee.
//
// 根据函数映射返回的不同值分组，返回一个二维切片
//
//	partitions := goutils.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
//	    if x < 0 {
//	        return "negative"
//	    } else if x%2 == 0 {
//	        return "even"
//	    }
//	    return "odd"
//	})
//
// [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}
func PartitionBy[T any, K comparable](collection []T, iteratee func(item T) K) [][]T {
	result := [][]T{}
	seen := map[K]int{}

	for _, item := range collection {
		key := iteratee(item)

		resultIndex, ok := seen[key]
		if !ok {
			resultIndex = len(result)
			seen[key] = resultIndex
			result = append(result, []T{})
		}

		result[resultIndex] = append(result[resultIndex], item)
	}

	return result

	// unc.Ordered:
	// groups := GroupBy[T, K](collection, iteratee)
	// return Values[K, []T](groups)
}

// Reduce reduces collection to a value which is the accumulated result of running each element in collection
// through accumulator, where each successive invocation is supplied the return value of the previous.
//
// 经典的Reduce方法
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}

	return initial
}

// ReduceRight helper is like Reduce except that it iterates over elements of collection from right to left.
//
// 从右向左的Reduce
func ReduceRight[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i := len(collection) - 1; i >= 0; i-- {
		initial = accumulator(initial, collection[i], i)
	}

	return initial
}

// Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
//
// 对切片数组原地洗牌
func Shuffle[T any](collection []T) []T {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})

	return collection
}

// Times invokes the iteratee n times, returning an array of the results of each invocation.
// The iteratee is invoked with index as argument.
//
// 执行iteratee count次数
func Times[T any](count int, iteratee func(index int) T) []T {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		result[i] = iteratee(i)
	}

	return result
}

// Interleave round-robin alternating input slices and sequentially appending value at index into result
//
// 交替输入切片到结果
//
// interleaved := goutils.Interleave([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//
// interleaved := goutils.Interleave([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func Interleave[T any](collections ...[]T) []T {
	if len(collections) == 0 {
		return []T{}
	}

	maxSize := 0
	totalSize := 0
	for _, c := range collections {
		size := len(c)
		totalSize += size
		if size > maxSize {
			maxSize = size
		}
	}

	if maxSize == 0 {
		return []T{}
	}

	// 预分配结果切片容量为所有切片长度之和（结果最大可能大小）
	result := make([]T, totalSize)

	resultIdx := 0
	for i := 0; i < maxSize; i++ {
		for j := range collections {
			if len(collections[j])-1 < i {
				continue
			}

			result[resultIdx] = collections[j][i]
			resultIdx++
		}
	}

	return result
}

// Repeat builds a slice with N copies of initial value.
//
// 重复initial n次，返回一个切片
func Repeat[T c.Clonable[T]](count int, initial T) []T {
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		result = append(result, initial.Clone())
	}

	return result
}

// Associate returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
//
// 传入对象切片，然后从对象中取值拼接成map
func Associate[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	result := make(map[K]V, len(collection))

	for _, t := range collection {
		k, v := transform(t)
		result[k] = v
	}

	return result
}

// SliceToMap returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.
// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.
// Alias of Associate().
func SliceToMap[T any, K comparable, V any](collection []T, transform func(item T) (K, V)) map[K]V {
	return Associate(collection, transform)
}

// SliceSafe returns a copy of a slice from `start` up to, but not including `end`. Like `slice[start:end]`, but does not panic on overflow.
// 类似于slice[start:end]，但是不会因为超出范围panic
func SliceSafe[T any](collection []T, start int, end int) []T {
	size := len(collection)

	if start >= end {
		return []T{}
	}

	if start > size {
		start = size
	}
	if start < 0 {
		start = 0
	}

	if end > size {
		end = size
	}
	if end < 0 {
		end = 0
	}

	return collection[start:end]
}

// ReplaceN returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
//
// 替换元素前n个,返回一个副本
func ReplaceN[T comparable](collection []T, old T, new T, n int) []T {
	result := make([]T, len(collection))
	copy(result, collection)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll returns a copy of the slice with all non-overlapping instances of old replaced by new.
//
// 替换所有元素,返回一个副本
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return ReplaceN(collection, old, new, -1)
}

// IsSorted checks if a slice is sorted.
//
// 判断切片是否排序好
func IsSorted[T c.Ordered](collection []T) bool {
	for i := 1; i < len(collection); i++ {
		if collection[i-1] > collection[i] {
			return false
		}
	}

	return true
}

// IsSortedByKey checks if a slice is sorted by iteratee.
//
// 使用函数映射后判断排序
func IsSortedByKey[T any, K c.Ordered](collection []T, iteratee func(item T) K) bool {
	size := len(collection)

	for i := 0; i < size-1; i++ {
		if iteratee(collection[i]) > iteratee(collection[i+1]) {
			return false
		}
	}

	return true
}

// Range 返回[start,end)区间的切片
func Range[T c.Number](start, end T) (res []T, err error) {
	if start > end {
		return nil, errors.New("start must be less than end")
	}
	return RangeWithStep(start, end, 1)
}

// RangeWithStep 返回[start,end，step)区间的切片，step可以为负
func RangeWithStep[T c.Number](start, end, step T) (res []T, err error) {
	empty := Empty[[]T]()
	if step == 0 {
		return empty, errors.New("step cannot be zero")
	}

	if (start < end && step < 0) || (start > end && step > 0) {
		return empty, errors.New("step direction is inconsistent with start and end values")
	}

	distance := Abs(end - start)
	// 暂时不处理可能的溢出问题。
	size := Abs(int((distance) / step))

	res = make([]T, size)
	if size > 0 {
		res[0] = start
	}

	for i := 1; i < size; i++ {
		res[i] = res[i-1] + step
	}
	return
}

// Concat 返回一个拼接好的新数组
func Concat[T any](collection []T, values ...T) []T {
	result := make([]T, 0, len(collection)+len(values))
	result = append(result, collection...)
	result = append(result, values...)
	return result
}

// Difference 生成第一个数组中独有的元素组成的新数组
// collection是检查的数组,excludes是需要排除的值的数组
func Difference[T comparable](collection []T, excludes []T) []T {
	// 预分配结果切片容量为collection的长度（结果最大可能大小）
	result := make([]T, 0, len(collection))
	// 初始化需要排除的map,便于查找
	excludesMap := map[T]struct{}{}
	for _, elem := range excludes {
		excludesMap[elem] = struct{}{}
	}

	// 遍历第一个数组,放入第二个数组中不存在的元素,
	for _, elem := range collection {
		if _, ok := excludesMap[elem]; !ok {
			result = append(result, elem)
		}
	}

	return result
}

// 生成第一个数组中独有的元素组成的新数组,对比较一个元素调用iteratee转换后的结果
func DifferenceBy[T comparable, R comparable](collection []T, excludes []T, iteratee func(item T) R) []T {
	// 预分配结果切片容量为collection的长度（结果最大可能大小）
	result := make([]T, 0, len(collection))
	// 初始化需要排除的map,便于查找
	excludesMap := map[R]struct{}{}
	for _, elem := range excludes {
		exclude := iteratee(elem)
		excludesMap[exclude] = struct{}{}
	}
	// 遍历第一个数组,放入第二个数组中不存在的元素,
	for _, elem := range collection {
		compareElem := iteratee(elem)
		if _, ok := excludesMap[compareElem]; !ok {
			result = append(result, elem)
		}
	}

	return result
}

// SortedIndex 判断一个数据应该放入增序排序数组的哪个位置
func SortedIndex[T c.Ordered](array []T, value T) int {
	length := len(array)
	if length == 0 {
		return 0
	}
	for i := 0; i < length; i++ {
		if value < array[i] {
			return i
		}
	}
	return length
}

// 创建一个数组，包含所有数组中的唯一值
func Xor[T comparable](arrays ...[]T) []T {
	// 创建一个 map，用于统计元素在几个数组中出现过
	m := make(map[T]int)

	// 计算所有数组的总长度作为结果切片的预分配容量（最大可能大小）
	totalLen := 0
	for _, a := range arrays {
		totalLen += len(a)
		for _, v := range a {
			m[v]++
		}
	}

	res := make([]T, 0, totalLen)
	// 遍历 map，找出出现在所有数组中的元素
	for k, v := range m {
		if v < 2 {
			res = append(res, k)
		}
	}

	return res
}

// 反向遍历
func ForEachRight[T any](collection []T, iteratee func(item T, index int)) {
	for i := len(collection) - 1; i >= 0; i-- {
		iteratee(collection[i], i)
	}
}
