package max_min

import "testing"

func Test_maxMin(t *testing.T) {
	type args struct {
		k   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "1",
			args: args{
				k:   3,
				arr: []int32{10, 100, 300, 200, 1000, 20, 30},
			},
			want: 20,
		},
		{
			name: "2",
			args: args{
				k:   4,
				arr: []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				k:   2,
				arr: []int32{2, 1, 2, 1, 2, 1},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				k:   5,
				arr: []int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7802, 3142, 9739, 5629, 5413, 7232},
			},
			want: 1335,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxMin(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("maxMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
