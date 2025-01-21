package main

import "fmt"

func main() {
	m := make(map[string]int)
	s := "hlo im hlo ammu im "
	result := ""
	for i := 0; i < len(s); i++ {

		if string(s[i]) == " " {

			m[result]++
			result = ""

		} else { //-->store the words
			result += string(s[i])
		}
	}

	fmt.Println(m)

}
