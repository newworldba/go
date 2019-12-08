package main

import "fmt"
import "time"

func main() {
	doWork := func(done, ch <-chan interface{}) <-chan interface{} {
		chRst := make(chan interface{})
		go func() {
			defer close(chRst)
			for {
				select {
				case <-done:
					// 如果done关闭，则退出该goroutine
					return
				case s := <-ch:
					fmt.Println(s)
				}
			}
			fmt.Println("goroutine exit")
		}()
		return chRst
	}
	done := make(chan interface{})
	chRst := doWork(done, nil)
	go func() {
		time.Sleep(time.Second * 2)
		close(done)
	}()
	fmt.Println(<-chRst)
}
