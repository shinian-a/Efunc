// Package utils provides English-named Map operation utility functions, covering key-value queries and conversions.
package utils

// GetKeys retrieves all keys from a Map.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - m: Map to operate on
//
// Returns:
//   - []K: slice of all keys
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	GetKeys(m)  // []string{"a", "b"}
func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetValues retrieves all values from a Map.
//
// Type Parameters:
//   - K: key type
//   - V: value type
//
// Parameters:
//   - m: Map to operate on
//
// Returns:
//   - []V: slice of all values
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	GetValues(m)  // []int{1, 2}
func GetValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// ContainsKey checks if a Map contains a specified key.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - m: Map to check
//   - key: key to search for
//
// Returns:
//   - bool: true if contains, false otherwise
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	ContainsKey(m, "a")  // true
//	ContainsKey(m, "c")  // false
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// SafeGet safely retrieves a value from Map, returns default value if key doesn't exist.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - m: Map to query
//   - key: key to search for
//   - defaultVal: default value to return if key doesn't exist
//
// Returns:
//   - V: value from Map or default value
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	SafeGet(m, "a", 0)  // 1
//	SafeGet(m, "c", 0)  // 0
func SafeGet[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if val, ok := m[key]; ok {
		return val
	}
	return defaultVal
}

// MergeMaps merges multiple Maps, later Maps will override same keys from earlier Maps.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - maps: list of Maps to merge
//
// Returns:
//   - map[K]V: merged Map
//
// Example:
//
//	m1 := map[string]int{"a": 1}
//	m2 := map[string]int{"b": 2}
//	MergeMaps(m1, m2)  // map[string]int{"a": 1, "b": 2}
func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// MapToPairs converts a Map to a 2D slice [[key, value], ...].
//
// Type Parameters:
//   - K: key type
//   - V: value type
//
// Parameters:
//   - m: Map to convert
//
// Returns:
//   - [][]interface{}: 2D slice
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	MapToPairs(m)  // [][]interface{}{{"a", 1}, {"b", 2}}
func MapToPairs[K comparable, V any](m map[K]V) [][]interface{} {
	result := make([][]interface{}, 0, len(m))
	for k, v := range m {
		result = append(result, []interface{}{k, v})
	}
	return result
}

// PairsToMap converts a 2D slice to a Map.
//
// Parameters:
//   - pairs: 2D slice [[key, value], ...]
//
// Returns:
//   - map[interface{}]interface{}: converted Map
//
// Example:
//
//	pairs := [][]interface{}{{"a", 1}, {"b", 2}}
//	PairsToMap(pairs)  // map[interface{}]interface{}{"a": 1, "b": 2}
func PairsToMap(pairs [][]interface{}) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	for _, pair := range pairs {
		if len(pair) >= 2 {
			result[pair[0]] = pair[1]
		}
	}
	return result
}

// FilterMap filters Map key-value pairs based on condition function.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - m: Map to filter
//   - condition: filter condition function, returns true to keep key-value pair
//
// Returns:
//   - map[K]V: filtered Map
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	FilterMap(m, func(k string, v int) bool { return v > 1 })
//	// map[string]int{"b": 2, "c": 3}
func FilterMap[K comparable, V any](m map[K]V, condition func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if condition(k, v) {
			result[k] = v
		}
	}
	return result
}

// MapValues applies transformation function to all values in Map.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: source value type
//   - U: transformed value type
//
// Parameters:
//   - m: Map to transform
//   - fn: transformation function
//
// Returns:
//   - map[K]U: transformed Map
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	MapValues(m, func(v int) string { return strconv.Itoa(v) })
//	// map[string]string{"a": "1", "b": "2"}
func MapValues[K comparable, V any, U any](m map[K]V, fn func(V) U) map[K]U {
	result := make(map[K]U, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// MapLength gets the length of a Map (number of key-value pairs).
//
// Type Parameters:
//   - K: key type
//   - V: value type
//
// Parameters:
//   - m: Map to query
//
// Returns:
//   - int: Map length
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	MapLength(m)  // 2
func MapLength[K comparable, V any](m map[K]V) int {
	return len(m)
}

// ClearMap clears all items from a Map.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - m: Map to clear
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	ClearMap(m)
//	// m = map[string]int{}
func ClearMap[K comparable, V any](m map[K]V) {
	for k := range m {
		delete(m, k)
	}
}

// ReverseMap creates a reversed Map where original values become keys and original keys become values.
//
// Type Parameters:
//   - K: original key type
//   - V: original value type (must be comparable as it becomes key)
//
// Parameters:
//   - m: Map to reverse
//
// Returns:
//   - map[V]K: reversed Map
//
// Note:
//   If original Map has duplicate values, later keys will override earlier ones.
//
// Example:
//
//	m := map[string]int{"a": 1, "b": 2}
//	ReverseMap(m)  // map[int]string{1: "a", 2: "b"}
func ReverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// ExtractKeys extracts all unique keys from multiple Maps.
//
// Type Parameters:
//   - K: key type (must be comparable)
//   - V: value type
//
// Parameters:
//   - maps: list of Maps to process
//
// Returns:
//   - map[K]struct{}: set of keys (represented with empty struct)
//
// Example:
//
//	m1 := map[string]int{"a": 1}
//	m2 := map[string]int{"b": 2}
//	ExtractKeys(m1, m2)  // map[string]struct{}{"a": {}, "b": {}}
func ExtractKeys[K comparable, V any](maps ...map[K]V) map[K]struct{} {
	result := make(map[K]struct{})
	for _, m := range maps {
		for k := range m {
			result[k] = struct{}{}
		}
	}
	return result
}

// ExtractValues extracts all unique values from multiple Maps.
//
// Type Parameters:
//   - K: key type
//   - V: value type (must be comparable)
//
// Parameters:
//   - maps: list of Maps to process
//
// Returns:
//   - map[V]struct{}: set of values
//
// Example:
//
//	m1 := map[string]int{"a": 1}
//	m2 := map[string]int{"b": 2}
//	ExtractValues(m1, m2)  // map[int]struct{}{1: {}, 2: {}}
func ExtractValues[K comparable, V comparable](maps ...map[K]V) map[V]struct{} {
	result := make(map[V]struct{})
	for _, m := range maps {
		for _, v := range m {
			result[v] = struct{}{}
		}
	}
	return result
}
