package successfulpairsofspellsandpotiions

import (
	"reflect"
	"testing"
)

func Test_successfulPairs(t *testing.T) {
	type args struct {
		spells  []int
		potions []int
		success int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				spells:  []int{5, 1, 3},
				potions: []int{1, 2, 3, 4, 5},
				success: 7,
			},
			want: []int{4, 0, 3},
		},
		{
			name: "2",
			args: args{
				spells:  []int{3, 1, 2},
				potions: []int{8, 5, 8},
				success: 16,
			},
			want: []int{2, 0, 2},
		},
		{
			name: "3",
			args: args{
				spells:  []int{1, 2, 3, 4, 5, 6, 7},
				potions: []int{1, 2, 3, 4, 5, 6, 7},
				success: 25,
			},
			want: []int{0, 0, 0, 1, 3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := successfulPairs(tt.args.spells, tt.args.potions, tt.args.success); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("successfulPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
