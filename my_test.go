package goutils

import (
	"math"
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	type args struct {
		collection []int
		values     []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3}, []int{4, 5}}, []int{1, 2, 3, 4, 5}},
		{"test2", args{[]int{2, 3}, []int{4, 5}}, []int{2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.collection, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceBy(t *testing.T) {
	type args struct {
		collection []float64
		excludes   []float64
		// iteratee   func(item int) string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"test1", args{[]float64{3.1, 2.2, 1.3}, []float64{4.4, 2.5}}, []float64{3.1, 1.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceBy(tt.args.collection, tt.args.excludes, math.Floor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args struct {
		collection []int
		excludes   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{3, 2, 1}, []int{4, 2}}, []int{3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.collection, tt.args.excludes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceWith(t *testing.T) {
	type args struct {
		collection []int
		excludes   []string
		// comparator func(left int, right string) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{3, 2, 1}, []string{"test4", "test2"}}, []int{3, 1}},
		{"test1", args{[]int{3, 2, 1}, []string{"test1", "test2"}}, []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceWith(tt.args.collection, tt.args.excludes, func(left int, right string) bool {
				num, err := strconv.Atoi(right[4:])
				if err != nil {
					t.Errorf("DifferenceWith atoi  failed %v", tt.name)
				}
				if num == left {
					return true
				}
				return false
			}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceWith() = %v, want %v", got, tt.want)
			}
		})
	}
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
		{"test1", args{[]string{"test1", "test2", "test9"}}, 2},
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

func TestSortedIndex(t *testing.T) {
	type args struct {
		array []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"最小", args{[]int{1, 2, 5}, 0}, 0},
		{"最大", args{[]int{1, 2, 5}, 6}, 3},
		{"空数组", args{[]int{}, 6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedIndex(tt.args.array, tt.args.value); got != tt.want {
				t.Errorf("SortedIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXor(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[][]int{{3, 2, 1}, {4, 2}}}, []int{3, 1, 4}},
		{"test2", args{[][]int{{2, 1}, {2, 3}}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Xor(tt.args.arrays...)
			sort.Ints(got)
			sort.Ints(tt.want)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestForEachRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEachRight([]string{"a", "b", "c"}, func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.ElementsMatch([]string{"c", "b", "a"}, callParams1)
	is.ElementsMatch([]int{2, 1, 0}, callParams2)
	is.IsDecreasing(callParams2)
}
