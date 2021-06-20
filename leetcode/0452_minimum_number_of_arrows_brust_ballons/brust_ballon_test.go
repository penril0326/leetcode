package brust_ballon

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				points: [][]int{
					{10, 16}, {2, 8}, {1, 6}, {7, 12}, {4, 7},
				},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				points: [][]int{
					{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8},
				},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				points: [][]int{
					{1, 9}, {7, 16}, {2, 5}, {7, 12}, {9, 11}, {2, 10}, {9, 16}, {3, 9}, {1, 3},
				},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				points: [][]int{
					{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
