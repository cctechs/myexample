package main

import (
	"log"
	"myexample/myrpc/server"
	"net/rpc"
)

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil{
		log.Fatalf("dialing:", err)
	}

	for i:= 1; i <100; i++{
		args := &server.Args{i,i+1}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil{
			log.Fatalf("arith error: %v", err)
		}
		log.Printf("result:%v\n", reply)
	}



}
