package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보내고 대기
		}
		close(c)   // 채널을 닫음
	}()

	for i := range c { // range를 통해서 채널에서 값을 꺼낸 후 대기
		fmt.Println(i) // 꺼낸 값을 출력
	}
}
