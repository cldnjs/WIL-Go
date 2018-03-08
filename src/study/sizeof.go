package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 int8 = 1
	var n2 int16 = 1
	var n3 int32 = 1
	var n4 int64 = 1

	fmt.Println(unsafe.Sizeof(n1))
	fmt.Println(unsafe.Sizeof(n2))
	fmt.Println(unsafe.Sizeof(n3))
	fmt.Println(unsafe.Sizeof(n4))
}
