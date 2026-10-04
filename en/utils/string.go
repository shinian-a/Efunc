// Package utils provides English-named string operation utility functions, covering reverse, formatting and other operations.
package utils

import "strings"

// StringReverse reverses the order of characters in a string.
//
// Parameters:
//   - s: string to reverse
//
// Returns:
//   - string: reversed string
//
// Example:
//
//	StringReverse("hello")  // "olleh"
func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// TrimSpace removes leading and trailing spaces from a string.
//
// Parameters:
//   - s: string to process
//
// Returns:
//   - string: string with leading/trailing spaces removed
//
// Example:
//
//	TrimSpace("  hello world  ")  // "hello world"
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimLeft removes all leading spaces from a string.
//
// Parameters:
//   - s: string to process
//
// Returns:
//   - string: string with leading spaces removed
//
// Example:
//
//	TrimLeft("  hello world")  // "hello world"
func TrimLeft(s string) string {
	return strings.TrimLeft(s, " ")
}

// TrimRight removes all trailing spaces from a string.
//
// Parameters:
//   - s: string to process
//
// Returns:
//   - string: string with trailing spaces removed
//
// Example:
//
//	TrimRight("hello world  ")  // "hello world"
func TrimRight(s string) string {
	return strings.TrimRight(s, " ")
}

// ToLower converts a string to lowercase.
//
// Parameters:
//   - s: string to convert
//
// Returns:
//   - string: lowercase string
//
// Example:
//
//	ToLower("Hello World")  // "hello world"
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts a string to uppercase.
//
// Parameters:
//   - s: string to convert
//
// Returns:
//   - string: uppercase string
//
// Example:
//
//	ToUpper("Hello World")  // "HELLO WORLD"
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// FirstLetterToUpper converts the first letter of a string to uppercase.
//
// Parameters:
//   - s: string to process
//
// Returns:
//   - string: string with first letter uppercase
//
// Example:
//
//	FirstLetterToUpper("hello")  // "Hello"
func FirstLetterToUpper(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = rune(strings.ToUpper(string(runes[0]))[0])
	return string(runes)
}

// CamelToSnake converts camelCase naming to snake_case (lowercase).
//
// Parameters:
//   - s: camelCase string to convert
//
// Returns:
//   - string: snake_case string
//
// Example:
//
//	CamelToSnake("helloWorld")  // "hello_world"
//	CamelToSnake("XMLParser")    // "xml_parser"
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(rune(strings.ToLower(string(r))[0]))
	}
	return result.String()
}

// SnakeToCamel converts snake_case naming to camelCase (first letter lowercase).
//
// Parameters:
//   - s: snake_case string to convert
//
// Returns:
//   - string: camelCase string
//
// Example:
//
//	SnakeToCamel("hello_world")  // "helloWorld"
//	SnakeToCamel("user_name")     // "userName"
func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 0 {
		return s
	}

	var result strings.Builder
	result.WriteString(strings.ToLower(parts[0]))

	for _, part := range parts[1:] {
		if part != "" {
			result.WriteString(strings.ToUpper(string(part[0])))
			if len(part) > 1 {
				result.WriteString(part[1:])
			}
		}
	}

	return result.String()
}

// CountCharacters counts the occurrences of each character in a string.
//
// Parameters:
//   - s: string to count
//
// Returns:
//   - map[rune]int: character count map
//
// Example:
//
//	CountCharacters("hello")  // map[h:1 e:1 l:2 o:1]
func CountCharacters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, r := range s {
		count[r]++
	}
	return count
}

// CountWords counts the number of words in a string (separated by spaces).
//
// Parameters:
//   - s: string to count
//
// Returns:
//   - int: word count
//
// Example:
//
//	CountWords("hello world")  // 2
//	CountWords("one")          // 1
func CountWords(s string) int {
	if s == "" {
		return 0
	}
	return len(strings.Fields(s))
}

// ContainsSubstring checks if a string contains a specified substring.
//
// Parameters:
//   - s: source string
//   - substr: substring to search for
//
// Returns:
//   - bool: true if contains, false otherwise
//
// Example:
//
//	ContainsSubstring("hello world", "world")  // true
//	ContainsSubstring("hello world", "xyz")   // false
func ContainsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

// HasPrefix checks if a string starts with a specified prefix.
//
// Parameters:
//   - s: source string
//   - prefix: prefix to check
//
// Returns:
//   - bool: true if starts with prefix, false otherwise
//
// Example:
//
//	HasPrefix("hello.txt", ".txt")  // false
//	HasPrefix("hello.txt", "hello") // true
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if a string ends with a specified suffix.
//
// Parameters:
//   - s: source string
//   - suffix: suffix to check
//
// Returns:
//   - bool: true if ends with suffix, false otherwise
//
// Example:
//
//	HasSuffix("hello.txt", ".txt")  // true
//	HasSuffix("hello.txt", "hello") // false
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// ReplaceAll replaces all occurrences of a substring in a string.
//
// Parameters:
//   - s: source string
//   - old: substring to replace
//   - new: replacement string
//
// Returns:
//   - string: replaced string
//
// Example:
//
//	ReplaceAll("hello world world", "world", "go")  // "hello go go"
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Split splits a string by separator into string slice.
//
// Parameters:
//   - s: string to split
//   - sep: separator
//
// Returns:
//   - []string: split string slice
//
// Example:
//
//	Split("a,b,c", ",")  // []string{"a", "b", "c"}
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join joins a slice of strings into a single string.
//
// Parameters:
//   - parts: slice of strings to join
//   - sep: separator
//
// Returns:
//   - string: joined string
//
// Example:
//
//	Join([]string{"a", "b", "c"}, ",")  // "a,b,c"
func Join(parts []string, sep string) string {
	return strings.Join(parts, sep)
}

// Repeat repeats a string specified number of times.
//
// Parameters:
//   - s: string to repeat
//   - count: number of repetitions
//
// Returns:
//   - string: repeated string
//
// Example:
//
//	Repeat("abc", 3)  // "abcabcabc"
func Repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

// Substring extracts a substring starting from specified position.
//
// Parameters:
//   - s: source string
//   - start: start position (inclusive)
//
// Returns:
//   - string: extracted substring
//
// Example:
//
//	Substring("hello world", 6)  // "world"
func Substring(s string, start int) string {
	if start >= len(s) {
		return ""
	}
	return s[start:]
}

// SubstringLength extracts a substring of specified length from start position.
//
// Parameters:
//   - s: source string
//   - start: start position (inclusive)
//   - length: extraction length
//
// Returns:
//   - string: extracted substring
//
// Example:
//
//	SubstringLength("hello world", 0, 5)  // "hello"
func SubstringLength(s string, start, length int) string {
	end := start + length
	if start >= len(s) {
		return ""
	}
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}
