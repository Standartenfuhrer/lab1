package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Goroutine says hello")
	}()

	time.Sleep(500 * time.Millisecond)
}

//Время не меняется, а количество выводимых строк меняется
