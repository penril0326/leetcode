package detonatethemaximumbombs

import "testing"

func Test_maximumDetonation(t *testing.T) {
	type args struct {
		bombs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				bombs: [][]int{
					{2, 1, 3},
					{6, 1, 4},
				},
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				bombs: [][]int{
					{1, 1, 5},
					{10, 10, 5},
				},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				bombs: [][]int{
					{1, 2, 3},
					{2, 3, 1},
					{3, 4, 2},
					{4, 5, 3},
					{5, 6, 4},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDetonation(tt.args.bombs); got != tt.want {
				t.Errorf("maximumDetonation() = %v, want %v", got, tt.want)
			}
		})
	}
}
