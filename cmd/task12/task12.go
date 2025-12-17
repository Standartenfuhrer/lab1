package main

import (
	"fmt"
	"sync"
)

func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n * factorial(n - 1)
}

func main(){
	numbers := []int{5, 2, 10, 7}
	var wg sync.WaitGroup
	for _, val := range numbers{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Printf("Факторилал %d = %d\n", val, factorial(val))
		}()
	}

	wg.Wait()
}