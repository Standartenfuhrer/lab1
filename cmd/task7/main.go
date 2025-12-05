package main

import "fmt"

func generate(ch chan<- int, n int){
	for i := 1; i <= n; i++{
		ch <- i
	}
	close(ch)
}

func main(){
	myCh := make(chan int)
	summ := 0

	go generate(myCh, 10)
	for v := range myCh{
		summ += v
	}

	fmt.Println(summ)
}