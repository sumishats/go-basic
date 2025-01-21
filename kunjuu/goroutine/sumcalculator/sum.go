package main

import "fmt"

func sum(arr []int, ch chan int) {
	var sum int
	for i := 0; i < len(arr)/2; i++ {
		sum = sum + arr[i]

	}
	for j := len(arr) / 2; j < len(arr); j++ {
		sum = sum + arr[j]
	}

	ch <- sum
	close(ch)
}
func main() {
	ammu := make(chan int)
	array := []int{10, 38, 28, 29, 17, 48}

	go sum(array, ammu)

	for val := range ammu {

		fmt.Println(val)
	}

}
