package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type File struct{
	FileName string
	Size int
}

func Source(fileNames []string) <- chan string{
	out := make(chan string)

	go func(){
		defer close(out)
		for i := 0; i < len(fileNames); i++{
			out <- fileNames[i]
		}
	}()
	return out
}

func filter(ch <-chan string) <- chan string{
	out := make(chan string)

	go func(){
		defer close(out)
		for fileName := range ch{
			if strings.HasSuffix(fileName, ".txt"){
				out <- fileName
			}
		}
	}()
	return out
}

func processing(in <- chan string) []File{
	results := []File{}
	var wg sync.WaitGroup
	wg.Add(3)
	go func(){
		defer wg.Done()
		d := time.Duration(rand.Intn(500))
		time.Sleep(d * time.Millisecond)
		for name := range in{
			results = append(results, File{FileName: name, Size: rand.Intn(100)}) 
		}
	}()
	go func(){
		defer wg.Done()
		d := time.Duration(rand.Intn(500))
		time.Sleep(d * time.Millisecond)
		for name := range in{
			results = append(results, File{FileName: name, Size: rand.Intn(100)}) 
		}
	}()
	go func(){
		defer wg.Done()
		d := time.Duration(rand.Intn(500))
		time.Sleep(d * time.Millisecond)
		for name := range in{
			results = append(results, File{FileName: name, Size: rand.Intn(100)}) 
		}
	}()

	wg.Wait()
	return results
}

func collector(results []File) <- chan File{
	out := make(chan File)
	var wg sync.WaitGroup
	summ := 0
	wg.Add(1)
	go func(){
		defer wg.Done()
		for _, val := range results{
			summ += val.Size
			out <- val
		}
		fmt.Println("Всего обработано строк:", summ)
	}()
	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}

func main(){
	filenames := []string{"data.txt", "image.png", "info.txt", "backup.zip", "ladno.txt"}

	c1 := Source(filenames)
	c2 := filter(c1)

	proc := processing(c2)
	collect := collector(proc)

	for result := range collect{
		fmt.Println(result)
	}
}