package luhn

import (
	"strconv"
	"strings"
)

func isValidInput(id string) bool {
	return len(id) > 1
}

func calculateNumber(v int, pos int) int {
	number := 0
	if pos%2 == 0 {
		number = 2 * v
	} else {
		number = v
	}
	if number > 9 {
		number -= 9
	}
	return number
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if !isValidInput(id) {
		return false
	}

	sum := 0

	for i := len(id) - 1; i >= 0; i-- {
		// process only numbers
		v, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		// calculate number to be summed
		sum += calculateNumber(v, len(id)-i)
	}

	return sum%10 == 0
}
