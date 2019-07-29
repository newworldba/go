package main

import "fmt"
import "time"
import "errors"

func main() {
	i, err := doWithTime(time.Second * 1)
	i, err = doWithTime(time.Second * 3)
	fmt.Println(i, err)
}

func doWithTime(t time.Duration) (int, error) {
	select {
	case i := <-do():
		return i, nil
	case <-time.After(t):
		fmt.Println("timeout", t)
		return 0, errors.New("timeout")
	}
}

func do() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- 1
	}()

	return ch
}
