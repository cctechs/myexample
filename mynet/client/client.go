package main

import (
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "")
	if err != nil{
		panic(err)
	}
	//defer conn.Close()
	conn.Write([]byte("hello world"))
	time.Sleep(time.Hour)
}
