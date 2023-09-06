package main

import (
	"fmt"
	"sync"
)

func countLetters(text string, letterCounts map[rune]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for _, char := range text {
		mu.Lock()
		letterCounts[char]++
		mu.Unlock()
	}
}

func main() {
	text := "hahahahahahah"
	letterCounts := make(map[rune]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, char := range text {
		if char != ' ' { // Skip spasi
			letterCounts[char] = 0
		}
	}

	for _, char := range text {
		if char != ' ' { // Skip spasi
			wg.Add(1)
			go countLetters(text, letterCounts, &wg, &mu)
		}
	}

	wg.Wait()

	for char, count := range letterCounts {
		fmt.Printf("Huruf %c: %d\n", char, count)
	}
}