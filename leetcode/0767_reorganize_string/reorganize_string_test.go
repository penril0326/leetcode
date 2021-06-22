package reorganize_string

import "testing"

func Test_reorganizeString(t *testing.T) {
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
				s: "abbabbaaab",
			},
			want: "ababababab",
		},
		{
			name: "2",
			args: args{
				s: "aab",
			},
			want: "aba",
		},
		{
			name: "3",
			args: args{
				s: "aaab",
			},
			want: "",
		},
		{
			name: "4",
			args: args{
				s: "ogccckcwmbmxtsbmozli",
			},
			want: "cocgctcwmsmimxbkbloz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.s); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
