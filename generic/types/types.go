package types

import (
	"errors"
	"reflect"
)

func GeneralizeMap(input interface{}) (map[interface{}]interface{}, error) {
	r := reflect.ValueOf(input)
	if r.Kind() != reflect.Map {
		return nil, errors.New("input is not a map")
	}

	m := make(map[interface{}]interface{})

	for _, k := range r.MapKeys() {
		v := r.MapIndex(k)
		underlyingK := k.Interface()
		underlyingV := v.Interface()
		m[underlyingK] = underlyingV
	}

	return m, nil
}

func GeneralizeSlice(input interface{}) ([]interface{}, error) {
	r := reflect.ValueOf(input)
	if r.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}

	s := make([]interface{}, r.Len())

	for i := 0; i < r.Len(); i++ {
		v := r.Index(i)
		underlyingV := v.Interface()
		s[i] = underlyingV
	}

	return s, nil
}
