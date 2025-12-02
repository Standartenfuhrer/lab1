package main
import "fmt"

func main() {
    ch := make(chan int)
    go func() {
        // Хотим отправить 3 числа
        ch <- 1
        ch <- 2
        ch <- 3
        close(ch) 
    }()

    // Пытаемся прочитать больше, чем отправили
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)

	v, ok := <- ch
	if !ok{
		fmt.Println("Канал закрыт.")
	}
    fmt.Println(v) // <-- Что тут произойдет и почему? Исправьте, используя v, ok
}
