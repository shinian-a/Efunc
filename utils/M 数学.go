// Package utils 提供中文命名的数学工具函数库，涵盖数值计算、统计等操作。
package utils

import "math"

// MathSum calculates the sum of an integer slice.
//
// Parameters:
//   - nums: integer slice to sum
//
// Returns:
//   - int: sum of all elements
//
// Example:
//
//	MathSum([]int{1, 2, 3, 4, 5})  // 15
func MathSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// MathAverage calculates the average of an integer slice.
//
// Parameters:
//   - nums: integer slice to calculate average
//
// Returns:
//   - float64: average value; 0 if empty slice
//
// Example:
//
//	MathAverage([]int{1, 2, 3, 4, 5})  // 3.0
func MathAverage(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	return float64(MathSum(nums)) / float64(len(nums))
}

// MathMax finds the maximum value in an integer slice.
//
// Parameters:
//   - nums: integer slice to find maximum
//
// Returns:
//   - int: maximum value; MinInt if empty
//
// Example:
//
//	MathMax([]int{1, 5, 3, 9, 2})  // 9
func MathMax(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// MathMin finds the minimum value in an integer slice.
//
// Parameters:
//   - nums: integer slice to find minimum
//
// Returns:
//   - int: minimum value; MaxInt if empty
//
// Example:
//
//	MathMin([]int{1, 5, 3, 9, 2})  // 1
func MathMin(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt
	}
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

// Factorial calculates the factorial of a given integer.
//
// Parameters:
//   - n: integer to calculate factorial (non-negative)
//
// Returns:
//   - int64: n!; 0 if n < 0
//
// Example:
//
//	Factorial(5)  // 120
func Factorial(n int) int64 {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

// Fibonacci calculates the nth Fibonacci number.
//
// Parameters:
//   - n: position in Fibonacci sequence (starting from 0)
//
// Returns:
//   - int64: nth Fibonacci number
//
// Example:
//
//	Fibonacci(0)  // 0
//	Fibonacci(1)  // 1
//	Fibonacci(10) // 55
func Fibonacci(n int) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var a, b int64 = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// IsPrime checks if an integer is a prime number.
//
// Parameters:
//   - n: integer to check
//
// Returns:
//   - bool: true if prime, false otherwise
//
// Example:
//
//	IsPrime(7)   // true
//	IsPrime(10)  // false
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Abs returns the absolute value of an integer.
//
// Parameters:
//   - n: integer to get absolute value
//
// Returns:
//   - int: absolute value
//
// Example:
//
//	Abs(-5)  // 5
//	Abs(5)   // 5
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Pow calculates base raised to the power of exponent.
//
// Parameters:
//   - base: base number
//   - exponent: exponent
//
// Returns:
//   - float64: base^exponent
//
// Example:
//
//	Pow(2, 3)  // 8.0
//	Pow(5, 0)  // 1.0
func Pow(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Sqrt calculates the square root of a given number.
//
// Parameters:
//   - n: number to calculate square root
//
// Returns:
//   - float64: square root of n
//
// Example:
//
//	Sqrt(16)  // 4.0
//	Sqrt(2)   // 1.414...
func Sqrt(n float64) float64 {
	return math.Sqrt(n)
}

// PI returns the value of pi (π).
//
// Returns:
//   - float64: π ≈ 3.141592653589793
func PI() float64 {
	return math.Pi
}

// E returns the value of Euler's number (e).
//
// Returns:
//   - float64: e ≈ 2.718281828459045
func E() float64 {
	return math.E
}
