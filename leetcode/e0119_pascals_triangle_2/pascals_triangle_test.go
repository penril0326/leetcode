package pascals_triangle2

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "2",
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
		{
			name: "3",
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
