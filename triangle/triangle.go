// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = Kind("NaT") // not a triangle
	Equ = Kind("Equ") // equilateral
	Iso = Kind("Iso") // isosceles
	Sca = Kind("Sca") // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a || math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return NaT
	}
	if a == b && a == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else if a+b >= c || b+c >= a || a+c >= b {
		return Sca
	}
	return NaT
}
