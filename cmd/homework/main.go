package main

import (
	"fmt"
	"strings"
	"sync"
)

func ConcurrentWordCount(sentences []string) map[string]int {
	words := map[string]int{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, val := range sentences{
		wg.Add(1)
		go func(val string){
			defer wg.Done()
			mu.Lock()
			for j := 0; j < len(strings.Fields(val)); j++{
				words[strings.Fields(val)[j]]++
			}
			mu.Unlock()
		}(val)
	}
	wg.Wait()
	return words
}

func main() {
	text := []string{
		"quick brown fox",
		"lazy dog",
		"quick brown fox jumps",
		"jumps over lazy dog",
	}
	a := ConcurrentWordCount(text)
	fmt.Println(a)
}