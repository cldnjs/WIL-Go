package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

//Rect 정의
type Rect struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/*
인터페이스를 사용하는 일반적인 예로 함수가 파라미터로 인터페이스를 받아들이는 경우를 들 수 있다.
함수 파라미터가 interface 인 경우, 이는 어떤 타입이든 해당 인터페이스를 구현하기만 하면
모두 입력 파라미터로 사용될 수 있다는 것을 의미한다.
*/
func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() //인터페이스 메서드 호출
		fmt.Println(a)
	}
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}
