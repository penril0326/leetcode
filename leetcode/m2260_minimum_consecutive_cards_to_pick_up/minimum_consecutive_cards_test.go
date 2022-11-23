package minimumconsecutivecardstopickup

import (
	"testing"
)

func Test_minimumCardPickup(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 7},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				cards: []int{0, 0},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCardPickup(tt.args.cards); got != tt.want {
				t.Errorf("minimumCardPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumCardPickup2(t *testing.T) {
	type args struct {
		cards []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 7},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				cards: []int{0, 0},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				cards: []int{3, 4, 2, 3, 4, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCardPickup2(tt.args.cards); got != tt.want {
				t.Errorf("minimumCardPickup2() = %v, want %v", got, tt.want)
			}
		})
	}
}
