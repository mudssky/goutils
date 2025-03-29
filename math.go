package goutils

import (
	c "github.com/mudssky/goutils/constraints"
)

// Abs returns the absolute value of a number.
// It works with any numeric type that satisfies the Number constraint.
//
// Abs 泛型，绝对值
func Abs[T c.Number](num T) (res T) {
	if num < 0 {
		return -num
	}
	return num
}

// MaxN returns the maximum value from a set of numbers.
// It takes a first value and a variadic list of additional values, comparing them to find the maximum.
// The function does not perform special handling for invalid parameter counts.
//
// Deprecated: After Go 1.21, use the built-in min and max functions instead.
//
// MaxN 泛型，返回多个数最大值
// 暂不对参数数量不符合的清情况做特殊处理
//
// Deprecated: go 1.21版本后，官方已经内置了min和max函数，弃用
func MaxN[T c.Number](first T, others ...T) (res T) {
	res = first
	for _, v := range others {
		if v > res {
			res = v
		}
	}
	return
}

// Max returns the maximum of two numbers.
// It compares two values and returns the larger one.
//
// Deprecated: After Go 1.21, use the built-in min and max functions instead.
//
// # Max 泛型，返回两个数中最大值
//
// Deprecated: go 1.21版本后，官方已经内置了min和max函数，弃用
func Max[T c.Number](num1 T, num2 T) (res T) {
	return MaxN(num1, num2)
}
