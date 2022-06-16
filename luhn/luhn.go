package luhn

import (
	"errors"
	"strings"
)

func isValidInput(id string) bool {
	return len(id) > 1
}

func calculateNumber(v byte, pos int) (int, error) {
	if v >= '0' && v <= '9' {
		number := 0
		if pos%2 == 0 {
			number = 2 * int(v-'0')
		} else {
			number = int(v - '0')
		}
		if number > 9 {
			number -= 9
		}
		return number, nil
	}

	return 0, errors.New("non-digit characters are disallowed")
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if !isValidInput(id) {
		return false
	}

	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		// calculate number to be summed
		num, err := calculateNumber(id[i], len(id)-i)
		if err != nil {
			return false
		}
		sum += num
	}

	return sum%10 == 0
}
