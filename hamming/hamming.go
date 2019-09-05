// Package hamming returns hamming difference between two DNAs
package hamming

import (
	"errors"
)

// Distance is func which returns hamming difference between two DNAs
func Distance(a, b string) (int, error) {
	count := 0
	var err error
	aRune := []rune(a)
	bRune := []rune(b)
	if len(aRune) == len(bRune) && (len(aRune) != 0 || len(bRune) != 0) {
		for i := 0; i < len(aRune); i++ {
			if aRune[i] != bRune[i] {
				count++
			}
		}
	} else if len(aRune) != len(bRune) {
		err = errors.New("DNAs should have the same length")
	}
	return count, err
}
