package main

import (
	"fmt"
	"strconv"
)

type Overlap struct {
	name string
	age int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Overlap:
		p := arg.(Overlap)
		return p.name + " " + strconv.Itoa(p.age)
	case *Overlap:
		p := arg.(*Overlap)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(Overlap{"chiwon", 12}))
	fmt.Println(formatString(Overlap{"minji", 12}))

	var andrew = new(Overlap)
	andrew.name = "andrew"
	andrew.age = 12

	fmt.Println(formatString(andrew))
}
