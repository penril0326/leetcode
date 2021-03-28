package addWithoutOperator

import "testing"

func TestAddWithoutOperator(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: 5,
				b: 7,
			},
			want: 12,
		},
		{
			name: "12",
			args: args{
				a: 85,
				b: 67,
			},
			want: 152,
		},
		{
			name: "3",
			args: args{
				a: 0,
				b: 7,
			},
			want: 7,
		},
		{
			name: "4",
			args: args{
				a: 5,
				b: 0,
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				a: -100,
				b: 7,
			},
			want: -93,
		},
		{
			name: "6",
			args: args{
				a: -123,
				b: -321,
			},
			want: -444,
		},
		{
			name: "7",
			args: args{
				a: 567432113,
				b: 789431124,
			},
			want: 1356863237,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddWithoutOperator(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AddWithoutOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
