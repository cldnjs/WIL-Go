package main

import "fmt"

func main() {
	var count int
	fmt.Scanln(&count)

	for i := 1; i <= count; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("X")
		default:
			fmt.Println(i)
		}
	}
}