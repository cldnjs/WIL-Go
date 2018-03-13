package main

import "fmt"

func main() {
	Loop: // Loop 레이블을 지정
		//fmt.Println("들어가면 혼나") -> 레이블과 for문 사이에 다른게 들어가면 안됨
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if j == 3 {
					break Loop
				}

				fmt.Println(i, j)
			}
		}
		fmt.Println("Loop End")
		fmt.Println()

		for i := 0; i < 5; i++ {
			if i == 2 {
				continue // 해당 조건에서는 밑의 실행 코드를 수행하지 않고 넘어감
			}
			fmt.Println(i)
		}
}
