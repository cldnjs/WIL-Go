package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // main()이 끝나기 직전에 사용
	hello()

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // Last in First out
	}
}
