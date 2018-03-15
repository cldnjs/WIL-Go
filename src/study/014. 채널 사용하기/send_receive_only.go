package main

import "fmt"

func producer(c chan<- int) { //보내기 전용 채널 생성
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100

	// fmt.Println(<-c)하면 보내기 전용 채널이므로 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널 생성
	for i := range c {
		fmt.Println(i)
	}

	//c <- 1을 하면 받기 전용 채널이므로 컴파일 에러
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
