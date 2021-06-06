package same_tree

import (
	"practice/leetcode"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *leetcode.TreeNode
		q *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				p: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				q: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				p: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
				},
				q: &leetcode.TreeNode{
					Val: 1,
					Left: &leetcode.TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: &leetcode.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
