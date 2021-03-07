package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func getIntPointer(x int) *int {
	return &x
}

func TestSort(t *testing.T) {
	pointers := map[int]*int{
		1: getIntPointer(1),
		2: getIntPointer(2),
		3: getIntPointer(3),
		4: getIntPointer(4),
	}
	var tests = []struct {
		m    interface{}
		want interface{}
	}{
		{
			m:    interface{}([]bool{true, false, false, true}),
			want: interface{}([]bool{false, false, true, true}),
		},
		{
			m:    interface{}([]string{"delta", "alpha", "beta", "gamma"}),
			want: interface{}([]string{"alpha", "beta", "delta", "gamma"}),
		},
		{
			m:    interface{}([]int{3, 1, 4, 2}),
			want: interface{}([]int{1, 2, 3, 4}),
		},
		{
			m:    interface{}([]float32{3.3, 1.2, 4.5, 2.6}),
			want: interface{}([]float32{1.2, 2.6, 3.3, 4.5}),
		},
		{
			m:    interface{}([]complex64{complex(1, 1), complex(0, 1), complex(2, 2), complex(0, 0)}),
			want: interface{}([]complex64{complex(0, 0), complex(0, 1), complex(1, 1), complex(2, 2)}),
		},
		{
			m:    interface{}([]*int{pointers[3], pointers[1], pointers[4], pointers[2]}),
			want: interface{}([]*int{pointers[1], pointers[2], pointers[3], pointers[4]}),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.m)
		t.Run(testname, func(t *testing.T) {
			result, err := Sort(tt.m)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}
