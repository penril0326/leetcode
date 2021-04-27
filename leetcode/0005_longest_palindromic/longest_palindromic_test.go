package longest_palindromic

import "testing"

func Test_longestPalindrome(t *testing.T) {
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
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
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
				s: "ac",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLongestEvenPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestEvenPalindrome(tt.args.s); got != tt.want {
				t.Errorf("getLongestEvenPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLongestOddPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLongestOddPalindrome(tt.args.s); got != tt.want {
				t.Errorf("getLongestOddPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
