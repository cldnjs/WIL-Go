package main

import (
	"fmt"
)


func main() {
	var input int64

	fmt.Println("1 ~ 5사이의 숫자를 입력하세요")
	fmt.Scanln(&input)

	switch input {
	case 1:
		fmt.Println("1입니다")
	case 2:
		fmt.Println("2입니다")
	case 3:
		fmt.Println("3입니다")
	case 4:
		fmt.Println("4입니다")
	case 5:
		fmt.Println("5입니다")
		// fallthrough -> 다음 case문을 실행할 수 있음. 맨 마지막에는 못쓴다.
	default:
		fmt.Println("범위 밖의 숫자입니다")
	}
}
