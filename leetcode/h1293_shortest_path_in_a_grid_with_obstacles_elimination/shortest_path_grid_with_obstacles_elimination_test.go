package shortestpathinagridwithobstacleselimination

import "testing"

func Test_shortestPath(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: [][]int{
					{0, 0, 0},
					{1, 1, 0},
					{0, 0, 0},
					{0, 1, 1},
					{0, 0, 0},
				},
				k: 1,
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				grid: [][]int{
					{0, 1, 1},
					{1, 1, 1},
					{1, 0, 0},
				},
				k: 1,
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				grid: [][]int{
					{0, 0},
					{1, 0},
					{1, 0},
					{1, 0},
					{1, 0},
					{1, 0},
					{0, 0},
					{0, 1},
					{0, 1},
					{0, 1},
					{0, 0},
					{1, 0},
					{1, 0},
					{0, 0},
				},
				k: 4,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.grid, tt.args.k); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
