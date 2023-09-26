package minimumoperationstohalvearraysum

import "testing"

func Test_halveArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{5, 19, 8, 1},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 8, 20},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halveArray(tt.args.nums); got != tt.want {
				t.Errorf("halveArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
