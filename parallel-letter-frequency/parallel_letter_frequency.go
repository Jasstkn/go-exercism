package letter

import (
	"sync"
)

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
	c := make(chan FreqMap)
	var wg sync.WaitGroup

	for _, s := range l {
		wg.Add(1)
		go func(s string) {
			c <- Frequency(s)
			defer wg.Done()
		}(s)
	}

	// close the channel to prevent memory leak
	go func() {
		// wait until counter is 0
		wg.Wait()
		close(c)
	}()

	// process data
	res := FreqMap{}
	for m := range c {
		for k, v := range m {
			res[k] += v
		}
	}

	return res
}
