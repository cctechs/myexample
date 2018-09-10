package main

import "fmt"

// "正统" 函数式编程 not golang
// 不可变性 不可能有状态 只有常量和函数
// 函数只能有一个参数
// golang 不做上述规定

// 闭包 -> { 自由变量 + 函数体{局部变量} }
//

func adder() func(value int)int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder  {
	fmt.Println("base:", base)
	return func(v int) (int, iAdder) {
		fmt.Printf("base:%v, v:%v\n", base, v)
		return base + v, adder2(base + v)
	}
}


func main() {
	/*
	f := adder()
	for i:= 0; i < 10; i++{
		//fmt.Println(f(i))
		fmt.Printf("0+...%d=%d\n", i, f(i))
	}
	*/
	{
		a := adder2(3)
		_, a = a(5)
		a(10)
		return
	}

	a := adder2(0)
	fmt.Println("111111111")
	for i :=0; i < 1;i++{
		var s int
		s, a = a(i)
		fmt.Printf("0+...%d=%d\n", i, s)
	}
}
