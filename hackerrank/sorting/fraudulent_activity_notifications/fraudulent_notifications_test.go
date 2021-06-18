package fraudulent_notification

import "testing"

func Test_activityNotifications(t *testing.T) {
	type args struct {
		expenditure []int32
		d           int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "1",
			args: args{
				expenditure: []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
				d:           5,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				expenditure: []int32{1, 2, 3, 4, 4},
				d:           4,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				expenditure: []int32{3, 3, 3, 3},
				d:           2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := activityNotifications(tt.args.expenditure, tt.args.d); got != tt.want {
				t.Errorf("activityNotifications() = %v, want %v", got, tt.want)
			}
		})
	}
}
