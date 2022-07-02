package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) (triplets []Triplet) {
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			cInt := math.Trunc(c)
			if cInt == c && a < b && float64(b) < c && c <= float64(max) {
				triplets = append(triplets, Triplet{a, b, int(c)})
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (triplets []Triplet) {
	for _, t := range Range(0, p) {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}
	return triplets
}
