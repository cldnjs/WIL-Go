package main

import "fmt"

type Rectangle struct { // 구조체 선언
	width int
	height int
}

func modifyAreaA(rect *Rectangle, n int)  { // 포인터를 매개변수로 받음 -> 값을 참조
	rect.width *= n
	rect.height *= n
}

func modifyAreaB(rect Rectangle, n int)  { // 구조체 인스턴스를 매개변수로 받음 -> 값을 복사
	rect.width *= n
	rect.height *= n
}

func main() {
	rt1 := Rectangle{10, 10}
	rt2 := Rectangle{10, 10}

	modifyAreaA(&rt1, 2) // 구조체 포인터를 사용하므로 변경 사항이 적용(값을 참조하기 때문)
	modifyAreaB(rt2, 2) // 구조체 인스턴스를 넘겨주면 변경 안됨(값을 복사하기때문)

	fmt.Println(rt1)
	fmt.Println(rt2)
}