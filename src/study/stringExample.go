package main

import (
	"fmt"
	"unicode/utf8"
	"runtime"
)

func main() {
	var s1 = "한글입니다"
	var s2 = "hello"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s1)) // 한글의 경우 실제 문자열의 길이 구함
	fmt.Println(runtime.NumCPU(), "개")
}
