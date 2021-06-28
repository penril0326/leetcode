package sherlock_array

import "testing"

func Test_balancedSums(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 0-1",
			args: args{
				arr: []int32{1, 2, 3},
			},
			want: "NO",
		},
		{
			name: "Test case 0-2",
			args: args{
				arr: []int32{1, 2, 3, 3},
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedSums(tt.args.arr); got != tt.want {
				t.Errorf("balancedSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
