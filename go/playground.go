package main

import (
	parallel "exercism/go/parallel-letter-frequency"
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

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
/*ConcurrentFrequency counts word occurances using parallel processes*/
func ConcurrentFrequency(texts []string) FreqMap {
	res := make(FreqMap)
	ch := make(chan FreqMap, 10)
	var wg sync.WaitGroup
	wg.Add(len(texts))
	for _, text := range texts {
		go func(t string) {
			ch <- Frequency(t)
			wg.Done()
		}(text)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for freqmap := range ch {
		for letter, freq := range freqmap {
			res[letter] += freq
		}
	}
	return res
}
func main() {
	startTime := time.Now()
	Frequency(parallel.Dostoevsky1)
	Frequency(parallel.Dostoevsky2)
	Frequency(parallel.Dostoevsky3)
	Frequency(parallel.Dostoevsky4)
	endTime := time.Now()
	pl("Time taken ", endTime.Sub(startTime).Microseconds())

	startTime = time.Now()
	l := []string{parallel.Dostoevsky1, parallel.Dostoevsky2, parallel.Dostoevsky3, parallel.Dostoevsky4}
	// pl(len(l))
	ConcurrentFrequency(l)
	endTime = time.Now()
	pl("Time taken ", endTime.Sub(startTime).Microseconds())

}
