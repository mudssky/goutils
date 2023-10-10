package goutils

import (
	c "github.com/mudssky/goutils/constraints"
)

// Abs 泛型，绝对值
func Abs[T c.Number](num T) (res T) {
	if num < 0 {
		return -num
	}
	return num
}

// MaxN 泛型，返回多个数最大值
// 暂不对参数数量不符合的清情况做特殊处理
func MaxN[T c.Number](first T, others ...T) (res T) {
	res = first
	for _, v := range others {
		if v > res {
			res = v
		}
	}
	return
}

// Max 泛型，返回两个数中最大值
func Max[T c.Number](num1 T, num2 T) (res T) {
	return MaxN(num1, num2)
}
