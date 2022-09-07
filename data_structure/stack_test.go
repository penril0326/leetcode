package stack

import (
	"reflect"
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := NewStack()

	stack.Push(123)
	stack.Push(23)
	stack.Push(4)

	t.Log(stack.Len())
	t.Log(stack.Top())
	t.Log(stack.Pop())

	t.Log(stack.Len())
	t.Log(stack.Top())
	t.Log(stack.Pop())

	t.Log(stack.Len())
	t.Log(stack.Top())
	t.Log(stack.Pop())

}

func TestMyStack_Top(t *testing.T) {
	tests := []struct {
		name string
		s    MyStack
		want interface{}
	}{
		{
			name: "1",
			s:    MyStack{5, 7, 3, 4, 1},
			want: 1,
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
		s    *MyStack
		want interface{}
	}{
		{
			name: "1",
			s:    &MyStack{1, 2, 3, 4, 5},
			want: 5,
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

func TestMyStack_Len(t *testing.T) {
	tests := []struct {
		name string
		s    MyStack
		want int
	}{
		{
			name: "1",
			s:    MyStack{1, 2, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("MyStack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
