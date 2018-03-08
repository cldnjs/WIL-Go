package main

import "fmt"

func main() {
	/* 슬라이스의 특징 */
	a := []int{1, 2, 3, 4}
	var b []int

	b = a
	a[0] = 10 // 복사가 아니라 참조기때문에 둘 중 하나를 변경하면 둘 다 영향을 받음

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---------------------------------")

	/* 슬라이스 복사 */
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := make([]int, 3)

	copy(sliceB, sliceA) // 슬라이스를 복사할 때는 copy메소드 사용

	fmt.Println(sliceA)
	fmt.Println(sliceB)
	fmt.Println("---------------------------------")

	/* 부분 슬라이스 */
	fmt.Println(a[0:2])
	fmt.Println(a[2:])
}
