package maps

func Keys(m map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values(m map[interface{}]interface{}) []interface{} {
	values := make([]interface{}, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}
