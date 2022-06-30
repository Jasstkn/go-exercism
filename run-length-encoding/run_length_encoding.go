package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func zipSlice(count int, previous string) string {
	if count > 1 {
		return strconv.Itoa(count) + previous
	}
	return previous
}

func RunLengthEncode(input string) (res string) {
	slice := ""
	previous := ""

	var b strings.Builder
	for i := 0; i < len(input); i++ {
		c := string(input[i])
		if c == previous {
			slice += c
		} else {
			b.WriteString(zipSlice(strings.Count(slice, previous), previous))
			slice = c
		}
		previous = c
		if i == len(input)-1 {
			b.WriteString(zipSlice(strings.Count(slice, previous), previous))
		}
	}

	return b.String()
}

func RunLengthDecode(input string) string {
	number := ""
	var b strings.Builder
	for _, v := range input {
		switch {
		// check if current symbol is a number
		case unicode.IsDigit(v):
			number += string(v)
		// number is empty and current symbol isn't a number as well -> it's a single letter
		case number == "":
			b.WriteRune(v)
		// default case when we got the number
		default:
			numberInt, _ := strconv.Atoi(number)
			b.WriteString(strings.Repeat(string(v), numberInt))
			number = ""
		}
	}
	return b.String()
}
