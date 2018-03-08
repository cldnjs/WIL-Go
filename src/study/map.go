package main

import "fmt"

func main() {
	a := map[string]int{
		"hello": 1,
		"world": 2,
		"coding": 3,
		"fun": 4,
	}

	fmt.Println(a["hello"])
	fmt.Println(a["world"])

	if _, ok := a["golang"]; ok { // 맵에서 데이터의 유효성 확인하는 if문
		fmt.Println("존재하는 키입니다")
	} else {
		fmt.Println("존재하지 않는 키입니다.")
	}

	for key, value := range a {
		fmt.Printf("%s: %d\n", key, value)
	}

	delete(a, "hello") // 맵에서 데이터 삭제
	fmt.Println(a)
}
