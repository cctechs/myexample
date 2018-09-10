package main

import (
	"fmt"
)

func main() {
	var c1, c2 chan int
	select {
	case n := <- c1:
		fmt.Println(n)
	case n := <- c2:
		fmt.Println(n)
	default:
		fmt.Println("no data received")
	}
}
