package goutils

import "reflect"

// 返回指定类型的空值
func Empty[T any]() T {
	var zero T
	return zero
}

// ToPtr returns a pointer copy of value.
//
// 返回指针
func ToPtr[T any](x T) *T {
	return &x
}

// 如果为非零值，返回指针，否则返回nil
func NullableToPtr[T any](x T) *T {
	isZero := reflect.ValueOf(&x).Elem().IsZero()
	if isZero {
		return nil
	}
	return &x
}

// FromPtr returns the pointer value or empty.
// 从指针获取值，空指针nil返回对应的零值
func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}

	return *x
}

// FromPtrOr returns the pointer value or the fallback value.
// 从指针获取值，为nil返回fallback
func FromPtrOr[T any](x *T, fallback T) T {
	if x == nil {
		return fallback
	}

	return *x
}

// ToSlicePtr returns a slice of pointer copy of value.
// 传入列表，返回指针的切片
func ToSlicePtr[T any](collection []T) []*T {
	return Map(collection, func(x T, _ int) *T {
		return &x
	})
}

// ToAnySlice returns a slice with all elements mapped to `any` type
// 传入一个切片，返回一个any类型的切片
func ToAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i, item := range collection {
		result[i] = item
	}
	return result
}

// FromAnySlice returns an `any` slice with all elements mapped to a type.
// Returns false in case of type conversion failure.
// 从any切片转为有类型的切片
func FromAnySlice[T any](in []any) (out []T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			out = []T{}
			ok = false
		}
	}()

	result := make([]T, len(in))
	for i, item := range in {
		result[i] = item.(T)
	}
	return result, true
}

// IsEmpty returns true if argument is a zero value.
// 判断是否为零值
func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

// IsNotEmpty returns true if argument is not a zero value.
// 判断是否不是零值
func IsNotEmpty[T comparable](v T) bool {
	return !IsEmpty[T](v)
}
