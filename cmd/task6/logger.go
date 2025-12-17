package main

import (
	"fmt"
	"time"
)

func logger(logs <-chan int){
	for{
		fmt.Println("[LOG]:", <-logs)
	}
}

func main(){
	logCh := make(chan int, 5)

	go logger(logCh)

	logCh <- 1
	logCh <- 2
	logCh <- 3
	logCh <- 4
	logCh <- 5
	logCh <- 6
	logCh <- 7
	logCh <- 8

	time.Sleep(2 * time.Second)
}