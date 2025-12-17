package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct{
	ID int
	Customer string
	Amount float64
}

func processOrder(or Order, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Обработка заказа №%d для %s\n", or.ID, or.Customer)
	time.Sleep(1 * time.Second)
	fmt.Printf("Заказ №%d на сумму %f успешно оформлен\n", or.ID, or.Amount)
}

func main(){
	var wg sync.WaitGroup
	orders := []Order{
		{ID: 1, Customer: "Тамерлан", Amount: 1000},
		{ID: 2, Customer: "Миша", Amount: 2000},
		{ID: 3, Customer: "Линда", Amount: 3000},
		{ID: 4, Customer: "Саша", Amount: 4000},
		{ID: 5, Customer: "Вова", Amount: 5000},
		{ID: 6, Customer: "Заур", Amount: 6000},
		{ID: 7, Customer: "Софа", Amount: 7000},
	}

	for _, val := range orders{
		wg.Add(1)
		go processOrder(val, &wg)
	}

	wg.Wait()
}