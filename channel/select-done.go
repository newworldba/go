package main

import "fmt"
import "time"

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()
loop:
	for {
		select {
		case <-done:
			fmt.Println("done...")
			break loop
		default:
			fmt.Println("do something1...")
		}
		time.Sleep(time.Second)
		fmt.Println("do something2...")
	}
}
