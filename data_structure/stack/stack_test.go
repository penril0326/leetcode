package stack

import (
	"log"
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	stk := NewStack()
	stk.Push(1)
	stk.Push(2)
	assert(stk.Top(), 2)
	assert(stk.Pop(), 2)
	assert(stk.IsEmpty(), false)

	assert(stk.Top(), 1)
	assert(stk.Pop(), 1)
	assert(stk.IsEmpty(), true)
}

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

func assert(src, target interface{}) bool {
	if !reflect.DeepEqual(src, target) {
		log.Fatalf("src = %v, target = %v\n", src, target)
	}

	return true
}
