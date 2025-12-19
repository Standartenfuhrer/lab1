package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct{
	ID int
	Number int
}

type Result struct{
	JobID    int
	Square    int
	Perimeter int
	WorkerID int
}

func worker(id int, jobs <- chan Job, results chan <- Result, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("Воркер %d начал смену\n", id)

	for job := range jobs{
		square := job.Number * job.Number
		perimeter := 4 * job.Number
		results <- Result{JobID: job.ID, Square: square, Perimeter: perimeter, WorkerID: id}
	}

	fmt.Printf("Воркер %d закончил смену\n", id)
}

func main(){
	inputs := []int{1, 5, 12, 5, 3, 8, 9}
	jobs := make(chan Job)
	var wg sync.WaitGroup

	numWorkers := 3
	results := make(chan Result, len(inputs))

	fmt.Println("Запускаем воркеров", len(inputs))
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func()  {
		fmt.Println("Начинаем отправку задач")
		for j := 1; j <= len(inputs); j++ {
			jobs <- Job{
				ID:     j,
				Number: j,
			}
			time.Sleep(100 * time.Millisecond)
		}
		close(jobs)
		fmt.Println("Все задачи отправлены")
	}()
	
	go func() {
		wg.Wait()
		close(results)
	}()
	
	for res := range results {
		fmt.Printf("Задача %d: %d^2 = %d, 4 * %d = %d (Считал воркер %d)\n",
			res.JobID, res.JobID, res.Square, res.JobID, res.Perimeter, res.WorkerID)
	}

	fmt.Println("Все задачи выполнены. Работа завершена.")
}