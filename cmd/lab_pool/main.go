package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Number int
}

type Result struct {
	JobID    int
	Value    int
	WorkerID int
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Воркер %d начал смену\n", id)

	for job := range jobs {
		res := factorial(job.Number)
		results <- Result{JobID: job.ID, Value: res, WorkerID: id}
	}

	fmt.Printf("Воркер %d закончил смену\n", id)
}

func main() {
	jobs := make(chan Job)
	var wg sync.WaitGroup

	numWorkers := 3
	numJobs := 20
	results := make(chan Result, 5)

	fmt.Println("Запускаем воркеров", numJobs)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		fmt.Println("Начинаем отправку задач")
		for j := 1; j <= numJobs; j++ {
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
		fmt.Printf("Задача %d: %d! = %d (Считал воркер %d)\n",
			res.JobID, res.JobID, res.Value, res.WorkerID)
	}

	fmt.Println("Все задачи выполнены. Работа завершена.")
}
