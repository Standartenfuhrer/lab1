package main

import (
	"fmt"
	"sync"
)

type PhoneBook struct {
	Phones map[string]string
	sync.RWMutex
}

func (p *PhoneBook) Set(name, phone string) {
	p.Lock()
	defer p.Unlock()
	p.Phones[name] = phone
}

func (p *PhoneBook) Get(name string) (string, bool) {
	p.RLock()
	defer p.RUnlock()
	phone, ok := p.Phones[name]
	return phone, ok
}

func main() {
	phones := PhoneBook{}
	var wg sync.WaitGroup
	phones.Phones = map[string]string{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			phones.Set(fmt.Sprintf("Record %d", i), fmt.Sprintf("+798904030%02d", i))
		}
	}()

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if val, ok := phones.Get(fmt.Sprintf("Record %d", i)); ok {
				fmt.Println(val)
			} else {
				fmt.Println("Ошибка")
			}
		}()
	}
	wg.Wait()
}
