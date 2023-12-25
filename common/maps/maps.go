package maps

// group
func GroupBy(maps []map[string]interface{}, key string) map[string][]map[string]interface{} {
	groups := make(map[string][]map[string]interface{})
	for _, m := range maps {
		k := m[key].(string) // XXX: will panic if m[key] is not a string.
		groups[k] = append(groups[k], m)
	}
	return groups
}

// filter and output
func Filter(maps map[string]interface{}, predicate func(K string) bool) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	for k, v := range maps {
		if predicate(k) {
			m[k] = v
		}
	}
	return m
}

// filter and output arrays
func FilterToArrays(maps map[string]interface{}, predicate func(K string) bool) []interface{} {
	arrays := make([]interface{}, 0)
	for k, v := range maps {
		if predicate(k) {
			arrays = append(arrays, v)
		}
	}
	return arrays
}
