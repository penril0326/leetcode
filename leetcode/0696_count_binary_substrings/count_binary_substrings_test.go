package count_binary_substrings

import (
	"testing"
)

func Test_countBinarySubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "00110011",
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				s: "10101",
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				s: "00110",
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				s: "1010001",
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				s: "00100",
			},
			want: 2,
		},
		{
			name: "6",
			args: args{
				s: "0011100",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
