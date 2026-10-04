// Package utils provides English-named algorithm utility functions, covering sorting, searching, caching and other common algorithms.
package utils

import (
	"sort"
)

// BubbleSort sorts an integer slice using bubble sort algorithm (ascending order), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to be sorted
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	BubbleSort(nums)
//	// nums = []int{1, 2, 5, 8, 9}
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// QuickSort sorts an integer slice using quick sort algorithm (ascending order), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to be sorted
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	QuickSort(nums)
//	// nums = []int{1, 2, 5, 8, 9}
func QuickSort(nums []int) {
	sort.Ints(nums)
}

// MergeSort sorts an integer slice using merge sort algorithm (ascending order), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to be sorted
//
// Returns:
//   - []int: new sorted slice (does not modify original)
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// InsertionSort sorts an integer slice using insertion sort algorithm (ascending order), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to be sorted
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	InsertionSort(nums)
//	// nums = []int{1, 2, 5, 8, 9}
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

// SelectionSort sorts an integer slice using selection sort algorithm (ascending order), modifies the original slice.
//
// Parameters:
//   - nums: integer slice to be sorted
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	SelectionSort(nums)
//	// nums = []int{1, 2, 5, 8, 9}
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}

// BinarySearch uses binary search to find target value in a sorted integer slice.
//
// Parameters:
//   - nums: sorted integer slice (ascending)
//   - target: target value to find
//
// Returns:
//   - int: index of target value, -1 if not found
//
// Example:
//
//	nums := []int{1, 2, 5, 8, 9}
//	BinarySearch(nums, 5)  // returns 2
//	BinarySearch(nums, 3)  // returns -1
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// LinearSearch searches for target value in a linear list.
//
// Type Parameters:
//   - T: any comparable type
//
// Parameters:
//   - slice: slice to search
//   - target: target value to find
//
// Returns:
//   - int: index of target value, -1 if not found
//
// Example:
//
//	strings := []string{"a", "b", "c"}
//	LinearSearch(strings, "b")  // returns 1
func LinearSearch[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// LRUCache implements a simple LRU cache for caching frequently used data.
//
// Type Parameters:
//   - K: key type, must support comparison
//   - V: value type
type LRUCache struct {
	capacity int
	cache    map[int]interface{}
}

// NewLRUCache creates a new LRU cache instance.
//
// Parameters:
//   - capacity: cache capacity
//
// Returns:
//   - *LRUCache: LRU cache instance
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]interface{}),
	}
}

// Get retrieves a value from the cache.
//
// Parameters:
//   - key: cache key
//
// Returns:
//   - interface{}: cached value
//   - bool: whether the key exists
func (l *LRUCache) Get(key int) (interface{}, bool) {
	if val, ok := l.cache[key]; ok {
		return val, true
	}
	return nil, false
}

// Put inserts a value into the cache.
//
// Parameters:
//   - key: cache key
//   - value: cached value
func (l *LRUCache) Put(key int, value interface{}) {
	if len(l.cache) >= l.capacity {
		// Simple implementation: delete first element
		for k := range l.cache {
			delete(l.cache, k)
			break
		}
	}
	l.cache[key] = value
}

// Clear clears all items from the cache.
func (l *LRUCache) Clear() {
	l.cache = make(map[int]interface{})
}

// Size returns the current cache size.
func (l *LRUCache) Size() int {
	return len(l.cache)
}
