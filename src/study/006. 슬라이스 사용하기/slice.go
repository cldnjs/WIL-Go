package main

import "fmt"

func main() {
	a := make([]int, 5) // 0으로 채워진 슬라이스 생성
	b := []int{10, 20, 30, 40} // 선언과 동시에 슬라이스 초기화

	for i := 0; i < 5; i++ {
		a = append(a, i+1) // a슬라이스 맨 뒤에서부터 추가(0 0 0 0 0 1 2 3 4 5)
	}

	fmt.Println(a)
	fmt.Println(b)
}
