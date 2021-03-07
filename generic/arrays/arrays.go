package arrays

import (
	"errors"
	"math"
	"reflect"
	"sort"
)

func getKindClass(kind reflect.Kind) string {
	switch kind {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return "number"
	case reflect.Complex64, reflect.Complex128:
		return "complex"
	case reflect.String:
		return "string"
	case reflect.Ptr:
		return "pointer"
	case reflect.Invalid, reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Slice, reflect.Struct, reflect.UnsafePointer:
		return "none"
	default:
		return "none"
	}
}

func compare(kindClass string, i, j int, iValue, jValue interface{}) bool {
	switch kindClass {
	case "number":
		var iCast float64
		var jCast float64

		switch v := iValue.(type) {
		case int:
			iCast = float64(v)
		case int8:
			iCast = float64(v)
		case int16:
			iCast = float64(v)
		case int32:
			iCast = float64(v)
		case int64:
			iCast = float64(v)
		case uint:
			iCast = float64(v)
		case uint8:
			iCast = float64(v)
		case uint16:
			iCast = float64(v)
		case uint32:
			iCast = float64(v)
		case uint64:
			iCast = float64(v)
		case float32:
			iCast = float64(v)
		case float64:
			iCast = v
		}

		switch v := jValue.(type) {
		case int:
			jCast = float64(v)
		case int8:
			jCast = float64(v)
		case int16:
			jCast = float64(v)
		case int32:
			jCast = float64(v)
		case int64:
			jCast = float64(v)
		case uint:
			jCast = float64(v)
		case uint8:
			jCast = float64(v)
		case uint16:
			jCast = float64(v)
		case uint32:
			jCast = float64(v)
		case uint64:
			jCast = float64(v)
		case float32:
			jCast = float64(v)
		case float64:
			jCast = v
		}

		return iCast <= jCast
	case "string":
		var iCast string
		var jCast string

		switch v := iValue.(type) {
		case string:
			iCast = v
		}

		switch v := jValue.(type) {
		case string:
			jCast = v
		}

		return iCast <= jCast
	case "complex":
		var iCast complex128
		var jCast complex128

		switch v := iValue.(type) {
		case complex64:
			iCast = complex128(v)
		case complex128:
			iCast = v
		}

		switch v := jValue.(type) {
		case complex64:
			jCast = complex128(v)
		case complex128:
			jCast = v
		}

		iDist := math.Pow(real(iCast), 2) + math.Pow(imag(iCast), 2)
		jDist := math.Pow(real(jCast), 2) + math.Pow(imag(jCast), 2)

		return iDist <= jDist
	case "boolean":
		var iCast int
		var jCast int

		switch v := iValue.(type) {
		case bool:
			if v == false {
				iCast = 0
			} else {
				iCast = 1
			}
		}

		switch v := jValue.(type) {
		case bool:
			if v == false {
				jCast = 0
			} else {
				jCast = 1
			}
		}

		return iCast <= jCast
	}

	return false
}

type TypedSlice struct {
	value reflect.Value
	kind  reflect.Kind
}

func (t TypedSlice) Len() int {
	return t.value.Len()
}

func (t TypedSlice) Less(i, j int) bool {
	iElem := t.value.Index(i)
	jElem := t.value.Index(j)
	iValue := iElem.Interface()
	jValue := jElem.Interface()

	kindClass := getKindClass(t.kind)
	if kindClass != "pointer" {
		return compare(kindClass, i, j, iValue, jValue)
	} else {
		ptrKind := iElem.Type().Elem().Kind()
		ptrKindClass := getKindClass(ptrKind)
		iPtrValue := iElem.Elem().Interface()
		jPtrValue := jElem.Elem().Interface()

		return compare(ptrKindClass, i, j, iPtrValue, jPtrValue)
	}
}

func (t TypedSlice) Swap(i, j int) {
	swap := reflect.Swapper(t.value.Interface())
	swap(i, j)
}

func Sort(input interface{}) (interface{}, error) {
	s := reflect.ValueOf(input)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}

	eKind := reflect.TypeOf(input).Elem().Kind()

	switch eKind {
	case reflect.Invalid, reflect.Array, reflect.Chan, reflect.Func, reflect.Interface,
		reflect.Map, reflect.Slice, reflect.Struct, reflect.UnsafePointer:
		return nil, errors.New("input collection is not sortable as it has no comparable elements")
	}

	typedSlice := TypedSlice{
		value: s,
		kind:  eKind,
	}

	sort.Sort(typedSlice)

	return typedSlice.value.Interface(), nil
}
