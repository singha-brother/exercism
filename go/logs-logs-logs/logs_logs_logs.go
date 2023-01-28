package logs

import "unicode/utf8"

var sChar = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather"}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if c, ok := sChar[char]; ok {
			return c
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newStr []int32
	for _, v := range log {
		if v == oldRune {
			newStr = append(newStr, newRune)
		} else {
			newStr = append(newStr, v)
		}
	}
	return string(newStr)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
