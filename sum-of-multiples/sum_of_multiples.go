package summultiples

func isDivisable(n int, divisors []int) bool {
	for _, d := range divisors {
		if d!=0 && n%d == 0 {
			return true
		}
	}
	return false
}

func SumMultiples(limit int, divisors ...int) (out int) {
	for i := 1; i < limit; i++ {
		if isDivisable(i, divisors) {
			out += i
		}
	}

	return out
}
