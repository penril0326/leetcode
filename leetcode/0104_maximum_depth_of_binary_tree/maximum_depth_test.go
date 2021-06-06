package maximum_depth

import (
	"practice/leetcode"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: &leetcode.TreeNode{
					Val: 3,
					Left: &leetcode.TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val: 20,
						Left: &leetcode.TreeNode{
							Val:   15,
							Left:  nil,
							Right: nil,
						},
						Right: &leetcode.TreeNode{
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
				root: &leetcode.TreeNode{
					Val:  1,
					Left: nil,
					Right: &leetcode.TreeNode{
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
				root: &leetcode.TreeNode{
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
