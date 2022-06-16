package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i, c := range strings.ToLower(word) {
		if unicode.IsLetter(c) && strings.ContainsRune(word[i+1:], c) {
			return false
		}
	}
	return true
}
