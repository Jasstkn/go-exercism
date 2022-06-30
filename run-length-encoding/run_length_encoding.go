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

	for i := 0; i < len(input); i++ {
		if string(input[i]) == previous {
			slice += string(input[i])
		} else {
			res += zipSlice(strings.Count(slice, previous), previous)
			slice = string(input[i])
		}
		previous = string(input[i])
		if i == len(input)-1 {
			res += zipSlice(strings.Count(slice, previous), previous)
		}
	}

	return res
}

func RunLengthDecode(input string) (res string) {
	number := ""
	for _, v := range input {
		switch {
		// check if current symbol is a number
		case unicode.IsDigit(v):
			number += string(v)
		// number is empty and current symbol isn't a number as well -> it's a single letter
		case number == "":
			res += string(v)
		// default case when we got the number
		default:
			numberInt, _ := strconv.Atoi(number)
			res += strings.Repeat(string(v), numberInt)
			number = ""
		}
	}
	return res
}
