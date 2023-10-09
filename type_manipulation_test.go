package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ToPtr([]int{1, 2})

	is.Equal(*result1, []int{1, 2})
}

func TestEmptyableToPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Nil(NullableToPtr(0))
	is.Nil(NullableToPtr(""))
	is.Nil(NullableToPtr[[]int](nil))
	is.Nil(NullableToPtr[map[int]int](nil))
	is.Nil(NullableToPtr[error](nil))

	is.Equal(*NullableToPtr(42), 42)
	is.Equal(*NullableToPtr("nonempty"), "nonempty")
	is.Equal(*NullableToPtr([]int{}), []int{})
	is.Equal(*NullableToPtr([]int{1, 2}), []int{1, 2})
	is.Equal(*NullableToPtr(map[int]int{}), map[int]int{})
	is.Equal(*NullableToPtr(assert.AnError), assert.AnError)
}

func TestFromPtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	ptr := &str1

	is.Equal("foo", FromPtr(ptr))
	is.Equal("", FromPtr[string](nil))
	is.Equal(0, FromPtr[int](nil))
	is.Nil(FromPtr[*string](nil))
	is.EqualValues(ptr, FromPtr(&ptr))
}

func TestFromPtrOr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	const fallbackStr = "fallback"
	str := "foo"
	ptrStr := &str

	const fallbackInt = -1
	i := 9
	ptrInt := &i

	is.Equal(str, FromPtrOr(ptrStr, fallbackStr))
	is.Equal(fallbackStr, FromPtrOr(nil, fallbackStr))
	is.Equal(i, FromPtrOr(ptrInt, fallbackInt))
	is.Equal(fallbackInt, FromPtrOr(nil, fallbackInt))
}

func TestToSlicePtr(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	str1 := "foo"
	str2 := "bar"
	result1 := ToSlicePtr([]string{str1, str2})

	is.Equal(result1, []*string{&str1, &str2})
}

func TestToAnySlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in1 := []int{0, 1, 2, 3}
	in2 := []int{}
	out1 := ToAnySlice(in1)
	out2 := ToAnySlice(in2)

	is.Equal([]any{0, 1, 2, 3}, out1)
	is.Equal([]any{}, out2)
}

func TestFromAnySlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.NotPanics(func() {
		out1, ok1 := FromAnySlice[string]([]any{"foobar", 42})
		out2, ok2 := FromAnySlice[string]([]any{"foobar", "42"})

		is.Equal([]string{}, out1)
		is.False(ok1)
		is.Equal([]string{"foobar", "42"}, out2)
		is.True(ok2)
	})
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	//nolint:unused
	type test struct{}

	is.Empty(Empty[string]())
	is.Empty(Empty[int64]())
	is.Empty(Empty[test]())
	is.Empty(Empty[chan string]())
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	//nolint:unused
	type test struct {
		foobar string
	}

	is.True(IsEmpty(""))
	is.False(IsEmpty("foo"))
	is.True(IsEmpty[int64](0))
	is.False(IsEmpty[int64](42))
	is.True(IsEmpty(test{foobar: ""}))
	is.False(IsEmpty(test{foobar: "foo"}))
}

func TestIsNotEmpty(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	//nolint:unused
	type test struct {
		foobar string
	}

	is.False(IsNotEmpty(""))
	is.True(IsNotEmpty("foo"))
	is.False(IsNotEmpty[int64](0))
	is.True(IsNotEmpty[int64](42))
	is.False(IsNotEmpty(test{foobar: ""}))
	is.True(IsNotEmpty(test{foobar: "foo"}))
}
