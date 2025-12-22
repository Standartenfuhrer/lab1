package main

import (
	"fmt"
	"sync"
)

func generator(limit int) <-chan int {
	out := make(chan int)
	go func(){
		defer close(out)
		for i := 1; i <= limit; i++{
			out <- i
		}
	}()
	return out 
}

func worker(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func(){
		defer close(out)
		for n := range in {
			fmt.Printf("Worker %d proccesing %d\n", id, n)
			out <- n * n
		}
	}()
	return out
}

func merge(cs ...<-chan int) <- chan int{
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int){
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}
	wg.Add(len(cs))
	for _, c := range cs{
		go output(c)
	}

	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}

func main(){
	in := generator(20)

	c1 := worker(1, in)
	c2 := worker(2, in)
	c3 := worker(3, in)
	c4 := worker(4, in)

	out := merge(c1, c2, c3, c4)

	for result := range out{
		fmt.Println("Result:", result)
	}
}