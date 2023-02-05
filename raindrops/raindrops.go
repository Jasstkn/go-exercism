package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(number int) string {
	var reply bytes.Buffer
	if number%3 == 0 {
		reply.WriteString("Pling")
	}
	if number%5 == 0 {
		reply.WriteString("Plang")
	}
	if number%7 == 0 {
		reply.WriteString("Plong")
	}
	if reply.Len() == 0 {
		reply.WriteString(strconv.Itoa(number))
	}

	return reply.String()
}
