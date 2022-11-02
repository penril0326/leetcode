package reversestring

import (
	"reflect"
	"testing"
)

// func Test_reverseString(t *testing.T) {
// 	type args struct {
// 		s []byte
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			before := copy(tt.args.s)
// 			reverseString(tt.args.s)
// 			t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
// 		})
// 	}
// }

// func assert(s []byte, want []byte) bool {

// }

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "2",
			args: args{
				s: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %s, want %s", tt.args.s, tt.want)
			}
		})
	}
}
