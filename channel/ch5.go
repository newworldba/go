package main

import "fmt"
import "sync"

func main() {
	inCh := gen(10)
	outCh := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go do(inCh, outCh, &wg)
	}
	go func() {
		wg.Wait()
		close(outCh)
	}()
	for v := range outCh {
		fmt.Println("read", v)
	}
}

func gen(s int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < s; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func do(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range inCh {
		outCh <- v * v
	}
}
