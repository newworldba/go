package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func main() {
	fmt.Println("vim-go")
	var ch = make(chan int, 2)
	wg.Add(10)
	go cFor(ch)
	for i := 0; i < 10; i++ {
		fmt.Println("income:", i)
		ch <- i
	}
	close(ch)
	wg.Wait()

}

func cFor(ch chan int) {
	for v := range ch {
		fmt.Println("output:", v)
		wg.Done()
	}
}
