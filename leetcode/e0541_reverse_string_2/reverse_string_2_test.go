package reversestring2

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyz",
				k: 3,
			},
			want: "cbadefihgjklonmpqrutsvwxzy",
		},
		{
			name: "2",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
