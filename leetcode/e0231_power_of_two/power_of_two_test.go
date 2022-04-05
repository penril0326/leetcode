package poweroftwo

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				n: -1,
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				n: -256,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				n: 64,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
