package main

import "fmt"

type Printer interface {
	print()
}

type MyInt int

func (i MyInt) print() {
	fmt.Println(i)
}

type Square struct {
	width int
	height int
}

func (sq Square) print() {
	fmt.Println(sq.width *  sq.height)
}

func main() {
	var i MyInt = 5
	var s Square

	var p Printer

	p = i
	p.print()

	p = s
	p.print()
}
