package main

import (
	"fmt"
	"math"
)

func main() {
	/* 부동소수점의 엡실론을 이용한 소수 계산 방법 */
	var a = 10.0
	 for i := 0; i < 10; i++ {
	 	a -= 0.1
	 }
	 fmt.Println(a)

	 const epsilon = 1e-14 // Go에서의 엡실론값
	 fmt.Println(math.Abs(a - 9.0) <= epsilon)
}
