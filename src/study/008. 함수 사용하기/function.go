package main

import "fmt"

func SumAndDiff(a int, b int) (int, int) { // Go는 여러개의 리턴값을 가질 수 있다
	return a + b, a - b
}

func total(n ...int) int { // 가변인자를 사용한 함수
	total := 0
	for _, val := range n {
		total += val
	}

	return total
}

func main() {
	sum, diff := SumAndDiff(3, 4)
	fmt.Printf("Sum: %d\nDiff: %d\n", sum, diff)

	total := total(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}
