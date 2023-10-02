package top_k_freq

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want map[int]bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: map[int]bool{
				1: true,
				2: true,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 0, 1, 0},
				k:    1,
			},
			want: map[int]bool{
				0: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.args.nums, tt.args.k)
			answer := map[int]bool{}
			for _, v := range got {
				answer[v] = true
			}

			if !reflect.DeepEqual(answer, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", answer, tt.want)
			}
		})
	}
}
