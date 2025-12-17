package main

import (
	"fmt"
<<<<<<< HEAD
	"time"
)

// printNumbers печатает числа от 1 до n с небольшими паузами.
func printNumbers(id int, n int) {
	for i := id; i <= n; i++ {
		fmt.Printf("Worker %d: number %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Start lab1")
	go printNumbers(1, 5)
	fmt.Println("End lab1")
=======
	
)

func main() {
   a := map[string]string{
      "asfaf": "asds",
      "tams": "azam",
      "misha": "cyka",
      "Linda": "lin",
   }

   for name, val := range a{
      fmt.Println(name, val)
   }
>>>>>>> 436fed9bf1941acdd224b531a89c7c8a75f41385
}
