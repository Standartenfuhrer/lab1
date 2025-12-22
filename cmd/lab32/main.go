package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Order struct{
	ID int
	Amount int
	Status string
}

func generateOrders(count int) <- chan Order{
	out := make(chan Order)

	for i := 1; i <= count; i++{
		out <- Order{
			ID: i,
			Amount: rand.Intn(10000),
			Status: "new",
		}
	}
	return out
}

func processOrders(in <- chan Order) <- chan Order{
	out := make(chan Order)
	var wg sync.WaitGroup
	wg.Add(3)
	go func(){
		defer wg.Done()
		for c := range in{
			c.Status = "processed"
			out <- c
		}
	}()
	go func(){
		defer wg.Done()
		for c := range in{
			c.Status = "processed"
			out <- c
		}
	}()
	go func(){
		defer wg.Done()
		for c := range in{
			c.Status = "processed"
			out <- c
		}
	}()

	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}

func filterOrders(in <- chan Order, minAmount int) <- chan Order{
	out := make(chan Order)

	for c := range in{
		if minAmount < c.Amount{
			out <- c
		}
	}

	return out
}

func main(){
	in := generateOrders(50)

	c1 := processOrders(in)

	out := filterOrders(c1, 100)

	for result := range out{
		fmt.Println("Result:", result)
	}
}