package maps

import (
	"errors"
	"reflect"
)

func KeysValues(input interface{}) (interface{}, interface{}, error) {
	m := reflect.ValueOf(input)
	if m.Kind() != reflect.Map {
		return nil, nil, errors.New("input is not a map")
	}

	mTyp := reflect.TypeOf(input)
	kTyp := mTyp.Key()
	vTyp := mTyp.Elem()

	ks := reflect.MakeSlice(reflect.SliceOf(kTyp), 0, m.Len())
	vs := reflect.MakeSlice(reflect.SliceOf(vTyp), 0, m.Len())

	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		ks = reflect.Append(ks, k)
		vs = reflect.Append(vs, v)
	}

	return ks.Interface(), vs.Interface(), nil
}

func Keys(input interface{}) (interface{}, error) {
	keys, _, err := KeysValues(input)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func Values(input interface{}) (interface{}, error) {
	_, values, err := KeysValues(input)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func FindKey(key interface{}, input interface{}) (bool, error) {
	m := reflect.ValueOf(input)
	if m.Kind() != reflect.Map {
		return false, errors.New("input is not a map")
	}

	for _, k := range m.MapKeys() {
		if reflect.DeepEqual(k.Interface(), key) {
			return true, nil
		}
	}

	return false, nil
}
