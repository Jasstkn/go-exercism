// Package hamming returns hamming difference between two DNAs
package hamming

import (
	"errors"

	"github.com/mozillazg/go-unidecode"
)

// Distance is func which returns hamming difference between two DNAs
func Distance(a, b string) (int, error) {
	count := 0
	var err error
	a = unidecode.Unidecode(a)
	b = unidecode.Unidecode(b)
	if len(a) == len(b) && (len(a) != 0 || len(b) != 0) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
	} else if len(a) != len(b) {
		err = errors.New("DNAs should have the same length")
	}
	return count, err
}
