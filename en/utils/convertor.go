// Package utils provides English-named type conversion utility functions, covering deep copy and struct conversion.
package utils

import (
	"encoding/json"
)

// DeepClone creates a deep copy of the passed value, suitable for basic types, slices, maps, etc.
// Uses JSON serialization for deep copy, cannot copy unexported fields and functions.
//
// Parameters:
//   - src: source value to copy
//
// Returns:
//   - interface{}: copied value
//   - error: error if copying fails
//
// Example:
//
//	src := map[string]int{"a": 1, "b": 2}
//	dst, err := DeepClone(src)
//	// dst = map[string]int{"a": 1, "b": 2}
func DeepClone(src interface{}) (interface{}, error) {
	if src == nil {
		return nil, nil
	}

	data, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	var dst interface{}
	err = json.Unmarshal(data, &dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

// StructToMap converts a struct to a Map, only converting exported fields (fields starting with uppercase).
//
// Parameters:
//   - struct: struct to convert
//
// Returns:
//   - map[string]interface{}: converted Map
//   - error: error if conversion fails
//
// Example:
//
//	type User struct {
//		Name string
//		Age  int
//	}
//	user := User{Name: "John", Age: 25}
//	m, err := StructToMap(user)
//	// m = map[string]interface{}{"Name": "John", "Age": 25}
func StructToMap(s interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// MapToStruct converts a Map to a specified struct.
//
// Parameters:
//   - m: Map to convert
//   - structPtr: pointer to target struct
//
// Returns:
//   - error: error if conversion fails
//
// Example:
//
//	m := map[string]interface{}{"Name": "John", "Age": 25}
//	type User struct {
//		Name string
//		Age  int
//	}
//	var user User
//	err := MapToStruct(m, &user)
//	// user = User{Name: "John", Age: 25}
func MapToStruct(m map[string]interface{}, structPtr interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, structPtr)
	if err != nil {
		return err
	}

	return nil
}

// PointerToValue returns the value pointed to by the pointer.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - ptr: pointer
//
// Returns:
//   - T: value pointed to by pointer
//
// Example:
//
//	x := 42
//	val := PointerToValue(&x)
//	// val = 42
func PointerToValue[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// ValueToPointer returns a pointer to the passed value.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - value: value to get address of
//
// Returns:
//   - *T: pointer to value
//
// Example:
//
//	ptr := ValueToPointer(42)
//	// ptr = &42
func ValueToPointer[T any](value T) *T {
	return &value
}

// SliceToPointers converts a slice of values to a slice of pointers.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - values: slice of values
//
// Returns:
//   - []*T: slice of pointers
//
// Example:
//
//	values := []int{1, 2, 3}
//	pointers := SliceToPointers(values)
//	// pointers = []*int{&1, &2, &3}
func SliceToPointers[T any](values []T) []*T {
	pointers := make([]*T, len(values))
	for i := range values {
		pointers[i] = &values[i]
	}
	return pointers
}

// PointersToSlice converts a slice of pointers to a slice of values.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - pointers: slice of pointers
//
// Returns:
//   - []T: slice of values
//
// Example:
//
//	pointers := []*int{&1, &2, &3}
//	values := PointersToSlice(pointers)
//	// values = []int{1, 2, 3}
func PointersToSlice[T any](pointers []*T) []T {
	values := make([]T, len(pointers))
	for i := range pointers {
		if pointers[i] != nil {
			values[i] = *pointers[i]
		} else {
			var zero T
			values[i] = zero
		}
	}
	return values
}

// CopyProperties copies fields with the same name from different structs.
//
// Parameters:
//   - source: source struct
//   - destination: target struct (must be struct pointer)
//
// Returns:
//   - error: error if copying fails
//
// Example:
//
//	type A struct {
//		Name string
//		Age  int
//	}
//	type B struct {
//		Name string
//		Age  int
//	}
//	a := A{Name: "John", Age: 25}
//	var b B
//	err := CopyProperties(a, &b)
//	// b = B{Name: "John", Age: 25}
func CopyProperties(source interface{}, destination interface{}) error {
	sourceData, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(sourceData, destination)
	if err != nil {
		return err
	}

	return nil
}

// EncodeByte encodes data to byte slice.
//
// Parameters:
//   - data: data to encode
//
// Returns:
//   - []byte: encoded byte slice
//   - error: error if encoding fails
func EncodeByte(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// DecodeByte decodes byte slice to target object.
//
// Parameters:
//   - data: byte slice to decode
//   - target: target object (must be pointer)
//
// Returns:
//   - error: error if decoding fails
func DecodeByte(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}
