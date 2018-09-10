package main

/*
#include <stdio.h>
int Add(int a, int b)
{
	return a + b;
}

int Sub(int a, int b)
{
	return a - b;
}

*/
import "C"

import (
	"fmt"
)

func main() {
	//cs := C.CString("hello")
	fmt.Println(C.Add(1, 2))
	fmt.Println(C.Sub(1, 2))
	//C.free(unsafe.Pointer(cs))
}
