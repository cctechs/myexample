package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"

	"bufio"
	"errors"
	"os"
	//"net/http"
	"./utils"
)

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func consts() {
	const (
		b = 1 << (10 * iota)
		kb
		mb
	)
	fmt.Println(b, kb, mb)
}

func converToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		result = fmt.Sprintf("%d", n%2) + result
	}
	return result
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "C"

	}
	return g
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func slices() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	s1 := arr[2:6]
	fmt.Println(s1)
	s2 := s1[3:5]
	fmt.Println(s2)
	fmt.Println(cap(s1))
	s3 := arr[:]
	fmt.Println(s3)
	s3 = append(s3, 11)
	fmt.Println(s3)

}

func convertToValue(str string) int64 {
	var val int64
	count := len(str)
	for i := 0; i < len(str); i++ {
		if v, err := strconv.Atoi(string(str[i])); err != err {
			break
		} else {
			val += int64(v << uint(count-i-1))
		}
	}
	return val
}

type TreeNode struct {
	value int32
}

func (node TreeNode) UpdateValue() {
	node.value = 10
}

func (node *TreeNode) uuuu() {
	node.value = 11
}

type queue []int

func (this *queue) push(v int) {
	*this = append(*this, v)
}

func (this *queue) pop() (int, error) {
	if this.empty() {
		return -1, errors.New("empty")
	}
	head := (*this)[0]
	*this = (*this)[1:]
	return head, nil
}

func (this *queue) empty() bool {
	return 0 == len(*this)
}

type Retriver interface {
	Get() string
}

type MyRetriver struct {
}

func (this *MyRetriver) Get() string {
	return "MyRetriver"
}

type YouRetriver struct {
}

func (obj YouRetriver) Get() string {
	return "YouRetriver"
}

func inspect(r Retriver) {
	switch v := r.(type) {
	case *MyRetriver:
		{
			fmt.Println(v.Get())
		}
	case YouRetriver:
		{
			fmt.Println("%T", v)
		}
	case *YouRetriver:
		{
			fmt.Println(v.Get())
		}
	}
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	//panic("qqweq")
	fmt.Println(3)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	writer.Write([]byte("str"))
}

type myFunc func() int

func (fn *myFunc) Test() {
	//	fmt.Println(fn())
}

func TestMyFunc() {

}

func MyHttp() {
	//http.Handle("/list/", )
	//http.StatusSeeOther
}

func TryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err)
		} else {
			panic(err)
		}

		//if err := recover(); ok{

		//}
	}()
	//panic("12345")
	//b := 0
	//a := 5/b
	//fmt.Println(a)
	panic("123")
}

func main() {
	{
		utils.MyFib()
		return
	}

	TryRecover()

	writeFile("123.txt")
	tryDefer()
	{
		inspect(&MyRetriver{})
		inspect(YouRetriver{})
	}

	fmt.Println("hello world")
	//euler()
	//consts()
	fmt.Println(converToBinary(10123))
	fmt.Println(convertToValue("10011110001011"))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	slices()

	node := TreeNode{1}
	fmt.Println(node)
	node.UpdateValue()
	fmt.Println(node)
	node.uuuu()
	fmt.Println(node)

	q := queue{}
	q.push(1)
	//q.push(2)
	fmt.Println("q:", q)
	fmt.Println(q.pop())
}
