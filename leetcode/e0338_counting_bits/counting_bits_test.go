package counting_bits

import (
	"reflect"
	"testing"
)

func Test_countingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				0,
			},
			want: []int{0},
		},
		{
			name: "2",
			args: args{
				1,
			},
			want: []int{0, 1},
		},
		{
			name: "3",
			args: args{
				5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countingBitsG(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				0,
			},
			want: []int{0},
		},
		{
			name: "2",
			args: args{
				1,
			},
			want: []int{0, 1},
		},
		{
			name: "3",
			args: args{
				5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingBitsG(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingBitsG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countingBitsS(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				0,
			},
			want: []int{0},
		},
		{
			name: "2",
			args: args{
				1,
			},
			want: []int{0, 1},
		},
		{
			name: "3",
			args: args{
				5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingBitsS(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingBitsS() = %v, want %v", got, tt.want)
			}
		})
	}
}
