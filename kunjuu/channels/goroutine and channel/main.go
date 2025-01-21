package main

import "fmt"

func getrank(x map[string]int) map[string]int {
	m := make(map[string]int)

	for k, v := range x {
		if v == 1 {
			m[k] = v
		}
	}
	return m

}

func main() {
	s := make(map[string]int)
	s["ammu"] = 1
	s["gadhu"] = 2
	s["akhil"] = 3
	s["sumisha"] = 1
	q := getrank(s)
	fmt.Println(q)

}

// type bank struct {
// 	name    string
// 	balance int
// }

// func (a *bank) addmoney(x int) {
// 	a.balance += x

// }

// func main() {

// 	a := bank{name: "ammu", balance: 10}
// 	a.addmoney(2045)
// 	fmt.Println(a)

// }
