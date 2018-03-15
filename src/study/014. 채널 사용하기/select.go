package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() { // 계속 c1에 10을 보냄
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() { // 계속 c2에 "Hello world"를 보냄
		for {
			c2 <- "Hello world"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:						// 채널 c1에 값이 들어왔다면 꺼내서 i에 대입
				fmt.Println("c1:", i)
			case s := <-c2:
				fmt.Println("c2: ", s)		// 채널 c2에 값이 들어왔다면 꺼내서 s에 대입
			}
		}
	}()

	time.Sleep(5 * time.Second) // 10초 동안 프로그램 실행
}
