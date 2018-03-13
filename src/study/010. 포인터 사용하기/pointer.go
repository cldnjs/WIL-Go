package main

import "fmt"

func main() {
	var num int = 1
	var numPt *int = &num

	fmt.Println(numPt)
	fmt.Println(&num)
	fmt.Println(*numPt)
}