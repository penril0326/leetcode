package minimumspeedtoarriceontime

import "testing"

func Test_minSpeedOnTime(t *testing.T) {
	type args struct {
		dist []int
		hour float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				dist: []int{1, 3, 2},
				hour: 6,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				dist: []int{1, 3, 2},
				hour: 2.7,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				dist: []int{1, 3, 2},
				hour: 1.9,
			},
			want: -1,
		},
		{
			name: "4",
			args: args{
				dist: []int{1, 1, 100000},
				hour: 2.01,
			},
			want: 10000000,
		},
		{
			name: "5",
			args: args{
				dist: []int{1, 1},
				hour: 1.0,
			},
			want: -1,
		},
		{
			name: "6",
			args: args{
				dist: []int{1, 3, 2},
				hour: 2.7,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
