package main

import (
    "fmt"
    "time"
)

// printNumbers печатает числа от 1 до n с небольшими паузами.
func printNumbers(id int, n int) {
	for i := id; i <= n; i++{
		fmt.Printf("Worker %d: number %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
    fmt.Println("Start lab1")

	go printNumbers(1, 5)
	go printNumbers(2, 5)
	go printNumbers(3, 5)
	go printNumbers(4, 5)
	go printNumbers(5, 5)

	time.Sleep(1 * time.Second)

    fmt.Println("End lab1")
}
