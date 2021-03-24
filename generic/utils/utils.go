package utils

import (
	"reflect"
)

func isTraversable(kind reflect.Kind) bool {
	switch kind {
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer, reflect.Struct, reflect.Array:
		return true
	default:
		return false
	}
}

func Traverse(input interface{}, handler func(i interface{}), handleBeforeTraverse bool, handleAfterTraverse bool) error {
	var err error
	r := reflect.ValueOf(input)

	if !isTraversable(r.Kind()) {
		handler(input)
		return nil
	}

	if handleBeforeTraverse {
		handler(input)
	}

	switch r.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < r.Len(); i++ {
			v := r.Index(i)
			underlyingV := v.Interface()
			err = Traverse(underlyingV, handler, handleBeforeTraverse, handleAfterTraverse)
		}
	case reflect.Struct:
		for i := 0; i < r.NumField(); i++ {
			v := r.Field(i)
			underlyingV := v.Interface()
			err = Traverse(underlyingV, handler, handleBeforeTraverse, handleAfterTraverse)
		}
	case reflect.Map:
		for _, k := range r.MapKeys() {
			v := r.MapIndex(k)
			underlyingK := k.Interface()
			err = Traverse(underlyingK, handler, handleBeforeTraverse, handleAfterTraverse)

			underlyingV := v.Interface()
			err = Traverse(underlyingV, handler, handleBeforeTraverse, handleAfterTraverse)
		}
	case reflect.UnsafePointer, reflect.Ptr, reflect.Interface:
		v := r.Elem()
		underlyingV := v.Interface()
		err = Traverse(underlyingV, handler, handleBeforeTraverse, handleAfterTraverse)
	}

	if handleAfterTraverse {
		handler(input)
	}

	return err
}
