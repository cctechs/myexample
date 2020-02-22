package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "myexample/mymicro/helloworld"
	"time"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:8801", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("dial error, %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := "world"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	start := time.Now().UnixNano()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name:name})
	if err != nil{
		log.Fatalf("greet failed %v", err)
	}
	fmt.Println("cost:", (time.Now().UnixNano() - start)/1000.0)
	log.Printf("greentins:%s", r.GetMessage())

}
