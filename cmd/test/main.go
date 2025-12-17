package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil{
		log.Fatal(err)
	}
	for i := 0; i < num; i++{
		for j := 0; j < num; j++{
			if (i + j) % 2 == 0{
				fmt.Print("x")
			} else{
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}