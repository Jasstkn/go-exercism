package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return Ints(nil)
	}
	res := Ints{}
	for _, num := range i {
		if filter(num) {
			res = append(res, num)
		}
	}
	return res
}

func (i Ints) Discard(filter func(int) bool) Ints {
	if i == nil {
		return Ints(nil)
	}
	res := Ints{}
	for _, num := range i {
		if !filter(num) {
			res = append(res, num)
		}
	}
	return res
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	res := Lists{}

	for _, v := range l {
		if filter(v) {
			res = append(res, v)
		}
	}

	return res
}

func (s Strings) Keep(filter func(string) bool) Strings {
	res := Strings{}

	for _, v := range s {
		if filter(v) {
			res = append(res, v)
		}
	}

	return res
}
