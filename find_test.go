package goutils

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := IndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 2)
	is.Equal(result2, -1)
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 2)
	result2 := LastIndexOf([]int{0, 1, 2, 1, 2, 3}, 6)

	is.Equal(result1, 4)
	is.Equal(result2, -1)
}

func TestFind(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	result1, ok1 := Find([]string{"a", "b", "c", "d"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d"}[index], item)
		index++
		return item == "b"
	})

	result2, ok2 := Find([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(ok1, true)
	is.Equal(result1, "b")
	is.Equal(ok2, false)
	is.Equal(result2, "")
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"a", "b", "c", "d", "b"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 1)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestFindLastIndexOf(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	index := 0
	item1, index1, ok1 := FindLastIndexOf([]string{"a", "b", "c", "d", "b"}, func(item string) bool {
		is.Equal([]string{"b", "d", "c", "b", "a"}[index], item)
		index++
		return item == "b"
	})
	item2, index2, ok2 := FindLastIndexOf([]string{"foobar"}, func(item string) bool {
		is.Equal("foobar", item)
		return item == "b"
	})

	is.Equal(item1, "b")
	is.Equal(ok1, true)
	is.Equal(index1, 4)
	is.Equal(item2, "")
	is.Equal(ok2, false)
	is.Equal(index2, -1)
}

func TestNth(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1, err1 := Nth([]int{0, 1, 2, 3}, 2)
	result2, err2 := Nth([]int{0, 1, 2, 3}, -2)
	result3, err3 := Nth([]int{0, 1, 2, 3}, 42)
	result4, err4 := Nth([]int{}, 0)
	result5, err5 := Nth([]int{42}, 0)
	result6, err6 := Nth([]int{42}, -1)

	is.Equal(result1, 2)
	is.Equal(err1, nil)
	is.Equal(result2, 2)
	is.Equal(err2, nil)
	is.Equal(result3, 0)
	is.Equal(err3, fmt.Errorf("nth: 42 out of slice bounds"))
	is.Equal(result4, 0)
	is.Equal(err4, fmt.Errorf("nth: 0 out of slice bounds"))
	is.Equal(result5, 42)
	is.Equal(err5, nil)
	is.Equal(result6, 42)
	is.Equal(err6, nil)
}

func TestMaxOfCollection(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := MaxOfCollection([]int{1, 2, 3})
	result2 := MaxOfCollection([]int{3, 2, 1})
	result3 := MaxOfCollection([]int{})

	is.Equal(result1, 3)
	is.Equal(result2, 3)
	is.Equal(result3, 0)
}
func TestSample(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Sample([]string{"a", "b", "c"})
	result2 := Sample([]string{})

	is.True(Contains([]string{"a", "b", "c"}, result1))
	is.Equal(result2, "")
}

func TestSamples(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	rand.Seed(time.Now().UnixNano())

	result1 := Samples([]string{"a", "b", "c"}, 3)
	result2 := Samples([]string{}, 3)

	sort.Strings(result1)

	is.Equal(result1, []string{"a", "b", "c"})
	is.Equal(result2, []string{})
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

func TestFindIndex(t *testing.T) {
	type args struct {
		array []string
		// predicate func(item string) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]string{"test1", "test2", "test3"}}, -1},
		{"test1", args{[]string{"test1", "test2", "test9"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndex(tt.args.array, func(item string) bool {
				return item == "test9"
			}); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	type args struct {
		array []string
		// predicate func(item string) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]string{"test1", "test9", "test2", "test3", "test9"}}, 4},
		{"test2", args{[]string{"test1", "test2", "test9"}}, 2},
		{"test3", args{[]string{"test1", "test2"}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLastIndex(tt.args.array, func(item string) bool {
				return item == "test9"
			}); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasDuplicates(t *testing.T) {
	type args struct {
		collection []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"has dup", args{collection: []string{"a", "b", "ggg", "d", "e", "ggg", "g", "h", "i"}}, true},
		{"no dup", args{collection: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicates(tt.args.collection); got != tt.want {
				t.Errorf("HasDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
