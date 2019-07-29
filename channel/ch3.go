package main

import "fmt"
import "time"

type Handler struct {
	Job chan int
}

func main() {
	fmt.Println("vim-go")

	var ch = &Handler{
		Job: make(chan int, 2),
	}
	for i := 0; i < 10; i++ {
		go ch.cHandle(i)
	}
	time.Sleep(time.Second * 2)

}

func (hd *Handler) cHandle(i int) {
	fmt.Println("func", i)
	select {
	case hd.Job <- i:
		fmt.Println("insert", i)
	case t := <-hd.Job:
		fmt.Println("read", t)
	}
}
