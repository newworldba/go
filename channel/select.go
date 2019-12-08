package main

import "fmt"

func main() {
	c1, c2 := make(chan interface{}), make(chan interface{})
	close(c1)
	close(c2)
	var countC1, countC2 int64
	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			countC1++
		case <-c2:
			countC2++
		}
	}
	fmt.Println(fmt.Sprintf("countC1:%d||countC2:%d", countC1, countC2))
	// countC1:484||countC2:516
}
