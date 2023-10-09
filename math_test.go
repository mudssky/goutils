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
