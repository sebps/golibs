package types

import (
	"fmt"
	"reflect"
	"testing"
)

type Tester interface {
	Have() interface{}
	Want() interface{}
}

// Maps
type StringStringMapTester struct {
	have map[string]string
	want map[interface{}]interface{}
}

func (t StringStringMapTester) Have() interface{} {
	return t.have
}
func (t StringStringMapTester) Want() interface{} {
	return t.want
}

type IntIntMapTester struct {
	have map[int]int
	want map[interface{}]interface{}
}

func (t IntIntMapTester) Have() interface{} {
	return t.have
}
func (t IntIntMapTester) Want() interface{} {
	return t.want
}

type IntStringMapTester struct {
	have map[int]string
	want map[interface{}]interface{}
}

func (t IntStringMapTester) Have() interface{} {
	return t.have
}
func (t IntStringMapTester) Want() interface{} {
	return t.want
}

func TestGeneralizeMap(t *testing.T) {
	var tests = []interface{}{
		[]StringStringMapTester{
			{
				have: map[string]string{"first_key": "first_value", "second_key": "second_value"},
				want: map[interface{}]interface{}{"first_key": "first_value", "second_key": "second_value"},
			},
		},
		[]IntIntMapTester{
			{
				have: map[int]int{1: 1, 2: 2},
				want: map[interface{}]interface{}{1: 1, 2: 2},
			},
		},
		[]IntStringMapTester{
			{
				have: map[int]string{1: "first_value", 2: "second_value"},
				want: map[interface{}]interface{}{1: "first_value", 2: "second_value"},
			},
		},
	}

	for i, _ := range tests {
		r := reflect.ValueOf(tests[i])
		for j := 0; j < r.Len(); j++ {
			tt := r.Index(j)
			testname := fmt.Sprintf("%s", tt.FieldByName("have").String())
			t.Run(testname, func(t *testing.T) {
				tester := tt.Interface().(Tester)
				have := tester.Have()
				want := tester.Want()
				got, err := GeneralizeMap(have)

				if err != nil {
					t.Errorf("got error %d, want %d", err, want)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %d, want %d", got, want)
				}
			})
		}
	}
}

// Slices
type StringSliceTester struct {
	have []string
	want []interface{}
}

func (t StringSliceTester) Have() interface{} {
	return t.have
}
func (t StringSliceTester) Want() interface{} {
	return t.want
}

type IntSliceTester struct {
	have []int
	want []interface{}
}

func (t IntSliceTester) Have() interface{} {
	return t.have
}
func (t IntSliceTester) Want() interface{} {
	return t.want
}

func TestGeneralizeSlice(t *testing.T) {
	var tests = []interface{}{
		[]StringSliceTester{
			{
				have: []string{"first_value", "second_value"},
				want: []interface{}{"first_value", "second_value"},
			},
		},
		[]IntSliceTester{
			{
				have: []int{1, 2},
				want: []interface{}{1, 2},
			},
		},
	}

	for i, _ := range tests {
		r := reflect.ValueOf(tests[i])
		for j := 0; j < r.Len(); j++ {
			tt := r.Index(j)
			testname := fmt.Sprintf("%s", tt.FieldByName("have").String())
			t.Run(testname, func(t *testing.T) {
				tester := tt.Interface().(Tester)
				have := tester.Have()
				want := tester.Want()
				got, err := GeneralizeSlice(have)

				if err != nil {
					t.Errorf("got error %d, want %d", err, want)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %d, want %d", got, want)
				}
			})
		}
	}
}
