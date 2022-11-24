package digit

import "testing"

func TestSumOfDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 142359,
			},
			want: 24,
		},
		{
			name: "2",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				n: 84759723840,
			},
			want: 57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigit(tt.args.n); got != tt.want {
				t.Errorf("SumOfDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
