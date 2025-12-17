package main

import (
	"fmt"
	"sync"
)

type BankAccount struct{
	Balance int
 	sync.Mutex
}

type PhoneBook struct{
	Phones map[string]string
	sync.RWMutex
}

func (p *PhoneBook) Set(name, phone string) {
	p.Lock()
	p.Phones[name] = phone
	p.Unlock()
}

func (p *PhoneBook) Get(name string) string{
	p.RLock()
	defer p.RUnlock()
	return p.Phones[name]
}

func (b *BankAccount) Deposit(wg *sync.WaitGroup){
	b.Lock()
	defer b.Unlock()
	b.Balance++
	wg.Done()
}

func main(){
	var phone PhoneBook
	var wg sync.WaitGroup
	phone.Phones = map[string]string{}
	pho := map[string]string{
		"Tamerlan": "+79890403016",
		"Linda": "+79896282745",
		"Misha": "+79896835245",
		"Sasha": "+79896647514",
		"Zaur": "+79893305620",
	}
	for name, phon := range pho{
		wg.Add(1)
		go func(name, phon string){
			defer wg.Done()
			phone.Set(name, phon)
			a := phone.Get(name)
			fmt.Println(a)
		}(name, phon)
	}
	wg.Wait()
}