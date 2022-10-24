package datastructure

import (
	"reflect"
	"testing"
)

func TestNewQue(t *testing.T) {
	que := NewQue()
	que.Enqueue(2)
	que.Enqueue(3)
	que.Enqueue(4)
	t.Log("length = ", que.Length())
	t.Log("front = ", que.Front())

	t.Log(que.Dequeue())
	t.Log(que.Dequeue())
	t.Log(que.Dequeue())
	t.Log(que.Dequeue())
}

func Test_myQueue_Front(t *testing.T) {
	type fields struct {
		head   *node
		tail   *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "1",
			fields: fields{
				head: &node{
					val: 1,
					next: &node{
						val: 2,
					},
				},
				length: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := myQueue{
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				length: tt.fields.length,
			}
			if got := q.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myQueue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myQueue_IsEmpty(t *testing.T) {
	type fields struct {
		head   *node
		tail   *node
		length int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "1",
			fields: fields{
				head: &node{
					val: 1,
					next: &node{
						val: 2,
					},
				},
				length: 2,
			},
			want: false,
		},
		{
			name: "2",
			fields: fields{
				head: &node{
					val: 1,
				},
				length: 1,
			},
			want: false,
		},
		{
			name: "3",
			fields: fields{
				length: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := myQueue{
				head:   tt.fields.head,
				tail:   tt.fields.tail,
				length: tt.fields.length,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("myQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
