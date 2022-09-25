package removealladjacentduplicatesinstring

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},
		{
			name: "2",
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
		{
			name: "3",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "4",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "5",
			args: args{
				s: "bbb",
			},
			want: "b",
		},
		{
			name: "6",
			args: args{
				s: "bbbb",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
