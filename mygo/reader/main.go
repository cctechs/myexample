package main

import (
	"bufio"
	"fmt"
)

type MyReader struct {
	str string
}

func (MyReader) Read(p []byte) (n int, err error) {
	panic("implement me")
}


type Retriver interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster){
	poster.Post("1235", map[string]string{
		"1":"1222",
	})
}


type RetriverPoster interface {
	Retriver
	Poster
}

func session(s RetriverPoster){
	//s.Post()
}


func main() {
	mr := &MyReader{str:"12345"}
	b := bufio.NewReader(mr)
	fmt.Println(b.Peek(10))
}
