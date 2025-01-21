package main

import "fmt"

//	type animals interface {
//		speak() string
//	}
type cat struct {
}

func (c cat) speak() string {
	return "meoww meoww"
}

type dog struct {
}

func (d dog) speak() string {
	return "boww bhoww"
}

type frog struct {
}

func (f frog) speak() string {
	return "peckroww peckroww"
}

func main() {

	c := cat{}

	fmt.Println(c.speak())
	d := dog{}

	fmt.Println(d.speak())

	// d := dog{}

	// c := cat{}
	// var a animals

	// a = d

	// fmt.Println(a.speak())

	// a = c
	// fmt.Println(a.speak())

	// akhil := []animals{cat{}, dog{}, frog{}}
	// for _, v := range akhil {
	// 	fmt.Println(v.speak())
	// }

}
