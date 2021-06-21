package longest_word

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"monkey", "ale", "apple", "plea"},
			},
			want: "apple",
		},
		{
			name: "2",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"a", "b", "c"},
			},
			want: "a",
		},
		{
			name: "3",
			args: args{
				s:          "abpcplea",
				dictionary: []string{"ale", "apple", "monkey", "plea", "abpcplaaa", "abpcllllll", "abccclllpppeeaaaa"},
			},
			want: "apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
