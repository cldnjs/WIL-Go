package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true 				// 고루틴에 true를 보냄, 값을 꺼낼 때까지 대기
			fmt.Println("고루틴:", i)// 반복문의 변수를 출력
			time.Sleep(1 * time.Second) // 1초 대기
		}
	}()

	for i := 0; i < count; i++ {		// 값을 꺼내지 않으면 무한 대기
		<-done							// 값을 꺼냄, 채널에 값이 들어올 때까지 대기
		fmt.Println("메인함수: ", i) // 반복문의 변수 출력
	}
}
