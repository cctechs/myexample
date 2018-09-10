package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	fmt.Println("hello world")
	runtime.Version()
	fmt.Println(runtime.Version())

	fmt.Println(math.MaxFloat32)

	fmt.Println(fmt.Sprintf("%v", math.MaxFloat32))
}
