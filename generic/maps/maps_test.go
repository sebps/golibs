package maps

import (
	"fmt"
	"github.com/sebps/golibs/generic/arrays"
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

			sortedKeys, err := arrays.Sort(keys, nil)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			sortedWant, err := arrays.Sort(tt.want, nil)
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

			sortedValues, err := arrays.Sort(values, nil)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}

			sortedWant, err := arrays.Sort(tt.want, nil)
			if err != nil {
				t.Errorf("got error %d, want %d", err, tt.want)
			}
			if !reflect.DeepEqual(sortedValues, sortedWant) {
				t.Errorf("got %d, want %d", values, tt.want)
			}
		})
	}
}

func TestFindKey(t *testing.T) {
	var tests = []struct {
		key  interface{}
		m    interface{}
		want struct {
			found bool
			err   error
		}
	}{
		{
			key: "first_key",
			m:   interface{}(map[string]string{"first_key": "first_value", "second_key": "second_value"}),
			want: struct {
				found bool
				err   error
			}{
				found: true,
				err:   nil,
			},
		},
		{
			key: "third_key",
			m:   interface{}(map[string]string{"first_key": "first_value", "second_key": "second_value"}),
			want: struct {
				found bool
				err   error
			}{
				found: false,
				err:   nil,
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.m)
		t.Run(testname, func(t *testing.T) {
			found, err := FindKey(tt.key, tt.m)
			if err != nil {
				t.Errorf("got error %#v, want %#v", err, tt.want)
			}
			result := struct {
				found bool
				err   error
			}{
				found: found,
				err:   err,
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %#v, want %#v", result, tt.want)
			}
		})
	}
}
