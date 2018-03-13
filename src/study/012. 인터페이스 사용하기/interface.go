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

func show(p... Printer)  {
	for index := range p {
		p[index].print()
	}
}

func main() {
	var i MyInt = 5
	s := Square{10, 20}

	show(i, s)
}
