package main

import (
	"fmt"
)

func main() {

	days := 2
	switch days {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	default:
		fmt.Println("not found")
	}

}
