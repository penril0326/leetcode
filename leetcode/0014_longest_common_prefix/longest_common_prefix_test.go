package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				strs: []string{"leets", "leetcode", "leet", "leeds"},
			},
			want: "lee",
		},
		{
			name: "2",
			args: args{
				strs: []string{"dog", "dog", "dog"},
			},
			want: "dog",
		},
		{
			name: "3",
			args: args{
				strs: []string{"dog", "dog", "doggy"},
			},
			want: "dog",
		},
		{
			name: "4",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "5",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "6",
			args: args{
				strs: []string{},
			},
			want: "",
		},
		{
			name: "7",
			args: args{
				strs: []string{"racecar"},
			},
			want: "racecar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
