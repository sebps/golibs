package maps

import (
	"fmt"
	"github.com/sebpsdev/golibs/generic/arrays"
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	var tests = []struct {
		m    interface{}
		want interface{}
	}{
		{
			m:    interface{}(map[string]string{"first_key": "first_value", "second_key": "second_value"}),
			want: interface{}([]string{"first_key", "second_key"}),
		},
		{
			m:    interface{}(map[int]string{1: "first_value", 2: "second_value"}),
			want: interface{}([]int{1, 2}),
		},
		{
			m:    interface{}(map[bool]string{true: "first_value", false: "second_value"}),
			want: interface{}([]bool{true, false}),
		},
		{
			m:    interface{}(map[string]int{"first_key": 1, "second_key": 2}),
			want: interface{}([]string{"first_key", "second_key"}),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.m)
		t.Run(testname, func(t *testing.T) {
			keys, err := Keys(interface{}(tt.m))
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			sortedKeys, err := arrays.Sort(keys)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			sortedWant, err := arrays.Sort(tt.want)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			if !reflect.DeepEqual(sortedKeys, sortedWant) {
				t.Errorf("got %d, want %d", keys, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	var tests = []struct {
		m    interface{}
		want interface{}
	}{
		{
			m:    interface{}(map[string]string{"first_key": "first_value", "second_key": "second_value"}),
			want: interface{}([]string{"first_value", "second_value"}),
		},
		{
			m:    interface{}(map[int]string{1: "first_value", 2: "second_value"}),
			want: interface{}([]string{"first_value", "second_value"}),
		},
		{
			m:    interface{}(map[bool]string{true: "first_value", false: "second_value"}),
			want: interface{}([]string{"first_value", "second_value"}),
		},
		{
			m:    interface{}(map[string]int{"first_key": 1, "second_key": 2}),
			want: interface{}([]int{1, 2}),
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.m)
		t.Run(testname, func(t *testing.T) {
			values, err := Values(tt.m)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			sortedValues, err := arrays.Sort(values)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			sortedWant, err := arrays.Sort(tt.want)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}
			if !reflect.DeepEqual(sortedValues, sortedWant) {
				t.Errorf("got %d, want %d", values, tt.want)
			}
		})
	}
}
