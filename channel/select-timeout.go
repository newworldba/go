package main

import "fmt"
import "time"

func main() {
	ch := make(chan interface{})
	select {
	case <-ch:
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}
