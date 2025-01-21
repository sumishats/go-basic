package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["akhil"] = 19
	m["ammu"] = 18
	fmt.Println(m)

	// value change in map
	m["akhil"] = 20
	fmt.Println(m["akhil"])

	// using loop on the map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// delete the key of map
	delete(m, "ammu")
	fmt.Println(m)

}
