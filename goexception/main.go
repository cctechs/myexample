package main

import (
	"fmt"
	"ivi.com/ucbase"
	"os"
)

func MyDefer() {
	defer func() {
		fmt.Println("mydefer")
	}()
	fmt.Println("qe")
}

func MyPanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("here")
			MyPanic()
		}
	}()

	panic("myfault")
	//here:
	fmt.Println("qqqqq")
}

func main() {

	os.Chdir("./goexception")

	//	MyPanic()

	ucbase.PProf("test.prof")
}
