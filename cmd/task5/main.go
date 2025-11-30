package main

import "fmt"

func inputStream(ch chan int){
	for i := 1; i <= 10; i++{
		ch <- i
	}
	close(ch)
}

func doubler(ch chan int, ch2 chan int){
	for i := 0; i < 10; i++{
		num := <- ch
		ch2 <- num * 2
	}
	close(ch2)
}

func printer(ch chan int){
	for i := 0; i < 10; i++{
		fmt.Printf("Результат: %d\n", <-ch)
	}
}

func main() {
	intCh := make(chan int)
	intCh2 := make(chan int)
	go inputStream(intCh)
	go doubler(intCh, intCh2)
	printer(intCh2)
}
