package main

import "fmt"

func main() {

	m := make(chan string) //unbuffered

	go func() {
		m <- "ammu"

	}()

	v := <-m

	fmt.Println(v)

	s := make(chan string, 3) //buffered

	s <- "hi"
	s <- "hlo"
	s <- "hy"
	close(s)
	for value := range s {
		fmt.Println(value)
	}

}
