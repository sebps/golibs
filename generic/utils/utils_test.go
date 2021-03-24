package utils

import (
	"fmt"
	"reflect"
	"testing"
)

// refactor tester :
// "got" should interface{} for either slice, struct, map, pointer, basic structure and forbidden struct ( ex function )
// "want" should be interface{}

type TraverseTester interface {
	Have() interface{}
	Want() interface{}
	Got() interface{}
	Handler() func(input interface{})
}

type TraverseSliceTester struct {
	have struct {
		input                [][]string
		handleBeforeTraverse bool
		handleAfterTraverse  bool
	}
	got  *[]string
	want []string
}

func (t TraverseSliceTester) Have() interface{} {
	return t.have
}
func (t TraverseSliceTester) Want() interface{} {
	return t.want
}
func (t TraverseSliceTester) Got() interface{} {
	return t.got
}
func (t TraverseSliceTester) Handler() func(input interface{}) {
	return func(input interface{}) {
		*t.got = append(*t.got, input.(string))
	}
}

func TestTraverse(t *testing.T) {
	var tests = []interface{}{
		[]TraverseSliceTester{
			{
				have: struct {
					input                [][]string
					handleBeforeTraverse bool
					handleAfterTraverse  bool
				}{
					input:                [][]string{[]string{"one_one", "one_two"}, []string{"two_one", "two_two"}},
					handleBeforeTraverse: false,
					handleAfterTraverse:  false,
				},
				want: []string{"one_one", "one_two", "two_one", "two_two"},
				got:  &[]string{},
			},
		},
	}

	for i, _ := range tests {
		r := reflect.ValueOf(tests[i])
		for j := 0; j < r.Len(); j++ {
			tt := r.Index(j)
			testname := fmt.Sprintf("%s", tt.FieldByName("have").String())
			t.Run(testname, func(t *testing.T) {
				tester := tt.Interface().(TraverseTester)
				have := tester.Have().(struct {
					input                [][]string
					handleBeforeTraverse bool
					handleAfterTraverse  bool
				})
				want := tester.Want()
				got := tester.Got().(*[]string)
				err := Traverse(have.input, tester.Handler(), have.handleBeforeTraverse, have.handleAfterTraverse)
				if err != nil {
					t.Errorf("got error %d, want %d", err, want)
				}
				if !reflect.DeepEqual(*got, want) {
					t.Errorf("got %#v, want %#v", got, want)
				}
			})
		}
	}
}
