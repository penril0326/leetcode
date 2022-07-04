package maximal_square

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]byte{
					{'0'},
				},
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				matrix: [][]byte{
					{'0', '1'},
					{'1', '0'},
				},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				matrix: [][]byte{
					{'0', '1', '1', '1', '0'},
					{'1', '1', '1', '1', '1'},
					{'0', '1', '1', '1', '1'},
					{'0', '1', '1', '1', '1'},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximalSquareDP(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				matrix: [][]byte{
					{'0'},
				},
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				matrix: [][]byte{
					{'0', '1'},
					{'1', '0'},
				},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				matrix: [][]byte{
					{'1', '0', '1', '0', '0'},
					{'1', '0', '1', '1', '1'},
					{'1', '1', '1', '1', '1'},
					{'1', '0', '0', '1', '0'},
				},
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				matrix: [][]byte{
					{'0', '1', '1', '1', '0'},
					{'1', '1', '1', '1', '1'},
					{'0', '1', '1', '1', '1'},
					{'0', '1', '1', '1', '1'},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquareDP(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquareDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
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
		{
			name: "1",
			args: args{
				a: 1,
				b: 2,
				c: 3,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				a: 1,
				b: 3,
				c: 2,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				a: 2,
				b: 1,
				c: 3,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				a: 2,
				b: 3,
				c: 1,
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				a: 3,
				b: 1,
				c: 2,
			},
			want: 1,
		},
		{
			name: "6",
			args: args{
				a: 3,
				b: 2,
				c: 1,
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				a: 1,
				b: 1,
				c: 3,
			},
			want: 1,
		},
		{
			name: "8",
			args: args{
				a: 1,
				b: 2,
				c: 1,
			},
			want: 1,
		},
		{
			name: "9",
			args: args{
				a: 2,
				b: 1,
				c: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
