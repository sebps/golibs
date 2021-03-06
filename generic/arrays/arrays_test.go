package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	var tests = []struct {
		m    []interface{}
		want []interface{}
	}{
		{
			m:    []interface{}{true, false, false, true},
			want: []interface{}{false, false, true, true},
		},
		{
			m:    []interface{}{"delta", "alpha", "beta", "gamma"},
			want: []interface{}{"alpha", "beta", "delta", "gamma"},
		},
		{
			m:    []interface{}{3, 1, 4, 2},
			want: []interface{}{1, 2, 3, 4},
		},
		{
			m:    []interface{}{3.3, 1.2, 4.5, 2.6},
			want: []interface{}{1.2, 2.6, 3.3, 4.5},
		},
		{
			m:    []interface{}{complex(1, 1), complex(0, 1), complex(2, 2), complex(0, 0)},
			want: []interface{}{complex(0, 0), complex(0, 1), complex(1, 1), complex(2, 2)},
		},
		{
			m:    []interface{}{complex(1, 1), complex(0, 1), complex(2, 2), complex(0, 0), complex(-2, 2)},
			want: []interface{}{complex(0, 0), complex(0, 1), complex(1, 1), complex(2, 2), complex(-2, 2)},
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
