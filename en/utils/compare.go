// Package utils provides English-named comparison utility functions, covering numerical and string comparisons.
package utils

import (
	"reflect"
	"golang.org/x/exp/constraints"
)

// Equal checks if two values are exactly equal (including type and value).
//
// Type Parameters:
//   - T: any comparable type
//
// Parameters:
//   - a: first value
//   - b: second value
//
// Returns:
//   - bool: true if equal, false otherwise
//
// Example:
//
//	Equal(1, 1)        // true
//	Equal("a", "a")    // true
//	Equal(1, "1")      // false
func Equal[T comparable](a, b T) bool {
	return a == b
}

// EqualValue checks if two values have the same value (ignoring type differences).
// Uses reflect.DeepEqual for deep comparison.
//
// Parameters:
//   - a: first value
//   - b: second value
//
// Returns:
//   - bool: true if values are equal, false otherwise
//
// Example:
//
//	EqualValue([]int{1,2,3}, []int{1,2,3})  // true
//	EqualValue(map[int]int{1:1}, map[int]int{1:1})  // true
func EqualValue(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// LessThan checks if left is less than right.
//
// Type Parameters:
//   - T: comparable type (numeric types)
//
// Parameters:
//   - left: left value
//   - right: right value
//
// Returns:
//   - bool: true if left < right
//
// Example:
//
//	LessThan(1, 2)    // true
//	LessThan(3, 3)    // false
func LessThan[T constraints.Ordered](left, right T) bool {
	return left < right
}

// GreaterThan checks if left is greater than right.
//
// Type Parameters:
//   - T: comparable type (numeric types)
//
// Parameters:
//   - left: left value
//   - right: right value
//
// Returns:
//   - bool: true if left > right
//
// Example:
//
//	GreaterThan(3, 2)    // true
//	GreaterThan(2, 2)    // false
func GreaterThan[T constraints.Ordered](left, right T) bool {
	return left > right
}

// LessOrEqual checks if left is less than or equal to right.
//
// Type Parameters:
//   - T: comparable type (numeric types)
//
// Parameters:
//   - left: left value
//   - right: right value
//
// Returns:
//   - bool: true if left <= right
//
// Example:
//
//	LessOrEqual(1, 2)    // true
//	LessOrEqual(2, 2)    // true
//	LessOrEqual(3, 2)    // false
func LessOrEqual[T constraints.Ordered](left, right T) bool {
	return left <= right
}

// GreaterOrEqual checks if left is greater than or equal to right.
//
// Type Parameters:
//   - T: comparable type (numeric types)
//
// Parameters:
//   - left: left value
//   - right: right value
//
// Returns:
//   - bool: true if left >= right
//
// Example:
//
//	GreaterOrEqual(3, 2)    // true
//	GreaterOrEqual(2, 2)    // true
//	GreaterOrEqual(1, 2)    // false
func GreaterOrEqual[T constraints.Ordered](left, right T) bool {
	return left >= right
}

// FloatApproxEqual checks if two floats are approximately equal within a delta.
//
// Parameters:
//   - a: first float
//   - b: second float
//   - delta: allowed error margin
//
// Returns:
//   - bool: true if within error margin
//
// Example:
//
//	FloatApproxEqual(3.14159, 3.14160, 0.0001)  // true
//	FloatApproxEqual(3.14159, 3.14160, 0.00001) // false
func FloatApproxEqual(a, b, delta float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= delta
}
