package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	sites := []string{"google.com", "yandex.ru", "github.com", "stackoverflow.com"}

	for _, site := range sites{
		wg.Add(1)
		go func(sit string){
			defer wg.Done()
			fmt.Printf("Начинаю проверку <%s>\n", sit)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Проверка <%s> завершена\n", sit)
		}(site)
	}
	wg.Wait()
}