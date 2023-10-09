package goutils

import (
	. "github.com/mudssky/goutils/constraints"
)

// 泛型，绝对值
func Abs[T Number](num T) (res T) {
	if num < 0 {
		return -num
	}
	return num
}

// 泛型，返回最大值
// 暂不对参数数量不符合的清情况做特殊处理
func Max[T Number](nums ...T) (res T) {
	length := len(nums)
	if length < 1 {
		panic("arguments is empty")
	}
	res = nums[0]
	for i := 1; i < length; i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return
}
