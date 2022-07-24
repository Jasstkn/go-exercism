package atbash

import (
	"strings"
)

const shift = 26

func Atbash(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	var counter int
	for i, char := range strings.ToLower(s) {
		if counter == 5 && i != len(s)-1 {
			counter = 0
			b.WriteRune(' ')
		}

		switch {
		case char >= 'a' && char <= 'z':
			b.WriteRune((('z' - char + rune(shift)) % 26) + 'a')
			counter++
		case char >= '0' && char <= '9':
			b.WriteRune(char)
			counter++
		}
	}
	return b.String()
}
