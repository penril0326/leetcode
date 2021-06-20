package sort_color

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 0},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 0, 2},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 0, 2, 2, 1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			t.Logf("%v", tt.args.nums)
		})
	}
}
