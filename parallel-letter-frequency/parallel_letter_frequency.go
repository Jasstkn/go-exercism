package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	freqs := make(chan FreqMap)

	for _, s := range l {
		go func(s string) {
			freqs <- Frequency(s)
		}(s)
	}

	res := FreqMap{}
	for range l {
		m := <-freqs
		for k, v := range m {
			res[k] += v
		}
	}
	// close the channel to prevent memory leak
	close(freqs)

	return res
}
