package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	ordA           = 65
	alphabetLength = 26
	numberLimit    = 1000
	namesLimit     = alphabetLength * alphabetLength * numberLimit
)

// generate pool required length
func generateNamePool(size int) []int {
	pool := make([]int, size)
	for i := 0; i < size; i++ {
		pool[i] = i
	}
	rand.Shuffle(size, func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})
	return pool
}

var namesPool = generateNamePool(namesLimit)

// Define the Robot type here.

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.Reset()
	}
	if r.name == "" {
		return "", errors.New("no uniq names remain")
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	if len(namesPool) == 0 {
		r.name = ""
		return
	}
	r.name = numberToText(namesPool[0])
	namesPool = namesPool[1:]
}

func numberToText(num int) string {
	letterPart := num / numberLimit
	part1 := letterPart / alphabetLength
	part2 := letterPart % alphabetLength
	part3 := num % numberLimit
	return fmt.Sprintf("%v%v%03d", numberToLetter(part1), numberToLetter(part2), part3)
}

// numberToLetter returns A for 0, B for 1, ..., Z for 25
func numberToLetter(number int) string {
	return string(rune(number + ordA))
}
