package logs

import (
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	mapping := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, v := range log {
		if v, ok := mapping[v]; ok {
			return v
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// runeStr := []rune(log)
	// for index, rune := range runeStr {
	// 	if rune == oldRune {
	// 		runeStr[index] = newRune
	// 	}
	// }
	// return string(runeStr)

	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	logRune := []rune(log)
	return len(logRune) <= limit
}
