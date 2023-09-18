package snakesandladders

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 35, -1, -1, 13, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 15, -1, -1, -1, -1},
				},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				board: [][]int{
					{-1, -1},
					{-1, 3},
				},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				board: [][]int{
					{1, 1, -1},
					{1, 1, 1},
					{-1, 1, 1},
				},
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				board: [][]int{
					{-1, -1, -1, 46, 47, -1, -1, -1},
					{51, -1, -1, 63, -1, 31, 21, -1},
					{-1, -1, 26, -1, -1, 38, -1, -1},
					{-1, -1, 11, -1, 14, 23, 56, 57},
					{11, -1, -1, -1, 49, 36, -1, 48},
					{-1, -1, -1, 33, 56, -1, 57, 21},
					{-1, -1, -1, -1, -1, -1, 2, -1},
					{-1, -1, -1, 8, 3, -1, 6, 56},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.args.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
