package equivalent_tree

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(2), ch)

	for c := range ch {
		t.Log(c)
	}
}

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				t1: tree.New(1),
				t2: tree.New(1),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				t1: tree.New(1),
				t2: tree.New(2),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				t1: tree.New(2),
				t2: tree.New(1),
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				t1: tree.New(2),
				t2: tree.New(2),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}
