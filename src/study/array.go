package main

import "fmt"

func main() {
	/* range를 이용한 배열 사용 */
	arr := [5]int{10, 20, 30, 40, 50}

	for index, val := range arr { // index에는 배열 번호, val에는 값이 들어감
		fmt.Printf("index: %d\n", index)
		fmt.Printf("value: %d\n", val )
		fmt.Println()
	}

	fmt.Println("--------------------------------")

	/* 2차원 배열 */
	a := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 왼쪽의 형식으로 배열 선언시 마지막 항목에 ,을 써줘야함
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("[%d][%d] : %d\n", i, j, a[i][j])
		}
	}
}
