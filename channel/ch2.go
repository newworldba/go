package main

import "fmt"
import "time"

func main() {
	fmt.Println("vim-go")
	var ch = make(chan int, 10)
	go cIf(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(2 * time.Second)
}

func cIf(ch chan int) {
	for {
		if v, ok := <-ch; ok {
			fmt.Println("output:", v)
		} else {
			break
		}
	}
}
