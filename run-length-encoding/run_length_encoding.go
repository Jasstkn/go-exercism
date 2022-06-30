package encode

import (
	"strconv"
	"strings"
)

type builder struct {
	strings.Builder
}

func (b *builder) writeEncoded(count int, r rune) {
	switch {
	case count == 1:
		b.WriteRune(r)
	case count > 1:
		b.WriteString(strconv.Itoa(count))
		b.WriteRune(r)
	}
}

func RunLengthEncode(input string) (res string) {
	var count int
	var previous rune

	var out builder
	out.Grow(len(input)) // allocate space in worst-case scenario

	for _, r := range input {
		if previous == r && count > 0 {
			count++
		} else {
			out.writeEncoded(count, previous)
			count = 1
			previous = r
		}
	}
	out.writeEncoded(count, previous)
	return out.String()
}

func RunLengthDecode(input string) string {
	var count int
	var out builder
	out.Grow(len(input) * 2)
	for _, r := range input {
		switch {
		// check if current symbol is a number
		case '0' <= r && r <= '9':
			count = count*10 + int(r-'0')
		// number is empty and current symbol isn't a number as well -> it's a single letter
		case count == 0:
			out.WriteRune(r)
		// default case when we got the number
		default:
			out.WriteString(strings.Repeat(string(r), count))
			count = 0
		}
	}
	return out.String()
}
