package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)			// 버퍼를 2개 이상 설정하면 비동기 채널
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true				// 채널에 true를 보냄, 버퍼가 가득 차면 대기
			fmt.Println("고루틴:", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done							// 값을 꺼냄, 버퍼에 값이 없으면 대기
		fmt.Println("메인함수:", i)
	}
}
