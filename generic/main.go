package main

import (
	"fmt"
	"github.com/sebps/golibs/generic/arrays"
	"github.com/sebps/golibs/generic/maps"
	"github.com/sebps/golibs/generic/types"
)

func main() {
	// Arrays
	sInt := []int{1, 4, 2, 7}
	arrays.Sort(sInt, nil)
	fmt.Println(sInt)
	// expect 1,2,4,7
	sStr := []string{"delta", "alpha", "beta", "gamma"}
	arrays.Sort(sStr, nil)
	fmt.Println(sStr)
	// expect "alpha","beta","delta","gamma"

	customSlice := []struct {
		highPriority int
		lowPriority  int
	}{
		struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 2,
			lowPriority:  1,
		},
		struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 2,
			lowPriority:  2,
		},
		struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 1,
			lowPriority:  2,
		},
		struct {
			highPriority int
			lowPriority  int
		}{
			highPriority: 1,
			lowPriority:  1,
		},
	}

	arrays.Sort(customSlice, customLess)
	fmt.Println(customSlice)
	// expect
	// {
	// 	highPriority: 1,
	// 	lowPriority:  1,
	// },
	// {
	// 	highPriority: 1,
	// 	lowPriority:  2,
	// },
	// {
	// 	highPriority: 2,
	// 	lowPriority:  1,
	// },
	// {
	// 	highPriority: 2,
	// 	lowPriority:  2,
	// }

	// Maps
	mIntStr := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	keys, _ := maps.Keys(mIntStr)
	fmt.Println(keys)
	// expect 1,2,3,4
	values, _ := maps.Values(mIntStr)
	fmt.Println(values)
	// expect "one","two","three","four" ( order can change )

	// Types
	s := []int{1, 4, 2, 7}
	genericS, _ := types.GeneralizeSlice(s)
	genericSType := fmt.Sprintf("%T", genericS)
	fmt.Println(genericSType)
	// expect []interface{}

	m := map[int]int{1: 1, 2: 2}
	genericM, _ := types.GeneralizeMap(m)
	genericMType := fmt.Sprintf("%T", genericM)
	fmt.Println(genericMType)
	// expect map[interface{}]interface{}
}

func customLess(a interface{}, b interface{}) bool {
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
	bLowPriority := b.(struct {
		highPriority int
		lowPriority  int
	}).lowPriority

	if aHighPriority != bHighPriority {
		return aHighPriority < bHighPriority
	} else {
		return aLowPriority < bLowPriority
	}
}
