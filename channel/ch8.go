package main

import "fmt"

func main() {
	ch := make(chan string, 10)
	fmt.Println(<-ch)
	close(ch)

}
