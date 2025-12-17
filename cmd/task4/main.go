package main

import (
	"fmt"
	"time"
)

// printNumbers печатает числа от 1 до n с небольшими паузами
func printNumbers(id int, n int) {
	for i := id; i <= n; i++ {
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

//Не все строки успевают напечататься за 100 миллисекунд
//Не все числа от 1 до 5 для каждого id видны за 100 миллисекунд
//time.Sleep ненадежный способ ожидания загрузки горутин, потому что мы не знаем
//сколько времени им нужно для завершения, если они будут работать дольше, чем мы
//указали, то программа завершится слишком рано, а если они будут работать
//меньше, чем мы указали, то мы в пустую потратим время на ожидание.
