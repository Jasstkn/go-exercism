package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) (triplets []Triplet) {
	for v1 := min; v1 <= max; v1++ {
		for v2 := min; v2 <= max; v2++ {
			for v3 := min; v3 <= max; v3++ {
				if v1*v1+v2*v2 == v3*v3 && v1 < v2 && v2 < v3 {
					triplets = append(triplets, Triplet{v1, v2, v3})
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) (triplets []Triplet) {
	for _, t := range Range(0, p) {
		if t[0]+t[1]+t[2] == p{
			triplets = append(triplets, t)
		}
	}
	return triplets
}
