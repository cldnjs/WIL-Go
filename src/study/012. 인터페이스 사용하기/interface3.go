package main

import "fmt"

type chiwon interface { // 인터페이스
	say()
}

type test struct { // chiwon을 구현한 struct
	name int
	what float64
}

func (rcv test) say() { // say() 구현
	fmt.Println("hello")
}

func main() {
	a := test{1, 1.0}

	if val, ok := interface{}(a).(chiwon); ok { // 특정 인터페이스를 구현하는지 확인하는 코드
		fmt.Println(val, ok)
	}
}
