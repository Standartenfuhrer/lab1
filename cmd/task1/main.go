package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
   if len(os.Args) != 2{
    log.Fatal("Ожидается аргумент: количество горутин")
   }
   countGoroutine, err := strconv.Atoi(os.Args[1])
   if err != nil {
    fmt.Println("Аргумент должен быть положительным числом")
   }
   for i := 0; i < countGoroutine; i++{
    
   }
}
