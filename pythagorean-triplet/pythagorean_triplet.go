package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) (triplets []Triplet) {
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if c == float64(int(c)) && int(c) <= max {
				triplets = append(triplets, Triplet{a, b, int(c)})
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (triplets []Triplet) {
	for _, t := range Range(1, p/2) {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}
	return triplets
}
