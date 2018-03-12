package main

import "fmt"

type People struct {
	name string
	age int
}

func (_ People) greeting() {
	fmt.Println("People method")
}

type Chiwon struct {
	People
	hobby string
}

func (rcv Chiwon) greeting() { // 부모 구조체의 메소드 오버라이딩
	fmt.Println("Chiwon method")
}

func main() {
	var people People
	var chiwon Chiwon

	people.greeting() // People method
	chiwon.People.greeting() // People method
	chiwon.greeting() // Chiwon method
}
