package slicex

// Column Return the values from a single column in the input slice
//  Returns a slice or map
//  slice: if the parameter indexKey is not specified
//         containing the columns in the input slice whose key is columnKey
//  map: if the parameter indexKey is specified
//       the key in the map is the column corresponding to indexKey
//       and the value is the column corresponding to columnKey
func Column(input []map[interface{}]interface{}, columnKey interface{}, indexKey ...interface{}) interface{} {
	// if the parameter indexKey is specified
	if len(indexKey) > 0 {
		result := make(map[interface{}]interface{})
		for _, s := range input {
			if v, ok := s[columnKey]; ok {
				if i, ok := s[indexKey[0]]; ok {
					result[i] = v
				}
			}
		}
		return result
	}

	// if the parameter indexKey is not specified
	result := make([]interface{}, 0, len(input))
	for _, s := range input {
		if v, ok := s[columnKey]; ok {
			result = append(result, v)
		}
	}
	return result
}

// In check if a value exists in a slice
//  true if needle is found in the haystack, false otherwise
func In(needle interface{}, haystack []interface{}) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// Diff Calculate the difference of slices
//  Returns a slice, containing all the values in slice1 but not in other slices
func Diff(slice1 []interface{}, sliceOthers ...[]interface{}) []interface{} {
	hash := make(map[interface{}]bool)
	result := make([]interface{}, 0)

	for i := 0; i < len(slice1); i++ {
		hash[slice1[i]] = false
	}

	for i := 0; i < len(sliceOthers); i++ {
		for j := 0; j < len(sliceOthers[i]); j++ {
			if _, exist := hash[sliceOthers[i][j]]; exist {
				hash[sliceOthers[i][j]] = true
			}
		}
	}

	for k, v := range hash {
		if !v {
			result = append(result, k)
		}
	}
	return result
}
