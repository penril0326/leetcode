package howmanyapplescanyouputintothebucket

import "testing"

func Test_maxNumberOfApples(t *testing.T) {
	type args struct {
		weight []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				weight: []int{100, 200, 150, 1000},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				weight: []int{900, 950, 800, 1000, 700, 800},
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				weight: []int{1000, 1000, 1000, 1000, 1000},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfApples(tt.args.weight); got != tt.want {
				t.Errorf("maxNumberOfApples() = %v, want %v", got, tt.want)
			}
		})
	}
}
