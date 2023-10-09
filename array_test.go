package goutils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
