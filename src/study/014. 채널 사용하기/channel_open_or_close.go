package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	val, ok := <-c	// 채널이 닫혔는지 확인
	fmt.Println(val, ok)

	close(c)
	val, ok = <-c
	fmt.Println(val, ok)
}
