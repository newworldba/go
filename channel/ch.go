package main

import "fmt"

func main() {
	// channel 拥有者
	chOwner := func() <-chan interface{} {
		// 实例化channel
		ch := make(chan interface{})
		go func() {
			// 写入完成后，关闭channel
			defer close(ch)
			for i := 0; i < 10; i++ {
				// 拥有者负责写入
				ch <- i
			}
		}()
		// 将只读channel 返回
		return ch
	}
	chCon := func(ch <-chan interface{}) {
		// 只需要知道channel何时被关闭的即可
		for i := range ch {
			fmt.Println(i)
		}
	}

	ch := chOwner()
	chCon(ch)
}
