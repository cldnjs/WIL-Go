package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ { // 100번 반복하여 고루틴 100개 생성
		go hello(i)
	}

	fmt.Scanln() // main함수가 끝나지 않도록 대기
}
