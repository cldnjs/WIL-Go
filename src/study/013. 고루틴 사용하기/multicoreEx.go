package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())// CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정값 출력

	s := "Hello, World"

	for i := 0; i < 100; i++ { // 익명 함수를 고루틴으로 실행
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln() // main함수가 종료되지 않게 대기
}
