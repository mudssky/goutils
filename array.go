package goutils

import (
	"errors"
)

// 测试方法注释222
func Range[T Number](start, end T) (res []T, err error) {
	if start > end {
		return nil, errors.New("start must be less than end")
	}
	return RangeWithStep(start, end, 1)
}

func RangeWithStep[T Number](start, end, step T) (res []T, err error) {
	if step == 0 {
		return nil, errors.New("step cannot be zero")
	}

	if (start < end && step < 0) || (start > end && step > 0) {
		return nil, errors.New("step direction is inconsistent with start and end values")
	}

	// 暂时不处理可能的溢出问题。
	size := Abs(int((end - start) / step))
	if size == 0 {
		size = 1
	}

	res = make([]T, size)
	res[0] = start
	for i := 1; i < size; i++ {
		res[i] = res[i-1] + step
	}
	return
}
