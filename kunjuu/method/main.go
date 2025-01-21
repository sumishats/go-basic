package main

import "fmt"

func main() {
	m := make(map[string]int)

	var count string
	m[count]++
	fmt.Println(m)
}

// type person struct {
// 	name string
// 	age  int
// }

// func (ammu *person) gone() {

// 	ammu.age = 34
// }

// func run(akhil *person) {

// 	akhil.age = 19
// }
// func main() {

// 	a := person{age: 30, name: "myraan akhil"}

// 	// normal function
// 	run(&a)
// 	fmt.Println(a)

// 	// method
// 	a.gone()

// 	fmt.Println(a)

// }
