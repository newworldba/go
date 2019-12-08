package main

import "fmt"

func main() {
	doWork := func(ch <-chan interface{}) <-chan interface{} {
		chRst := make(chan interface{})
		go func() {
			defer close(chRst)
			// 如果ch为nil， 则会一直阻塞,该方法不会完成
			for i := range ch {
				fmt.Println(i)
			}
		}()
		return chRst
	}
	chRst := doWork(nil)
	fmt.Println(<-chRst)
}
