package main

import (
	"fmt"
	"time"
)

func ammu() {
	fmt.Println("hello all ")
}
func main() {
	go ammu()
	time.Sleep(1 * time.Second)
	fmt.Println("akhil is a good boy")

}
