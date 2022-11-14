package findplayerwithzerooneloss

import (
	"reflect"
	"testing"
)

func Test_findWinners(t *testing.T) {
	type args struct {
		matches [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				matches: [][]int{
					{1, 100000},
				},
			},
			want: [][]int{
				{1},
				{100000},
			},
		},
		{
			name: "2",
			args: args{
				matches: [][]int{
					{1, 3},
					{2, 3},
					{3, 6},
					{5, 6},
					{5, 7},
					{4, 5},
					{4, 8},
					{4, 9},
					{10, 4},
					{10, 9},
				},
			},
			want: [][]int{
				{1, 2, 10},
				{4, 5, 7, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinners(tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWinners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
