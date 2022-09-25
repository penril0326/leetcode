package finalpriceswithaspecialdiscountinashop

import (
	"reflect"
	"testing"
)

func Test_finalPricesBruteForce(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				prices: []int{2, 3, 4, 5, 1},
			},
			want: []int{1, 2, 3, 4, 1},
		},
		{
			name: "2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "3",
			args: args{
				prices: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 1, 1, 1, 1},
		},
		{
			name: "4",
			args: args{
				prices: []int{8, 4, 6, 2, 3},
			},
			want: []int{4, 2, 4, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPricesBruteForce(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPricesBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				prices: []int{8, 4, 6, 2, 3},
			},
			want: []int{4, 2, 4, 2, 3},
		},
		{
			name: "2",
			args: args{
				prices: []int{2, 3, 4, 5, 1},
			},
			want: []int{1, 2, 3, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPrices(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
