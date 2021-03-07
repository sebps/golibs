# Generic
Some methods for datastructures of elements of type interface{} ( passed by empty interface )

## Examples

```go
package main

import (
  "github.com/sebpsdev/golibs/generic/arrays"
  "github.com/sebpsdev/golibs/generic/maps"
  "github.com/sebpsdev/golibs/generic/types"
  "fmt"
)

func main() {
  // Arrays
  s := []interface{}{1,4,2,7}
  s := arrays.Sort(m)
  fmt.Println(s)
  // expect 1,2,4,7
  s = []interface{}{"delta","alpha","beta","gamma"}
  s := arrays.Sort(m)
  fmt.Println(s)
  // expect "alpha","beta","delta","gamma"

  // Maps
  s := []interface{}{1,4,2,7}
  s := arrays.Sort(m)
  fmt.Println(s)
  // expect 1,2,4,7
  m = map[interface{}]interface{}{
    1: "one",
    2: "two",
    3: "three",
    4: "four",
  }
  var keys []interface{}
  keys = maps.Keys(m)
  fmt.Println(keys)
  // expect 1,2,3,4
  var values []interface{}
  values = maps.Values(m)
  fmt.Println(values)
  // expect "one","two","three","four"

  // Types
  s := []interface{}{1,4,2,7}
  var genericS []interface{} 
  genericS = types.GeneralizeSlice(s)
  genericSType := fmt.Sprintf("%T", genericS)
  fmt.Println(genericSType)
  // expect []interface{}
    
  m := map[interface{}]interface{}{1:1,2:2}
  var genericM []interface{} 
  genericM = types.GeneralizeMap(m)
  genericMType := fmt.Sprintf("%T", genericM)
  fmt.Println(genericMType)
  // expect map[interface{}]interface{}
}

```