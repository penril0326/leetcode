package longestsubstringwithoutrepeatingchars

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
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
				s: " ",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringOptimized(t *testing.T) {
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
				s: " ",
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				s: "abba",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringOptimized(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringOptimized() = %v, want %v", got, tt.want)
			}
		})
	}
}
