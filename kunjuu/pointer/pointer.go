package main

import "fmt"

func main() {
	x := 10
	pt := &x
	fmt.Println(x)   //value of x
	fmt.Println(*pt) //value of pointer

	*pt = 25         //changing value of pointer
	fmt.Println(*pt) //value of changing pointer

}
