package main

import (
	"log"
	"myexample/myrpc/server"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main(){
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil{
		log.Fatalf("listen error:", err)
	}
	go http.Serve(l, nil)
	time.Sleep(time.Hour)
}
