// Package utils 提供中文命名的比较工具函数库，涵盖数值、字符串等类型的比较操作。
package utils

import (
	"reflect"
	"golang.org/x/exp/constraints"
)

// CompareEqual 检查两个值是否完全相等（包括类型和值）。
//
// Type Parameters:
//   - T: 任意可比较类型
//
// Parameters:
//   - a: 第一个值
//   - b: 第二个值
//
// Returns:
//   - bool: 相等返回 true，否则返回 false
//
// Example:
//
//	CompareEqual(1, 1)        // true
//	CompareEqual("a", "a")    // true
//	CompareEqual(1, "1")      // false
func CompareEqual[T comparable](a, b T) bool {
	return a == b
}

// CompareValueEqual 检查两个值的值是否相等（忽略类型差异）。
// 使用 reflect.DeepEqual 进行深度比较。
//
// Parameters:
//   - a: 第一个值
//   - b: 第二个值
//
// Returns:
//   - bool: 值相等返回 true，否则返回 false
//
// Example:
//
//	CompareValueEqual([]int{1,2,3}, []int{1,2,3})  // true
//	CompareValueEqual(map[int]int{1:1}, map[int]int{1:1})  // true
func CompareValueEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// LessThan 检查 left 是否小于 right。
//
// Type Parameters:
//   - T: 支持比较运算的类型 (numeric types)
//
// Parameters:
//   - left: 左值
//   - right: 右值
//
// Returns:
//   - bool: left < right 返回 true
//
// Example:
//
//	LessThan(1, 2)    // true
//	LessThan(3, 3)    // false
func LessThan[T constraints.Ordered](left, right T) bool {
	return left < right
}

// GreaterThan 检查 left 是否大于 right。
//
// Type Parameters:
//   - T: 支持比较运算的类型 (numeric types)
//
// Parameters:
//   - left: 左值
//   - right: 右值
//
// Returns:
//   - bool: left > right 返回 true
//
// Example:
//
//	GreaterThan(3, 2)    // true
//	GreaterThan(2, 2)    // false
func GreaterThan[T constraints.Ordered](left, right T) bool {
	return left > right
}

// LessOrEqual 检查 left 是否小于或等于 right。
//
// Type Parameters:
//   - T: 支持比较运算的类型 (numeric types)
//
// Parameters:
//   - left: 左值
//   - right: 右值
//
// Returns:
//   - bool: left <= right 返回 true
//
// Example:
//
//	LessOrEqual(1, 2)    // true
//	LessOrEqual(2, 2)    // true
//	LessOrEqual(3, 2)    // false
func LessOrEqual[T constraints.Ordered](left, right T) bool {
	return left <= right
}

// GreaterOrEqual 检查 left 是否大于或等于 right。
//
// Type Parameters:
//   - T: 支持比较运算的类型 (numeric types)
//
// Parameters:
//   - left: 左值
//   - right: 右值
//
// Returns:
//   - bool: left >= right 返回 true
//
// Example:
//
//	GreaterOrEqual(3, 2)    // true
//	GreaterOrEqual(2, 2)    // true
//	GreaterOrEqual(1, 2)    // false
func GreaterOrEqual[T constraints.Ordered](left, right T) bool {
	return left >= right
}

// FloatApproxEqual 检查两个浮点数是否在指定误差范围内相等。
//
// Parameters:
//   - a: 第一个浮点数
//   - b: 第二个浮点数
//   - delta: 允许的误差范围
//
// Returns:
//   - bool: 在误差范围内相等返回 true
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
