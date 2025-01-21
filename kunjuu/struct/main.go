package main

import "fmt"

type person struct {
	name  string
	age   int
	study string
}

func main() {
	one := person{
		name:  "akhildas",
		age:   19,
		study: "go lang",
	}
	fmt.Println(one)

}
