package findthedifference

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "1",
			args: args{
				s: "abc",
				t: "cbad",
			},
			want: 'd',
		},
		{
			name: "2",
			args: args{
				s: "",
				t: "t",
			},
			want: 't',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
