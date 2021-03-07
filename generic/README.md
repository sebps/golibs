# Generic
Some untyped generic methods for datastructures.
No specific type is expected as argument for the methods. 
Errors can however be returned in case of inappropriate underlying type for the called behavior ( e.g calling "array.sort" method on a slice of Struct elements would result in an error as Struct is not a comparable type )

## Examples

```go
package main

import (
	"fmt"
	"github.com/sebpsdev/golibs/generic/arrays"
	"github.com/sebpsdev/golibs/generic/maps"
	"github.com/sebpsdev/golibs/generic/types"
)

func main() {
	// Arrays
	sInt := []int{1, 4, 2, 7}
	arrays.Sort(sInt)
	fmt.Println(sInt)
	// expect 1,2,4,7
	sStr := []string{"delta", "alpha", "beta", "gamma"}
	arrays.Sort(sStr)
	fmt.Println(sStr)
	// expect "alpha","beta","delta","gamma"

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
```