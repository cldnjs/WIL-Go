package main

import "fmt"

type Person struct {
	name string
	age int
}

func (_ Person) greeting() {
	fmt.Println("hello~")
}

type Student1 struct {
	p Person	// 구조체가 Person타입을 가지고 있다(has a 관계)
	school string
	grade int
}

type Student2 struct {
	Person	// 필드명 없이 타입만 선언하면 포함 관계가 됨(is a)
	school string
	grade int
}

func main() {
	var student Student1
	student.p.greeting() // 가지고 있는것이므로 p를 통해서 접근

	var student2 Student2
	student2.greeting() // 바로 호출 가능


}
