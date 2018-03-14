package main

import (
	"runtime"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, World"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 반복문에서의 변수를 클로저에서 바로 출력 -> 100이 계속 출력됨
		}()		              // 고루틴은 반복문이 끝나고 생성되므로 고루틴이 생성된 시점에서 i = 100
	}

	for i := 0; i < 100; i++ { // 클로저를 고루틴으로 실행할 때 반복문에 의한 변수는 반드시 매개변수로 넘기기
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
