package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 10000; i++ {
       // TODO: запустить горутину, которая что-нибудь печатает
	   go fmt.Printf("Goroutine %d says hello\n", i)
    }

    time.Sleep(500 * time.Millisecond)
}

//Время не меняется, а количество выводимых строк меняется