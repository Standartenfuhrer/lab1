package main

import (
	"fmt"
	"time"
)

func newsFeed(ch chan <- string){
	for i := 0; i < 5; i++{
		time.Sleep(1 * time.Second)
		ch <- "эщкереее"
	}
}

func socialMedia(ch chan <- string){
	for i := 0; i < 5; i++{
		time.Sleep(2500 * time.Millisecond)
		ch <- "не эщкере."
	}
}

func main(){
	news := make(chan string)
	media := make(chan string)

	go newsFeed(news)
	go socialMedia(media)
	
	timeout := time.After(5 * time.Second)
	for {
		select{
		case msg := <- news:
			fmt.Println("Новости:", msg)
		case msg := <- media:
			fmt.Println("Соцсети:", msg)
		case <- timeout:
			fmt.Println("Время вышло, выключаемся")
			return
		default:
			fmt.Println(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}