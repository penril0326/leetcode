package coin_change

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				coins:  []int{1},
				amount: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
