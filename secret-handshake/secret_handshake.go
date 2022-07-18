package secret

var codeMapping = map[int]string{
	1:    "wink",
	10:   "double blink",
	100:  "close your eyes",
	1000: "jump",
}

func reverse(input []string) []string {
	i := 0
	j := len(input) - 1
	for i < j {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}
	return input
}

func Handshake(code uint) (out []string) {
	if code&1 > 0 {
		out = append(out, codeMapping[1])
	}
	if code&2 > 0 {
		out = append(out, codeMapping[10])
	}
	if code&4 > 0 {
		out = append(out, codeMapping[100])
	}
	if code&8 > 0 {
		out = append(out, codeMapping[1000])
	}

	if code&16 > 0 {
		out = reverse(out)
	}

	return out
}
