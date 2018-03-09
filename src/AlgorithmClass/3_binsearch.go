package main

import (
	"fmt"
	"sort"
)

func main() {
	var count int // 숫자의 개수
	var findNum int // 찾고자하는 값
	fmt.Scanln(&count)

	numList := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Scan(&numList[i])
	}
	sort.Sort(sort.IntSlice(numList))

	fmt.Scanln(&findNum)

	index := BinSearch(numList, findNum)

	fmt.Println("Index:", index)
}

func BinSearch(target []int, findValue int) int { // 이진탐색을 실행하는 함수
	startIndex := 0
	endIndex := len(target) - 1

	for (endIndex - startIndex) >= 0 {
		median := (startIndex + endIndex) / 2
		if target[median] == findValue {
			return median
		} else if target[median] < findValue {
			startIndex = median + 1
		} else {
			endIndex = median - 1
		}
	}

	return -1
}