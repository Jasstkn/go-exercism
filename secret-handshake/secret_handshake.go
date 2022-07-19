package secret

func reverse(input []string) []string {
	for i, j := 0, len(input) - 1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

func Handshake(code uint) (out []string) {
	// 1 is wink
	// 10 is double blink
	// 100 is close your eyes
	// 1000 is jump
	// 10000 is to revert array
	if code&1 > 0 {
		out = append(out, "wink")
	}
	if code&2 > 0 {
		out = append(out, "double blink")
	}
	if code&4 > 0 {
		out = append(out, "close your eyes")
	}
	if code&8 > 0 {
		out = append(out, "jump")
	}

	if code&16 > 0 {
		out = reverse(out)
	}

	return out
}
