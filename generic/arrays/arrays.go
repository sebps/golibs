package arrays

import (
	"errors"
	"math"
	"reflect"
	"sort"
)

func Sort(slice []interface{}) ([]interface{}, error) {
	var typ string
	for _, element := range slice {
		currentTyp := reflect.TypeOf(element)
		if typ == "" {
			typ = currentTyp.Name()
		} else if typ != currentTyp.Name() {
			return nil, errors.New("all input array elements must be of the same type")
		}
	}

	sorted := make([]interface{}, len(slice))

	switch typ {
	case "bool":
		s := make([]int, len(slice))
		for i, v := range slice {
			var bitSetVar int
			if v.(bool) {
				bitSetVar = 1
			}
			s[i] = bitSetVar
		}
		sort.Ints(s)
		for i, v := range s {
			var bitSetVar bool
			if v == 0 {
				bitSetVar = false
			} else {
				bitSetVar = true
			}
			sorted[i] = bitSetVar
		}
	case "int", "int8", "int16", "int32", "int64", "uint ", "uint8 ", "uint16 ", "uint32 ", "uint64 ", "uintptr", "byte", "rune":
		s := make([]int, len(slice))
		for i, v := range slice {
			s[i] = v.(int)
		}
		sort.Ints(s)
		for i, v := range s {
			sorted[i] = v
		}
	case "string":
		s := make([]string, len(slice))
		for i, v := range slice {
			s[i] = v.(string)
		}
		sort.Strings(s)
		for i, v := range s {
			sorted[i] = v
		}
	case "float32 ", "float64":
		s := make([]float64, len(slice))
		for i, v := range slice {
			s[i] = v.(float64)
		}
		sort.Float64s(s)
		for i, v := range s {
			sorted[i] = v
		}
	case "complex64 ", "complex128":
		m := make(map[float64][]complex128)
		var s []float64
		for _, v := range slice {
			d := math.Pow(real(v.(complex128)), 2) + math.Pow(imag(v.(complex128)), 2)
			if _, ok := m[d]; !ok {
				s = append(s, d)
			}
			m[d] = append(m[d], v.(complex128))
		}
		sort.Float64s(s)
		for i, d := range s {
			for j, c := range m[d] {
				sorted[i+j] = c
			}
		}
	}

	return sorted, nil
}

// func Reverse(slice []interface{}) ([]interface{}, error) {
// 	var typ string
// 	for _, element := range slice {
// 		currentTyp := reflect.TypeOf(element)
// 		if typ == "" {
// 			typ = currentTyp.Name()
// 		} else if typ != currentTyp.Name() {
// 			return nil, errors.New("all input array elements must be of the same type")
// 		}
// 	}

// 	sorted := make([]interface{}, len(slice))

// 	switch typ {
// 	case "bool":
// 		s := make([]int, len(slice))
// 		for i, v := range slice {
// 			var bitSetVar int
// 			if v.(bool) {
// 				bitSetVar = 1
// 			}
// 			s[i] = bitSetVar
// 		}
// 		sort.Ints(s)
// 		for i, v := range s {
// 			var bitSetVar bool
// 			if v == 0 {
// 				bitSetVar = false
// 			} else {
// 				bitSetVar = true
// 			}
// 			sorted[i] = bitSetVar
// 		}
// 	case "int", "int8", "int16", "int32", "int64", "uint ", "uint8 ", "uint16 ", "uint32 ", "uint64 ", "uintptr", "byte", "rune":
// 		s := make([]int, len(slice))
// 		for i, v := range slice {
// 			s[i] = v.(int)
// 		}
// 		sort.Ints(s)
// 		for i, v := range s {
// 			sorted[i] = v
// 		}
// 	case "string":
// 		s := make([]string, len(slice))
// 		for i, v := range slice {
// 			s[i] = v.(string)
// 		}
// 		sort.Strings(s)
// 		for i, v := range s {
// 			sorted[i] = v
// 		}
// 	case "float32 ", "float64":
// 		s := make([]float64, len(slice))
// 		for i, v := range slice {
// 			s[i] = v.(float64)
// 		}
// 		sort.Float64s(s)
// 		for i, v := range s {
// 			sorted[i] = v
// 		}
// 	case "complex64 ", "complex128":
// 		s := make([]float64, len(slice))
// 		for i, v := range slice {
// 			s[i] = math.Pow(real(v.(complex128)), 2) + math.Pow(imag(v.(complex128)), 2)
// 		}
// 		sort.Float64s(s)
// 		for i, v := range s {
// 			sorted[i] = v
// 		}
// 	}

// 	return sorted, nil
// }
