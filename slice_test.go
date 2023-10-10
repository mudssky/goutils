package goutils

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})

	is.Equal(r1, []int{2, 4})

	r2 := Filter([]string{"", "foo", "", "bar", ""}, func(x string, _ int) bool {
		return len(x) > 0
	})

	is.Equal(r2, []string{"foo", "bar"})
}

func TestMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Map([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := Map([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	is.Equal(len(result1), 4)
	is.Equal(len(result2), 4)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "3", "4"})
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := FilterMap([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, bool) {
		if x%2 == 0 {
			return strconv.FormatInt(x, 10), true
		}
		return "", false
	})
	r2 := FilterMap([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})

	is.Equal(len(r1), 2)
	is.Equal(len(r2), 2)
	is.Equal(r1, []string{"2", "4"})
	is.Equal(r2, []string{"xpu", "xpu"})
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := FlatMap([]int{0, 1, 2, 3, 4}, func(x int, _ int) []string {
		return []string{"Hello"}
	})
	result2 := FlatMap([]int64{0, 1, 2, 3, 4}, func(x int64, _ int) []string {
		result := make([]string, 0, x)
		for i := int64(0); i < x; i++ {
			result = append(result, strconv.FormatInt(x, 10))
		}
		return result
	})

	is.Equal(len(result1), 5)
	is.Equal(len(result2), 10)
	is.Equal(result1, []string{"Hello", "Hello", "Hello", "Hello", "Hello"})
	is.Equal(result2, []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"})
}

func TestTimes(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []string{"0", "1", "2"})
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	result2 := Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 10)

	is.Equal(result1, 10)
	is.Equal(result2, 20)
}

func TestReduceRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})

	is.Equal(result1, []int{4, 5, 2, 3, 0, 1})
}

func TestForEach(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	// check of callback is called for every element and in proper order

	callParams1 := []string{}
	callParams2 := []int{}

	ForEach([]string{"a", "b", "c"}, func(item string, i int) {
		callParams1 = append(callParams1, item)
		callParams2 = append(callParams2, i)
	})

	is.ElementsMatch([]string{"a", "b", "c"}, callParams1)
	is.ElementsMatch([]int{0, 1, 2}, callParams2)
	is.IsIncreasing(callParams2)
}

func TestUniq(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Uniq([]int{1, 2, 2, 1})

	is.Equal(len(result1), 2)
	is.Equal(result1, []int{1, 2})
}

func TestUniqBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := UniqBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, []int{0, 1, 2})
}

func TestGroupBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	is.Equal(len(result1), 3)
	is.Equal(result1, map[int][]int{
		0: {0, 3},
		1: {1, 4},
		2: {2, 5},
	})
}

func TestChunk(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	result2 := Chunk([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	result3 := Chunk([]int{}, 2)
	result4 := Chunk([]int{0}, 2)

	is.Equal(result1, [][]int{{0, 1}, {2, 3}, {4, 5}})
	is.Equal(result2, [][]int{{0, 1}, {2, 3}, {4, 5}, {6}})
	is.Equal(result3, [][]int{})
	is.Equal(result4, [][]int{{0}})
	is.PanicsWithValue("Second parameter must be greater than 0", func() {
		Chunk([]int{0}, 0)
	})
}

func TestPartitionBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}

	result1 := PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	result2 := PartitionBy([]int{}, oddEven)

	is.Equal(result1, [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}})
	is.Equal(result2, [][]int{})
}

func TestFlatten(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Flatten([][]int{{0, 1}, {2, 3, 4, 5}})

	is.Equal(result1, []int{0, 1, 2, 3, 4, 5})
}

func TestInterleave(t *testing.T) {
	tests := []struct {
		name        string
		collections [][]int
		want        []int
	}{
		{
			"nil",
			[][]int{nil},
			[]int{},
		},
		{
			"empty",
			[][]int{},
			[]int{},
		},
		{
			"empties",
			[][]int{{}, {}},
			[]int{},
		},
		{
			"same length",
			[][]int{{1, 3, 5}, {2, 4, 6}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"different length",
			[][]int{{1, 3, 5, 6}, {2, 4}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"many slices",
			[][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interleave(tt.collections...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Shuffle([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result2 := Shuffle([]int{})

	is.NotEqual(result1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	is.Equal(result2, []int{})
}

func TestReverse(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Reverse([]int{0, 1, 2, 3, 4, 5})
	result2 := Reverse([]int{0, 1, 2, 3, 4, 5, 6})
	result3 := Reverse([]int{})

	is.Equal(result1, []int{5, 4, 3, 2, 1, 0})
	is.Equal(result2, []int{6, 5, 4, 3, 2, 1, 0})
	is.Equal(result3, []int{})
}

func TestFill(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Fill([]foo{{"a"}, {"a"}}, foo{"b"})
	result2 := Fill([]foo{}, foo{"a"})

	is.Equal(result1, []foo{{"b"}, {"b"}})
	is.Equal(result2, []foo{})
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := Repeat(2, foo{"a"})
	result2 := Repeat(0, foo{"a"})

	is.Equal(result1, []foo{{"a"}, {"a"}})
	is.Equal(result2, []foo{})
}

func TestKeyBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	result1 := KeyBy([]string{"a", "aa", "aaa"}, func(str string) int {
		return len(str)
	})

	is.Equal(result1, map[int]string{1: "a", 2: "aa", 3: "aaa"})
}

func TestAssociate(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in     []*foo
		expect map[string]int
	}{
		{
			in:     []*foo{{baz: "apple", bar: 1}},
			expect: map[string]int{"apple": 1},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(Associate(testCase.in, transform), testCase.expect)
		})
	}
}

func TestSliceToMap(t *testing.T) {
	t.Parallel()

	type foo struct {
		baz string
		bar int
	}
	transform := func(f *foo) (string, int) {
		return f.baz, f.bar
	}
	testCases := []struct {
		in     []*foo
		expect map[string]int
	}{
		{
			in:     []*foo{{baz: "apple", bar: 1}},
			expect: map[string]int{"apple": 1},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}},
			expect: map[string]int{"apple": 1, "banana": 2},
		},
		{
			in:     []*foo{{baz: "apple", bar: 1}, {baz: "apple", bar: 2}},
			expect: map[string]int{"apple": 2},
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			is := assert.New(t)
			is.Equal(SliceToMap(testCase.in, transform), testCase.expect)
		})
	}
}

func TestDrop(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{1, 2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{2, 3, 4}, Drop([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{3, 4}, Drop([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{4}, Drop([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, Drop([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropRight(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRight([]int{0, 1, 2, 3, 4}, 1))
	is.Equal([]int{0, 1, 2}, DropRight([]int{0, 1, 2, 3, 4}, 2))
	is.Equal([]int{0, 1}, DropRight([]int{0, 1, 2, 3, 4}, 3))
	is.Equal([]int{0}, DropRight([]int{0, 1, 2, 3, 4}, 4))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 5))
	is.Equal([]int{}, DropRight([]int{0, 1, 2, 3, 4}, 6))
}

func TestDropWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{4, 5, 6}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 4
	}))

	is.Equal([]int{}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return true
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, DropWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))
}

func TestDropRightWhile(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal([]int{0, 1, 2, 3}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 3
	}))

	is.Equal([]int{0, 1}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 1
	}))

	is.Equal([]int{0, 1, 2, 3, 4, 5, 6}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t == 10
	}))

	is.Equal([]int{}, DropRightWhile([]int{0, 1, 2, 3, 4, 5, 6}, func(t int) bool {
		return t != 10
	}))
}

func TestCount(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := Count([]int{1, 2, 1}, 1)
	count2 := Count([]int{1, 2, 1}, 3)
	count3 := Count([]int{}, 1)

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	count1 := CountBy([]int{1, 2, 1}, func(i int) bool {
		return i < 2
	})

	count2 := CountBy([]int{1, 2, 1}, func(i int) bool {
		return i > 2
	})

	count3 := CountBy([]int{}, func(i int) bool {
		return i <= 2
	})

	is.Equal(count1, 2)
	is.Equal(count2, 0)
	is.Equal(count3, 0)
}

func TestCountValues(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.Equal(map[int]int{}, CountValues([]int{}))
	is.Equal(map[int]int{1: 1, 2: 1}, CountValues([]int{1, 2}))
	is.Equal(map[int]int{1: 1, 2: 2}, CountValues([]int{1, 2, 2}))
	is.Equal(map[string]int{"": 1, "foo": 1, "bar": 1}, CountValues([]string{"foo", "bar", ""}))
	is.Equal(map[string]int{"foo": 1, "bar": 2}, CountValues([]string{"foo", "bar", "bar"}))
}

func TestCountValuesBy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	oddEven := func(v int) bool {
		return v%2 == 0
	}
	length := func(v string) int {
		return len(v)
	}

	result1 := CountValuesBy([]int{}, oddEven)
	result2 := CountValuesBy([]int{1, 2}, oddEven)
	result3 := CountValuesBy([]int{1, 2, 2}, oddEven)
	result4 := CountValuesBy([]string{"foo", "bar", ""}, length)
	result5 := CountValuesBy([]string{"foo", "bar", "bar"}, length)

	is.Equal(map[bool]int{}, result1)
	is.Equal(map[bool]int{true: 1, false: 1}, result2)
	is.Equal(map[bool]int{true: 2, false: 1}, result3)
	is.Equal(map[int]int{0: 1, 3: 2}, result4)
	is.Equal(map[int]int{3: 3}, result5)
}

func TestSlice(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 2, 3, 4}

	out1 := SliceSafe(in, 0, 0)
	out2 := SliceSafe(in, 0, 1)
	out3 := SliceSafe(in, 0, 5)
	out4 := SliceSafe(in, 0, 6)
	out5 := SliceSafe(in, 1, 1)
	out6 := SliceSafe(in, 1, 5)
	out7 := SliceSafe(in, 1, 6)
	out8 := SliceSafe(in, 4, 5)
	out9 := SliceSafe(in, 5, 5)
	out10 := SliceSafe(in, 6, 5)
	out11 := SliceSafe(in, 6, 6)
	out12 := SliceSafe(in, 1, 0)
	out13 := SliceSafe(in, 5, 0)
	out14 := SliceSafe(in, 6, 4)
	out15 := SliceSafe(in, 6, 7)
	out16 := SliceSafe(in, -10, 1)
	out17 := SliceSafe(in, -1, 3)
	out18 := SliceSafe(in, -10, 7)
	out19 := SliceSafe(in, -10, -7)

	is.Equal([]int{}, out1)
	is.Equal([]int{0}, out2)
	is.Equal([]int{0, 1, 2, 3, 4}, out3)
	is.Equal([]int{0, 1, 2, 3, 4}, out4)
	is.Equal([]int{}, out5)
	is.Equal([]int{1, 2, 3, 4}, out6)
	is.Equal([]int{1, 2, 3, 4}, out7)
	is.Equal([]int{4}, out8)
	is.Equal([]int{}, out9)
	is.Equal([]int{}, out10)
	is.Equal([]int{}, out11)
	is.Equal([]int{}, out12)
	is.Equal([]int{}, out13)
	is.Equal([]int{}, out14)
	is.Equal([]int{}, out15)
	is.Equal([]int{0}, out16)
	is.Equal([]int{0, 1, 2}, out17)
	is.Equal([]int{0, 1, 2, 3, 4}, out18)
	is.Equal([]int{}, out19)
}

func TestReplace(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := ReplaceN(in, 0, 42, 2)
	out2 := ReplaceN(in, 0, 42, 1)
	out3 := ReplaceN(in, 0, 42, 0)
	out4 := ReplaceN(in, 0, 42, -1)
	out5 := ReplaceN(in, 0, 42, -1)
	out6 := ReplaceN(in, -1, 42, 2)
	out7 := ReplaceN(in, -1, 42, 1)
	out8 := ReplaceN(in, -1, 42, 0)
	out9 := ReplaceN(in, -1, 42, -1)
	out10 := ReplaceN(in, -1, 42, -1)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 0}, out1)
	is.Equal([]int{42, 1, 0, 1, 2, 3, 0}, out2)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out3)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out4)
	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out5)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out6)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out7)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out8)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out9)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out10)
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	in := []int{0, 1, 0, 1, 2, 3, 0}

	out1 := ReplaceAll(in, 0, 42)
	out2 := ReplaceAll(in, -1, 42)

	is.Equal([]int{42, 1, 42, 1, 2, 3, 42}, out1)
	is.Equal([]int{0, 1, 0, 1, 2, 3, 0}, out2)
}

func TestCompact(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Compact([]int{2, 0, 4, 0})

	is.Equal(r1, []int{2, 4})

	r2 := Compact([]string{"", "foo", "", "bar", ""})

	is.Equal(r2, []string{"foo", "bar"})

	r3 := Compact([]bool{true, false, true, false})

	is.Equal(r3, []bool{true, true})

	type foo struct {
		bar int
		baz string
	}

	// slice of structs
	// If all fields of an element are zero values, Compact removes it.

	r4 := Compact([]foo{
		{bar: 1, baz: "a"}, // all fields are non-zero values
		{bar: 0, baz: ""},  // all fields are zero values
		{bar: 2, baz: ""},  // bar is non-zero
	})

	is.Equal(r4, []foo{{bar: 1, baz: "a"}, {bar: 2, baz: ""}})

	// slice of pointers to structs
	// If an element is nil, Compact removes it.

	e1, e2, e3 := foo{bar: 1, baz: "a"}, foo{bar: 0, baz: ""}, foo{bar: 2, baz: ""}
	// NOTE: e2 is a zero value of foo, but its pointer &e2 is not a zero value of *foo.
	r5 := Compact([]*foo{&e1, &e2, nil, &e3})

	is.Equal(r5, []*foo{&e1, &e2, &e3})
}

func TestIsSorted(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsSorted([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	is.True(IsSorted([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}))

	is.False(IsSorted([]int{0, 1, 4, 3, 2, 5, 6, 7, 8, 9, 10}))
	is.False(IsSorted([]string{"a", "b", "d", "c", "e", "f", "g", "h", "i", "j"}))
}

func TestIsSortedByKey(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	is.True(IsSortedByKey([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	}))

	is.False(IsSortedByKey([]string{"aa", "b", "ccc"}, func(s string) int {
		return len(s)
	}))

	is.True(IsSortedByKey([]string{"1", "2", "3", "11"}, func(s string) int {
		ret, _ := strconv.Atoi(s)
		return ret
	}))
}

func TestRange(t *testing.T) {
	// 因为目前我没有很好的办法测试泛型，所以暂时只测试int类型
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
		wantErr bool
	}{
		{
			"no error", args{0, 3}, []int{0, 1, 2}, false,
		},
		{"start must be less than end", args{3, 2}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := Range(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Range() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestRangeWithStep(t *testing.T) {
	type args struct {
		start int
		end   int
		step  int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
		wantErr bool
	}{
		{
			"no error", args{0, 3, 1}, []int{0, 1, 2}, false,
		}, {
			"step cannot be zero", args{0, 3, 0}, nil, true,
		},

		{
			"step direction is inconsistent with start and end values", args{0, 3, -1}, nil, true,
		},
		{
			"size equal 0", args{0, 0, 1}, []int{}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := RangeWithStep(tt.args.start, tt.args.end, tt.args.step)
			if (err != nil) != tt.wantErr {
				t.Errorf("RangeWithStep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotRes, tt.wantRes)
			// if !reflect.DeepEqual(gotRes, tt.wantRes) {
			// 	t.Errorf("RangeWithStep() = %v, want %v", gotRes, tt.wantRes)
			// }
		})
	}
}

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
