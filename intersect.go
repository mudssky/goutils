package goutils

// Intersect returns the intersection between two collections.
//
// 两个列表取交集
func Intersect[T comparable](list1 []T, list2 []T) []T {
	// 预分配结果切片容量为较小列表的长度（交集最大可能大小）
	minLen := len(list1)
	if len(list2) < minLen {
		minLen = len(list2)
	}
	result := make([]T, 0, minLen)
	seen := map[T]struct{}{}

	for _, elem := range list1 {
		seen[elem] = struct{}{}
	}

	for _, elem := range list2 {
		if _, ok := seen[elem]; ok {
			result = append(result, elem)
		}
	}

	return result
}

// IntersectN 求任意多个数组的交集数组(n>=1),注意传入的数组都需要经过去重
func IntersectN[T comparable](first []T, arrays ...[]T) []T {
	length := len(arrays) + 1

	// 找出所有数组中最小的长度，作为结果切片的预分配容量
	minLen := len(first)
	for _, a := range arrays {
		if len(a) < minLen {
			minLen = len(a)
		}
	}

	// 创建一个 map，用于统计元素在几个数组中出现过
	m := make(map[T]int)

	for _, v := range first {
		m[v]++
	}
	for _, a := range arrays {
		for _, v := range a {
			m[v]++
		}
	}

	res := make([]T, 0, minLen)
	// 遍历 map，找出出现在所有数组中的元素
	for k, v := range m {
		if v == length {
			res = append(res, k)
		}
	}

	return res
}

// Union returns all distinct elements from given collections.
// result returns will not change the order of elements relatively.
//
// 求并集
func Union[T comparable](lists ...[]T) []T {
	// 计算所有列表的总长度作为结果切片的预分配容量（并集最大可能大小）
	totalLen := 0
	for _, list := range lists {
		totalLen += len(list)
	}

	result := make([]T, 0, totalLen)
	seen := map[T]struct{}{}

	for _, list := range lists {
		for _, e := range list {
			if _, ok := seen[e]; !ok {
				seen[e] = struct{}{}
				result = append(result, e)
			}
		}
	}

	return result
}

// Uniq returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array.
//
// 去重
func Uniq[T comparable](collection []T) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[T]struct{}, len(collection))

	for _, item := range collection {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// UniqBy returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
// The order of result values is determined by the order they occur in the array. It accepts `iteratee` which is
// invoked for each element in array to generate the criterion by which uniqueness is computed.
//
// 使用iteratee映射来判断重复
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	result := make([]T, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for _, item := range collection {
		key := iteratee(item)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, item)
	}

	return result
}

// Contains returns true if an element is present in a collection.
//
// 类似js array.contains
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}

	return false
}

// Includes Contains的别名
func Includes[T comparable](collection []T, element T) bool {
	return Contains(collection, element)
}

// ContainsBy returns true if predicate function return true.
//
// 用函数判断Contain
func ContainsBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// Every returns true if all elements of a subset are contained into a collection or if the subset is empty.
//
// 判断subset子集中的内容是否都在collection列表中存在
func Every[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if !Contains(collection, elem) {
			return false
		}
	}

	return true
}

// EveryBy returns true if the predicate returns true for all of the elements in the collection or if the collection is empty.
//
// 类似js的array.every
func EveryBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// Some returns true if at least 1 element of a subset is contained into a collection.
// If the subset is empty Some returns false.
func Some[T comparable](collection []T, subset []T) bool {
	// 如果子集为空，则返回 false
	if len(subset) == 0 {
		return false
	}

	for _, elem := range subset {
		if Contains(collection, elem) {
			return true
		}
	}

	return false
}

// SomeBy returns true if the predicate returns true for any of the elements in the collection.
// If the collection is empty SomeBy returns false.
func SomeBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return true
		}
	}

	return false
}

// None returns true if no element of a subset are contained into a collection or if the subset is empty.
func None[T comparable](collection []T, subset []T) bool {
	for _, elem := range subset {
		if Contains(collection, elem) {
			return false
		}
	}

	return true
}

// NoneBy returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.
func NoneBy[T any](collection []T, predicate func(item T) bool) bool {
	for _, v := range collection {
		if predicate(v) {
			return false
		}
	}

	return true
}
