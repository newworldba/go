package main

import "fmt"
import "time"

func main() {
	ch := gen(10)
	consume(ch)
	time.Sleep(time.Second)
}

func gen(s int) <-chan int {
	ch := make(chan int, s)
	go func() {
		for i := 0; i < s*2; i++ {
			ch <- i
			fmt.Println("gen", i)
		}
		close(ch)
	}()
	return ch
}
func consume(ch <-chan int) {
	for v := range ch {
		fmt.Println("read", v)
	}
}
