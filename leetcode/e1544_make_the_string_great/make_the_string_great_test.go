package makethestringgreat

import (
	"testing"
)

func Test_makeGood(t *testing.T) {
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
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "2",
			args: args{
				s: "ss",
			},
			want: "ss",
		},
		{
			name: "3",
			args: args{
				s: "abcdDCBA",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeGoodRecursive(t *testing.T) {
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
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "2",
			args: args{
				s: "ss",
			},
			want: "ss",
		},
		{
			name: "3",
			args: args{
				s: "abcdDCBA",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGoodRecursive(tt.args.s); got != tt.want {
				t.Errorf("makeGoodRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeGoodBruteForce(t *testing.T) {
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
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "2",
			args: args{
				s: "ss",
			},
			want: "ss",
		},
		{
			name: "3",
			args: args{
				s: "abcdDCBA",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGoodBruteForce(tt.args.s); got != tt.want {
				t.Errorf("makeGoodBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
