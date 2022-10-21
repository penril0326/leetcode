package datastructure

import (
	"reflect"
	"testing"
)

func TestMyStack_Top(t *testing.T) {
	tests := []struct {
		name string
		s    myStack
		want interface{}
	}{
		{
			name: "1",
			s: myStack{
				stack: []interface{}{1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyStack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *myStack
		want interface{}
	}{
		{
			name: "1",
			s: &myStack{
				stack: []interface{}{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "2",
			s: &myStack{
				stack: []interface{}{"a", "b", "ab"},
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
