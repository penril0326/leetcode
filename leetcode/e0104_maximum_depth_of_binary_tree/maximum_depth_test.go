package maximum_depth

import (
	"practice/data_structure/node"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *node.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: &node.TreeNode{
					Val: 3,
					Left: &node.TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &node.TreeNode{
						Val: 20,
						Left: &node.TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &node.TreeNode{
							Val:   7,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				root: &node.TreeNode{
					Val:  1,
					Left: nil,
					Right: &node.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				root: &node.TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
