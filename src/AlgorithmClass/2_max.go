package main

import "fmt"

func main() {
	a := []int{3, 4, 6, 8, 1, 14}
	var max, index = 0, 0

	for i := 0; i < len(a) - 1; i++ {
		if a[i] < a[i + 1] {
			max = a[i+1]
			index = i+1
		} else {
			max = a[i]
			index = i
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Index:", index)
}
