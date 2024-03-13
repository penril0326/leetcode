package nqueensii

import "testing"

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				n: 5,
			},
			want: 10,
		},
		{
			name: "3",
			args: args{
				n: 6,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
