package maxconsecutiveonesiii

import (
	"testing"
)

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3,
			},
			want: 10,
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 0, 0, 1},
				k:    4,
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				nums: []int{0, 0, 0, 0},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestOnesOptimized(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
				k:    2,
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				k:    3,
			},
			want: 10,
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 0, 0, 1},
				k:    4,
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				nums: []int{0, 0, 0, 0},
				k:    2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnesOptimized(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnesOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}
