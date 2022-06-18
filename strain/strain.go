package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (out Ints) {
	if i == nil {
		return Ints(nil)
	}

	for _, num := range i {
		if filter(num) {
			out = append(out, num)
		}
	}
	return out
}

func (i Ints) Discard(filter func(int) bool) (out Ints) {
	if i == nil {
		return Ints(nil)
	}
	for _, num := range i {
		if !filter(num) {
			out = append(out, num)
		}
	}
	return out
}

func (l Lists) Keep(filter func([]int) bool) (out Lists) {
	for _, v := range l {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}

func (s Strings) Keep(filter func(string) bool) (out Strings) {
	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}
