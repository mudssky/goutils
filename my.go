package goutils

// 返回一个拼接好的新数组
func Concat[T any](collection []T, values ...T) []T {
	result := make([]T, 0, len(collection)+len(values))
	result = append(result, collection...)
	result = append(result, values...)
	return result
}

// 生成第一个数组中独有的元素组成的新数组
// collection是检查的数组,excludes是需要排除的值的数组
func Difference[T comparable](collection []T, excludes []T) []T {
	result := []T{}
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
	result := []T{}
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

// 和Difference一样,但是用一个比较函数判断
func DifferenceWith[T comparable, R comparable](collection []T, excludes []R, comparator func(left T, right R) bool) []T {
	result := []T{}
	// 缓存map,把已经比较的结果存起来,加速重复的比较,通过的是true
	cacheMap := map[T]bool{}

	// 遍历第一个数组,放入第二个数组中不存在的元素,
	for _, elem := range collection {
		// 用缓存提升执行效率,如果缓存中比较过了通过,那么直接放入到结果中
		if isPass, ok := cacheMap[elem]; ok {
			// 不通过的情况跳过
			if !isPass {
				continue
			}
			result = append(result, elem)
		}
		isInclude := false
		// 缓存中没有比较过,遍历excludes比较
		for _, ex := range excludes {
			isEqual := comparator(elem, ex)
			if isEqual {
				isInclude = true
				break
			}
		}
		if isInclude {
			cacheMap[elem] = false

		} else {
			// 不包含的情况添加
			result = append(result, elem)
			cacheMap[elem] = true
		}
	}

	return result
}

// This method is like _.find except that it returns the index of the first element predicate returns truthy for instead of the element itself.
// 和find一样,但是返回的是数组下标
func FindIndex[T comparable](collection []T, predicate func(item T) bool) int {
	for i, item := range collection {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// 类似于FindIndex,但是从后往前找
func FindLastIndex[T comparable](collection []T, predicate func(item T) bool) int {
	length := len(collection)
	for i := length - 1; i >= 0; i-- {
		if predicate(collection[i]) {
			return i
		}
	}
	return -1
}

// 求任意多个数组的交集数组,注意传入的数组都需要经过去重
func Intersect2[T comparable](arrays ...[]T) []T {
	// 创建一个 map，用于统计元素在几个数组中出现过
	m := make(map[T]int)
	for _, a := range arrays {
		for _, v := range a {
			m[v]++
		}
	}

	res := []T{}
	// 遍历 map，找出出现在所有数组中的元素
	for k, v := range m {
		if v == len(arrays) {
			res = append(res, k)
		}
	}

	return res
}

// 判断一个数据应该放入增序排序数组的哪个位置
func SortedIndex[T Ordered](array []T, value T) int {
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

// 判断一个数据应该放入增序排序数组的哪个位置,先用iteratee处理后比较
func SortedIndexBy[T Ordered](array []T, value T, iteratee func(item T, index int)) int {
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
	for _, a := range arrays {
		for _, v := range a {
			m[v]++
		}
	}

	res := []T{}
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
