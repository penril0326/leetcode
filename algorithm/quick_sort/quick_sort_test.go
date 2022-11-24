package quick_sort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums:  []int{3, 2, 1, 5, 6, 4},
				left:  0,
				right: 5,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.nums, tt.args.left, tt.args.right)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("quickSort() = %v, want %v", tt.args.nums, tt.want)
			}

		})
	}
}
