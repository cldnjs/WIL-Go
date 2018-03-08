package main

import "fmt"

func f() {
	defer func() {
		s := recover() // 패닉 발생 시에 넘어오는 메시지를 s로 받아서 프린트
		fmt.Println(s)
	}()

	a := []int{1, 2}
	for i := 0; i < len(a) + 1; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	f()
	fmt.Println("hello world")
}
