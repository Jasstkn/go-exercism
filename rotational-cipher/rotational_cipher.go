package rotationalcipher

import (
	"strings"
)

func RotationalCipher(plain string, shiftKey int) string {
	var b strings.Builder
	b.Grow(len(plain))

	for _, char := range plain {
		switch {
		case char >= 'a' && char <= 'z':
			b.WriteRune(((char - 'a' + rune(shiftKey)) % 26) + 'a')
		case char >= 'A' && char <= 'Z':
			b.WriteRune(((char - 'A' + rune(shiftKey)) % 26) + 'A')
		default:
			b.WriteRune(char)
		}
	}
	return b.String()
}
