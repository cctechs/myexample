package utils

import (
	"fmt"
	"strings"
)

// MyFunc
type MyFunc struct {
}

// this is Add
func MyAdd(a, b int32) int32 {
	return a + b
}

// this is MyString
func MyString(str string) int32 {
	arr := strings.Split(str, ";")
	return int32(len(arr))
}

func Fib() func(int) int {
	sum := 0
	return func(i int) int {
		sum = sum + i
		return sum
	}
}

func MyFib() {
	f := Fib()
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(5))
}
