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
		printNumbers(1, 5)
       fmt.Println("End lab1")
   }
