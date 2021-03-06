package maps

import (
	"fmt"
	"github.com/sebpsdev/golibs/generic/arrays"
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	var tests = []struct {
		m    map[interface{}]interface{}
		want []interface{}
	}{
		{
			m:    map[interface{}]interface{}{"first_key": "first_value", "second_key": "second_value"},
			want: []interface{}{"first_key", "second_key"},
		},
		{
			m:    map[interface{}]interface{}{1: "first_value", 2: "second_value"},
			want: []interface{}{1, 2},
		},
		{
			m:    map[interface{}]interface{}{true: "first_value", false: "second_value"},
			want: []interface{}{true, false},
		},
		{
			m:    map[interface{}]interface{}{"first_key": 1, "second_key": 2},
			want: []interface{}{"first_key", "second_key"},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.m)
		t.Run(testname, func(t *testing.T) {
			keys := Keys(tt.m)
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
		m    map[interface{}]interface{}
		want []interface{}
	}{
		{
			m:    map[interface{}]interface{}{"first_key": "first_value", "second_key": "second_value"},
			want: []interface{}{"first_value", "second_value"},
		},
		{
			m:    map[interface{}]interface{}{1: "first_value", 2: "second_value"},
			want: []interface{}{"first_value", "second_value"},
		},
		{
			m:    map[interface{}]interface{}{true: "first_value", false: "second_value"},
			want: []interface{}{"first_value", "second_value"},
		},
		{
			m:    map[interface{}]interface{}{"first_key": 1, "second_key": 2},
			want: []interface{}{1, 2},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.m)
		t.Run(testname, func(t *testing.T) {
			values := Values(tt.m)
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
