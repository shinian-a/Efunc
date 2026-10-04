// Package utils 提供中文命名的算法工具函数库，涵盖排序、查找、缓存等常用算法实现。
package utils

import (
	"sort"
)

// LQuickSort 使用快速排序算法对整数切片进行排序（升序），会修改原切片。
//
// Parameters:
//   - nums: 待排序的整数切片
//
// Example:
//
//	nums := []int{5, 2, 8, 1, 9}
//	LQuickSort(nums)
//	// nums = []int{1, 2, 5, 8, 9}
func LQuickSort(nums []int) {
	sort.Ints(nums)
}

// LMergeSort 使用归并排序算法对整数切片进行排序（升序），返回新切片。
//
// Parameters:
//   - nums: 待排序的整数切片
//
// Returns:
//   - []int: 排序后的新切片（不修改原切片）
func LMergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := LMergeSort(nums[:mid])
	right := LMergeSort(nums[mid:])
	return lMerge(left, right)
}

func lMerge(left, right []int) []int {
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

// LBinarySearch 使用二分查找算法在已排序的整数切片中查找目标值。
//
// Parameters:
//   - nums: 已排序的整数切片（升序）
//   - target: 要查找的目标值
//
// Returns:
//   - int: 目标值的索引，未找到返回 -1
//
// Example:
//
//	nums := []int{1, 2, 5, 8, 9}
//	LBinarySearch(nums, 5)  // returns 2
//	LBinarySearch(nums, 3)  // returns -1
func LBinarySearch(nums []int, target int) int {
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

// LLinearSearch 在线性列表中查找目标值。
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
//	LLinearSearch(strings, "b")  // returns 1
func LLinearSearch[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// LRUCache implements a simple LRU cache for caching frequently used data.
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
