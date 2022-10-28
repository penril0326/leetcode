package longestcontinuoussubarray

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums  []int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:  []int{8, 2, 4, 7},
				limit: 4,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums:  []int{4, 2, 2, 2, 4, 4, 2, 2},
				limit: 0,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				nums:  []int{2, 2, 2, 4, 4, 2, 5, 5, 5, 5, 5, 2},
				limit: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
