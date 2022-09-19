package backspacestringcompare

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s: "a###",
				t: "####",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s: "####a",
				t: "###b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backspaceCompareWithoutStack(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s: "a###",
				t: "####",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s: "####a",
				t: "###b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompareWithoutStack(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompareWithoutStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
