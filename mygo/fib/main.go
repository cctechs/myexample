package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)

// the bigger the interface
// the weaker the abstraction


func fibnonacci() intGen{
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrent if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


func main() {
	/*
	f := fibnonacci()
	for i := 0; i < 10; i++{
		fmt.Printf("%d\n", f())
	}
	*/

	f:= fibnonacci()
	printFileContents(f)
}
