package heap

import (
	"reflect"
	"testing"
)

func Test_mHeap_Len(t *testing.T) {
	type fields struct {
		h   []interface{}
		cmp func(i, j interface{}) bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				h:   []interface{}{3, 5, 6, 9},
				cmp: nil,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := mHeap{
				h:   tt.fields.h,
				cmp: tt.fields.cmp,
			}
			if got := h.Len(); got != tt.want {
				t.Errorf("mHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mHeap_Top(t *testing.T) {
	type fields struct {
		h   []interface{}
		cmp func(i, j interface{}) bool
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "1",
			fields: fields{
				h:   []interface{}{"123", "sss", "ajskd"},
				cmp: nil,
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := mHeap{
				h:   tt.fields.h,
				cmp: tt.fields.cmp,
			}
			if got := h.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mHeap.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mHeap_Push(t *testing.T) {
	// Min heap
	m := mHeap{
		h: nil,
		cmp: func(i, j interface{}) bool {
			return i.(int) <= j.(int)
		},
	}

	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields *mHeap
		args   args
		want   []interface{}
	}{
		{
			name:   "1",
			fields: &m,
			args: args{
				v: 4,
			},
			want: []interface{}{4},
		},
		{
			name:   "2",
			fields: &m,
			args: args{
				v: 3,
			},
			want: []interface{}{3, 4},
		},
		{
			name:   "3",
			fields: &m,
			args: args{
				v: 1,
			},
			want: []interface{}{1, 4, 3},
		},
		{
			name:   "4",
			fields: &m,
			args: args{
				v: 8,
			},
			want: []interface{}{1, 4, 3, 8},
		},
		{
			name:   "5",
			fields: &m,
			args: args{
				v: 5,
			},
			want: []interface{}{1, 4, 3, 8, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.Push(tt.args.v)
			for idx, val := range tt.fields.h {
				if !reflect.DeepEqual(val, tt.want[idx]) {
					t.Errorf("mHeap.Push() = %v, want %v", tt.fields.h, tt.want)
				}
			}
		})
	}
}

func Test_mHeap_Pop(t *testing.T) {
	m := NewHeap([]interface{}{4, 5, 7, 1, 2, 3}, func(i, j interface{}) bool {
		return i.(int) >= j.(int)
	})
	tests := []struct {
		name   string
		fields MyHeap
		want   interface{}
	}{
		{
			name:   "1",
			fields: m,
			want:   7,
		},
		{
			name:   "2",
			fields: m,
			want:   5,
		},
		{
			name:   "3",
			fields: m,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
