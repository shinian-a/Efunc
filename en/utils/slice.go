// Package utils provides English-named slice operation utility functions, covering deduplication, sorting, searching and other operations.
package utils

import "sort"

// Deduplicate removes duplicate elements from an integer slice, maintaining original order.
//
// Parameters:
//   - nums: integer slice to deduplicate
//
// Returns:
//   - []int: deduplicated integer slice
//
// Example:
//
//	Deduplicate([]int{1, 2, 2, 3, 4, 4, 5})  // []int{1, 2, 3, 4, 5}
func Deduplicate(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

// GenericDeduplicate generic version of deduplication function for any comparable type.
//
// Type Parameters:
//   - T: comparable type
//
// Parameters:
//   - slice: slice to deduplicate
//
// Returns:
//   - []T: deduplicated slice
//
// Example:
//
//	strings := []string{"a", "b", "a", "c"}
//	GenericDeduplicate(strings)  // []string{"a", "b", "c"}
func GenericDeduplicate[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}

	seen := make(map[T]bool)
	result := make([]T, 0)

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// Difference calculates the difference set between two integer slices (elements in a but not in b).
//
// Parameters:
//   - a: first slice
//   - b: second slice
//
// Returns:
//   - []int: difference set result
//
// Example:
//
//	a := []int{1, 2, 3, 4, 5}
//	b := []int{3, 4, 5, 6, 7}
//	Difference(a, b)  // []int{1, 2}
func Difference(a, b []int) []int {
	bSet := make(map[int]bool)
	for _, num := range b {
		bSet[num] = true
	}

	result := make([]int, 0)
	for _, num := range a {
		if !bSet[num] {
			result = append(result, num)
		}
	}

	return result
}

// Intersection calculates the intersection of two integer slices (elements present in both slices).
//
// Parameters:
//   - a: first slice
//   - b: second slice
//
// Returns:
//   - []int: intersection result
//
// Example:
//
//	a := []int{1, 2, 3, 4, 5}
//	b := []int{3, 4, 5, 6, 7}
//	Intersection(a, b)  // []int{3, 4, 5}
func Intersection(a, b []int) []int {
	bSet := make(map[int]bool)
	for _, num := range b {
		bSet[num] = true
	}

	result := make([]int, 0)
	for _, num := range a {
		if bSet[num] {
			result = append(result, num)
		}
	}

	return result
}

// Contains checks if a slice contains a specified element.
//
// Type Parameters:
//   - T: comparable type
//
// Parameters:
//   - slice: slice to check
//   - item: item to search for
//
// Returns:
//   - bool: true if contains, false otherwise
//
// Example:
//
//	Contains([]int{1, 2, 3}, 2)  // true
//	Contains([]int{1, 2, 3}, 4)  // false
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Reverse reverses the order of an integer slice, modifies the original slice.
//
// Parameters:
//   - nums: integer slice to reverse
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5}
//	Reverse(nums)
//	// nums = []int{5, 4, 3, 2, 1}
func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// GenericReverse generic version of array reverse function.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - slice: slice to reverse
//
// Example:
//
//	strings := []string{"a", "b", "c"}
//	GenericReverse(strings)
//	// strings = []string{"c", "b", "a"}
func GenericReverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// TakeFirstN gets the first N elements from a slice.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - slice: source slice
//   - n: number of elements to get
//
// Returns:
//   - []T: first N elements
//
// Example:
//
//	TakeFirstN([]int{1, 2, 3, 4, 5}, 3)  // []int{1, 2, 3}
func TakeFirstN[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice
	}
	result := make([]T, n)
	copy(result, slice[:n])
	return result
}

// TakeLastN gets the last N elements from a slice.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - slice: source slice
//   - n: number of elements to get
//
// Returns:
//   - []T: last N elements
//
// Example:
//
//	TakeLastN([]int{1, 2, 3, 4, 5}, 3)  // []int{3, 4, 5}
func TakeLastN[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return slice
	}
	start := len(slice) - n
	result := make([]T, n)
	copy(result, slice[start:])
	return result
}

// Filter filters slice elements based on condition function.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - slice: slice to filter
//   - condition: filter condition function, returns true to keep element
//
// Returns:
//   - []T: filtered slice
//
// Example:
//
//	nums := []int{1, 2, 3, 4, 5, 6}
//	Filter(nums, func(n int) bool { return n%2 == 0 })
//	// []int{2, 4, 6}
func Filter[T any](slice []T, condition func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range slice {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map applies transformation function to each element in slice.
//
// Type Parameters:
//   - T: source element type
//   - U: transformed element type
//
// Parameters:
//   - slice: source slice
//   - fn: transformation function
//
// Returns:
//   - []U: transformed slice
//
// Example:
//
//	nums := []int{1, 2, 3}
//	Map(nums, func(n int) string { return strconv.Itoa(n) })
//	// []string{"1", "2", "3"}
func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Sort sorts an integer slice using quick sort (ascending), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to sort
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	Sort(nums)
//	// nums = []int{1, 2, 5, 8, 9}
func Sort(nums []int) {
	sort.Ints(nums)
}

// SortDescending sorts an integer slice using quick sort (descending), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to sort
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	SortDescending(nums)
//	// nums = []int{9, 8, 5, 2, 1}
func SortDescending(nums []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
}

// GenericSort generic sort function with custom comparison rule.
//
// Type Parameters:
//   - T: any type
//
// Parameters:
//   - slice: slice to sort
//   - less: comparison function, returns true if a should come before b
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//	people := []Person{{"Alice", 30}, {"Bob", 25}}
//	GenericSort(people, func(a, b Person) bool { return a.Age < b.Age })
func GenericSort[T any](slice []T, less func(a, b T) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
}
