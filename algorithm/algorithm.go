package main

import (
	"fmt"
	"sort"
)

type mystrings []string

func (s mystrings) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s mystrings) Len() int  {
	return len(s)
}

func (s mystrings) Less(i, j int) bool{
	if s[i] < s[j]{
		return true
	}
	return false
}

//0 1 1 3 5 8
//0 1 2 3 4 5 6
var m int  = 0
func fib(n int) int{
	if n == 0{
		return 0
	}

	if n == 1{
		return 1
	}
	k := fib(n - 1) + fib(n - 2)
	return k
}

func main(){
	s := mystrings{"z12", "asd", "acd"}
	sort.Sort(s)
	fmt.Println(s)
	n := sort.Search(len(s), func(i int) bool {
		if s[i] == "ac"{
			return true
		}
		return false
	})

	fmt.Println(n)
	fmt.Println("hello world")

	for i := 0; i < 10; i++{
		fmt.Println(fib(i))
	}
}
