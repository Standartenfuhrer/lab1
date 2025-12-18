package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("Воркер %d: начал смену\n", id)
	
	for job := range jobs {
		fmt.Printf("Воркер %d: выполняет задачу %d\n", id, job.ID)
		time.Sleep(500 * time.Millisecond)
	}
	
	fmt.Printf("Воркер %d: закончил смену\n", id)
}

func main() {
	jobs := make(chan Job, 10) 
	var wg sync.WaitGroup

	numWorkers := 3
    fmt.Println("Запуск воркеров...")
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}
    numJobs := 10
	fmt.Println("Отправка задач...")
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	
	close(jobs)
	fmt.Println("Все задачи отправлены, канал закрыт.")
	wg.Wait()
	fmt.Println("Все задачи выполнены. Программа завершена.")
}
