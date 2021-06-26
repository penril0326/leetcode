package missing_numbers

import (
	"reflect"
	"testing"
)

func Test_missingNumbers(t *testing.T) {
	type args struct {
		arr []int32
		brr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test case 0",
			args: args{
				arr: []int32{203, 204, 205, 206, 207, 208, 203, 204, 205, 206},
				brr: []int32{203, 204, 204, 205, 206, 207, 205, 208, 203, 206, 205, 206, 204},
			},
			want: []int32{204, 205, 206},
		},
		{
			name: "Test case 4",
			args: args{
				arr: []int32{11, 4, 11, 7, 13, 4, 12, 11, 10, 14},
				brr: []int32{11, 4, 11, 7, 3, 7, 10, 13, 4, 8, 12, 11, 10, 14, 12},
			},
			want: []int32{3, 7, 8, 10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumbers(tt.args.arr, tt.args.brr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
