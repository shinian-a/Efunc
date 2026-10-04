// Package utils provides English-named condition judgment utility functions, covering boolean logic and ternary operations.
package utils

// Bool returns the truthy value of anything.
//
// Parameters:
//   - value: value to evaluate
//
// Returns:
//   - bool: true for truthy values, false for falsy values
//
// Example:
//
//	Bool(true)       // true
//	Bool(1)          // true
//	Bool("")         // false
//	Bool(0)          // false
func Bool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		return v != ""
	case int:
		return v != 0
	case float64:
		return v != 0
	default:
		return v != nil
	}
}

// And returns true if and only if both a and b are true.
//
// Parameters:
//   - a: first boolean value
//   - b: second boolean value
//
// Returns:
//   - bool: result of a && b
//
// Example:
//
//	And(true, true)    // true
//	And(true, false)   // false
func And(a, b bool) bool {
	return a && b
}

// Or returns false if and only if both a and b are false.
//
// Parameters:
//   - a: first boolean value
//   - b: second boolean value
//
// Returns:
//   - bool: result of a || b
//
// Example:
//
//	Or(true, false)   // true
//	Or(false, false)  // false
func Or(a, b bool) bool {
	return a || b
}

// Xor returns false if a and b are the same, true if different.
//
// Parameters:
//   - a: first boolean value
//   - b: second boolean value
//
// Returns:
//   - bool: result of a XOR b
//
// Example:
//
//	Xor(true, true)    // false
//	Xor(true, false)   // true
func Xor(a, b bool) bool {
	return a != b
}

// Xnor returns true if both a and b are true or both are false.
//
// Parameters:
//   - a: first boolean value
//   - b: second boolean value
//
// Returns:
//   - bool: result of a XNOR b
//
// Example:
//
//	Xnor(true, true)    // true
//	Xnor(true, false)   // false
func Xnor(a, b bool) bool {
	return a == b
}

// Nor returns true if and only if both a and b are false.
//
// Parameters:
//   - a: first boolean value
//   - b: second boolean value
//
// Returns:
//   - bool: NOR(a, b) result
//
// Example:
//
//	Nor(false, false)  // true
//	Nor(true, false)   // false
func Nor(a, b bool) bool {
	return !(a || b)
}

// Nand returns false if both a and b are true, otherwise returns true.
//
// Parameters:
//   - a: first boolean value
//   - b: second boolean value
//
// Returns:
//   - bool: NAND(a, b) result
//
// Example:
//
//	Nand(true, true)    // false
//	Nand(true, false)   // true
func Nand(a, b bool) bool {
	return !(a && b)
}

// TernaryOperator generic ternary operator, returns one of two values based on condition.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - condition: condition expression, returns trueVal if true, falseVal if false
//   - trueVal: return value when condition is true
//   - falseVal: return value when condition is false
//
// Returns:
//   - T: selected value
//
// Example:
//
//	TernaryOperator(true, "yes", "no")  // "yes"
//	TernaryOperator(false, 1, 2)        // 2
func TernaryOperator[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
