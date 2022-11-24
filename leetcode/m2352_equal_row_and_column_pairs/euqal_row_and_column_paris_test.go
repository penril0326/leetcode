package equalrowandcolumnpairs

import "testing"

func Test_equalPairs(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				grid: [][]int{
					{11, 1},
					{1, 11},
				},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				grid: [][]int{
					{3, 2, 1},
					{1, 7, 6},
					{2, 7, 7},
				},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				grid: [][]int{
					{3, 1, 2, 2},
					{1, 4, 4, 5},
					{2, 4, 2, 2},
					{2, 4, 2, 2},
				},
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				grid: [][]int{
					{3, 1, 2, 2},
					{1, 4, 4, 4},
					{2, 4, 2, 2},
					{2, 4, 2, 2},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalPairs(tt.args.grid); got != tt.want {
				t.Errorf("equalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertArrayToString(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 2, 3},
			},
			want: "1,2,3,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertArrayToString(tt.args.arr); got != tt.want {
				t.Errorf("convertArrayToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
