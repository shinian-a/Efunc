// Package utils provides English-named retry utility functions, supporting repeated execution with backoff strategies.
package utils

import (
	"time"
)

// Retry executes a function with specified maximum retries and interval, automatically retries on failure.
//
// Type Parameters:
//   - T: return value type
//
// Parameters:
//   - fn: function to execute, returns value and error
//   - maxRetries: maximum number of retries
//   - interval: wait time between retries
//
// Returns:
//   - T: return value on success
//   - error: last error if all retries fail
//
// Example:
//
//	result, err := Retry(func() (string, error) {
//		return "success", nil
//	}, 3, time.Second)
func Retry[T any](fn func() (T, error), maxRetries int, interval time.Duration) (T, error) {
	var result T
	var err error

	for i := 0; i <= maxRetries; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if i < maxRetries {
			time.Sleep(interval)
		}
	}

	var zero T
	return zero, err
}

// ExponentialBackoffRetry implements retry with exponential backoff strategy.
//
// Type Parameters:
//   - T: return value type
//
// Parameters:
//   - fn: function to execute, returns value and error
//   - maxRetries: maximum number of retries
//   - baseInterval: base interval time
//   - maxInterval: maximum interval (prevents infinite growth)
//
// Returns:
//   - T: return value on success
//   - error: last error if all retries fail
//
// Example:
//
//	result, err := ExponentialBackoffRetry(func() (string, error) {
//		return "success", nil
//	}, 5, time.Second, time.Minute)
func ExponentialBackoffRetry[T any](fn func() (T, error), maxRetries int, baseInterval, maxInterval time.Duration) (T, error) {
	var result T
	var err error

	interval := baseInterval
	for i := 0; i <= maxRetries; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if i < maxRetries {
			time.Sleep(interval)
			// Exponential growth, but not exceeding maximum
			interval = interval * 2
			if interval > maxInterval {
				interval = maxInterval
			}
		}
	}

	var zero T
	return zero, err
}

// RetryWithCondition implements retry based on error type condition.
//
// Type Parameters:
//   - T: return value type
//
// Parameters:
//   - fn: function to execute, returns value and error
//   - shouldRetry: function to determine if retry should happen, takes current error
//   - maxRetries: maximum number of retries
//   - interval: wait time between retries
//
// Returns:
//   - T: return value on success
//   - error: last error if all retries fail
//
// Example:
//
//	result, err := RetryWithCondition(
//		func() (string, error) {
//			return "success", nil
//		},
//		func(err error) bool {
//			return err != nil && err.Error() == "retryable"
//		},
//		3,
//		time.Second,
//	)
func RetryWithCondition[T any](fn func() (T, error), shouldRetry func(error) bool, maxRetries int, interval time.Duration) (T, error) {
	var result T
	var err error

	for i := 0; i <= maxRetries; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if !shouldRetry(err) || i >= maxRetries {
			break
		}
		time.Sleep(interval)
	}

	var zero T
	return zero, err
}
