package main

import "fmt"
import "time"

// 定义方法，因为or方法会有递归调用
var or func(...<-chan interface{}) <-chan interface{}

func main() {
	// 方法主体，将多个channel 合并成一个channel 返回
	or = func(chs ...<-chan interface{}) <-chan interface{} {
		// 因为有递归调用，所以需要一个跳出条件
		switch len(chs) {
		case 0:
			// 如果没有传参，则直接返回一个nil
			return nil
		case 1:
			//如果只有一个参数channel，则将该参数channel返回即可
			return chs[0]
		}
		// 初始化done通知channel
		orDone := make(chan interface{})
		go func() {
			// 完成后 关闭orChan
			defer close(orDone)
			switch len(chs) {
			case 2:
				select {
				case <-chs[0]:
				case <-chs[1]:
				}
			default:
				select {
				case <-chs[0]:
				case <-chs[1]:
				case <-chs[2]:
				// 将其余的channel参数递归传入到or方法中
				case <-or(append(chs[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
	doWork := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			// 规定时间过后，关闭channel
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	t := time.Now()
	defer func(time.Time) {
		// 打印方法执行时间
		fmt.Println(fmt.Sprintf("cost:%s", time.Since(t)))
	}(t)
	<-or(
		doWork(time.Second*5),
		doWork(time.Second*3),
		doWork(time.Minute*2),
		doWork(time.Hour*2),
	)

}
