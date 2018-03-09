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

func (rect *Rectangle) area() int { // class가 없는 대신 구조체에 메소드 연결 가능, 리시버 변수는 java의 this와 비슷함
	return rect.height * rect.width // 포인터이므로 원래의 값이 변함
}

func (rect Rectangle) area2() int { // 인스턴스이므로 원래의 값 안변함(복사함)
	return rect.height * rect.width
}

func main() {
	rt1 := Rectangle{10, 10}
	rt2 := Rectangle{10, 10}

	modifyAreaA(&rt1, 2) // 구조체 포인터를 사용하므로 변경 사항이 적용(값을 참조하기 때문)
	modifyAreaB(rt2, 2) // 구조체 인스턴스를 넘겨주면 변경 안됨(값을 복사하기때문)

	fmt.Println(rt1)
	fmt.Println(rt2)

	fmt.Println(rt1.area())
	fmt.Println(rt1.area2())
}