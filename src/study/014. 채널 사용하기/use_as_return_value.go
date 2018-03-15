package main

import "fmt"

func num(a, b int) <-chan int {
	out := make(chan int)

	go func() {
		out <- a
		out <- b
		close(out)
	}()
	return out
}

func add(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		total := 0
		for i := range c{
			total += i
		}
		out <- total
	}()

	return out
}

func main() {
	a := num(1, 2) // 1과 2가 들어 있는 채널이 리턴됨
	sum := add(a) // 채널 a를 매개변수에 담아서 모두 더함

	fmt.Println(<-sum)
}
