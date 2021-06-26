package cut_the_tree

import "testing"

func Test_cutTheTree(t *testing.T) {
	type args struct {
		data  []int32
		edges [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test case 0",
			args: args{
				data: []int32{100, 200, 100, 500, 100, 600},
				edges: [][]int32{
					{1, 2},
					{2, 3},
					{2, 5},
					{4, 5},
					{5, 6},
				},
			},
			want: 400,
		},
		{
			name: "Test case 1",
			args: args{
				data: []int32{205, 573, 985, 242, 830, 514, 592, 263, 142, 915},
				edges: [][]int32{
					{2, 8},
					{10, 5},
					{1, 7},
					{6, 9},
					{4, 3},
					{8, 10},
					{5, 1},
					{7, 6},
					{9, 4},
				},
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutTheTree(tt.args.data, tt.args.edges); got != tt.want {
				t.Errorf("cutTheTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
