package evaluatedivision

import (
	"reflect"
	"testing"
)

func Test_calcEquationBFS(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "1",
			args: args{
				equations: [][]string{
					{"a", "b"},
					{"b", "c"},
				},
				values: []float64{2.0, 3.0},
				queries: [][]string{
					{"a", "c"},
					{"b", "a"},
					{"a", "e"},
					{"a", "a"},
					{"x", "x"},
				},
			},
			want: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name: "2",
			args: args{
				equations: [][]string{
					{"a", "b"},
					{"b", "c"},
					{"bc", "cd"},
				},
				values: []float64{1.5, 2.5, 5.0},
				queries: [][]string{
					{"a", "c"},
					{"c", "b"},
					{"bc", "cd"},
					{"cd", "bc"},
				},
			},
			want: []float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			name: "3",
			args: args{
				equations: [][]string{
					{"a", "b"},
				},
				values: []float64{0.5},
				queries: [][]string{
					{"a", "b"},
					{"b", "a"},
					{"a", "c"},
					{"x", "y"},
				},
			},
			want: []float64{0.5, 2.0, -1.0, -1.0},
		},
		{
			name: "4",
			args: args{
				equations: [][]string{
					{"x1", "x2"},
					{"x2", "x3"},
					{"x3", "x4"},
					{"x4", "x5"},
				},
				values: []float64{3.0, 4.0, 5.0, 6.0},
				queries: [][]string{
					{"x1", "x5"},
					{"x5", "x2"},
					{"x2", "x4"},
					{"x2", "x2"},
					{"x2", "x9"},
					{"x9", "x9"},
				},
			},
			want: []float64{360.0, 1.0 / 120.0, 20.0, 1.0, -1.0, -1.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquationBFS(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquationBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
