package numberofarithmetictriplets

import "testing"

func Test_arithmeticTriplets(t *testing.T) {
	type args struct {
		nums []int
		diff int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, 1, 4, 6, 7, 10},
				diff: 3,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{4, 5, 6, 7, 8, 9},
				diff: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arithmeticTriplets(tt.args.nums, tt.args.diff); got != tt.want {
				t.Errorf("arithmeticTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arithmeticTripletsHash(t *testing.T) {
	type args struct {
		nums []int
		diff int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, 1, 4, 6, 7, 10},
				diff: 3,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums: []int{4, 5, 6, 7, 8, 9},
				diff: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arithmeticTripletsHash(tt.args.nums, tt.args.diff); got != tt.want {
				t.Errorf("arithmeticTripletsHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
