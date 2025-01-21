package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 10; j++ {
				fmt.Println(j)
			}
			wg.Done()
		}()

	}

	wg.Wait()

}
