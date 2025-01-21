package main

import "fmt"

func main() {
	m := make(map[string]int)
	a := "sumishaammu"
	for i := 0; i < len(a); i++ {
		// fmt.Println(a[i]) --->()bytes
		// fmt.Println(string(a[i]))-->character words
		m[string(a[i])]++
	}
	fmt.Println(m)
}
