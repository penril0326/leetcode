package kclosestpointstoorigin

import (
	"strconv"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				points: [][]int{
					{1, 3}, {-2, 2},
				},
				k: 1,
			},
			want: [][]int{
				{-2, 2},
			},
		},
		{
			name: "2",
			args: args{
				points: [][]int{
					{3, 3}, {5, -1}, {-2, 4},
				},
				k: 2,
			},
			want: [][]int{
				{-2, 4}, {3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.args.points, tt.args.k)
			resultMap := gotMap(got)
			for _, w := range tt.want {
				x, y := w[0], w[1]
				if _, ok := resultMap[strconv.Itoa(x)+strconv.Itoa(y)]; !ok {
					t.Errorf("kClosest() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func gotMap(ss [][]int) map[string]struct{} {
	h := map[string]struct{}{}
	for _, s := range ss {
		x, y := s[0], s[1]
		h[strconv.Itoa(x)+strconv.Itoa(y)] = struct{}{}
	}

	return h
}
