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
	custom := map[int]struct {
		highPriority int
		lowPriority  int
	}{
		1: struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 1,
			lowPriority:  1,
		},
		2: struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 1,
			lowPriority:  2,
		},
		3: struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 2,
			lowPriority:  1,
		},
		4: struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 2,
			lowPriority:  2,
		},
	}

	var tests = []struct {
		have struct {
			slice interface{}
			less  func(a interface{}, b interface{}) bool
		}
		want interface{}
	}{
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]bool{true, false, false, true}),
				less:  nil,
			},
			want: interface{}([]bool{false, false, true, true}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]string{"delta", "alpha", "beta", "gamma"}),
				less:  nil,
			},
			want: interface{}([]string{"alpha", "beta", "delta", "gamma"}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]int{3, 1, 4, 2}),
				less:  nil,
			},
			want: interface{}([]int{1, 2, 3, 4}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]float32{3.3, 1.2, 4.5, 2.6}),
				less:  nil,
			},
			want: interface{}([]float32{1.2, 2.6, 3.3, 4.5}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]complex64{complex(1, 1), complex(0, 1), complex(2, 2), complex(0, 0)}),
				less:  nil,
			},
			want: interface{}([]complex64{complex(0, 0), complex(0, 1), complex(1, 1), complex(2, 2)}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]complex64{complex(1, 1), complex(0, 1), complex(2, 2), complex(0, 0)}),
				less:  nil,
			},
			want: interface{}([]complex64{complex(0, 0), complex(0, 1), complex(1, 1), complex(2, 2)}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]*int{pointers[3], pointers[1], pointers[4], pointers[2]}),
				less:  nil,
			},
			want: interface{}([]*int{pointers[1], pointers[2], pointers[3], pointers[4]}),
		},
		{
			have: struct {
				slice interface{}
				less  func(a interface{}, b interface{}) bool
			}{
				slice: interface{}([]struct {
					highPriority int
					lowPriority  int
				}{custom[3], custom[1], custom[4], custom[2]}),
				less: func(a interface{}, b interface{}) bool {
					aHighPriority := a.(struct {
						highPriority int
						lowPriority  int
					}).highPriority
					aLowPriority := a.(struct {
						highPriority int
						lowPriority  int
					}).lowPriority
					bHighPriority := b.(struct {
						highPriority int
						lowPriority  int
					}).highPriority
					bLowPriority := a.(struct {
						highPriority int
						lowPriority  int
					}).lowPriority

					if aHighPriority != bHighPriority {
						return aHighPriority < bHighPriority
					} else {
						return aLowPriority < bLowPriority
					}
				},
			},
			want: interface{}([]struct {
				highPriority int
				lowPriority  int
			}{custom[1], custom[2], custom[3], custom[4]}),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.have.slice)
		t.Run(testname, func(t *testing.T) {
			result, err := Sort(tt.have.slice, tt.have.less)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %d, want %d", result, tt.want)
			}
		})
	}
}
