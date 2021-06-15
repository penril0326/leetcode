package merge_intervals

import (
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
		},
		{
			name: "3",
			args: args{
				intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			},
		},
		{
			name: "4",
			args: args{
				intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			},
		},
		{
			name: "5",
			args: args{
				intervals: [][]int{{1, 4}, {0, 2}, {3, 5}},
			},
		},
		{
			name: "6",
			args: args{
				intervals: [][]int{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.args.intervals)
			t.Logf("%v", got)
		})
	}
}
