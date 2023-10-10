package goutils

import (
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"+1", args{num: 1}, 1},
		{"0", args{num: 0}, 0},
		{"-1", args{num: -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Abs(tt.args.num); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Abs() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMaxN(t *testing.T) {
	type args struct {
		first  int
		others []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"first bigger than others", args{10, []int{1, 2, 3}}, 10},
		{"first smaller than others", args{2, []int{5, 7, 9}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := MaxN(tt.args.first, tt.args.others...); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("MaxN() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"num1 > num2", args{1, 2}, 2},
		{"num1 = num2", args{2, 2}, 2},
		{"num1 < num2", args{1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Max(tt.args.num1, tt.args.num2); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Max() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
