package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}
	ch := make(chan int)
	var sum int
	var sum2 int

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 0; i < len(arr)/2; i++ {
			sum += arr[i]

		}
		ch <- sum
		wg.Done()

	}()

	go func() {
		for j := len(arr) / 2; j < len(arr); j++ {
			sum2 += arr[j]
		}
		ch <- sum2
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var total int
	for value := range ch { //resecive channel
		fmt.Println(value)
		total += value

	}

	fmt.Println(total)

}
