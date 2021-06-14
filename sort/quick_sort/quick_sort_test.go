package quick_sort

import "testing"

func Test_quickSort(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums:  []int{3, 2, 1, 5, 6, 4},
				left:  0,
				right: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, tt.args.left, tt.args.right)
			t.Logf("%v", tt.args.nums)
		})
	}
}
