package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "myexample/mymicro/helloworld"
	"net"
)

const (
	port = ":8801"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	//panic("implement me")
	log.Printf("received:%v", req.GetName())
	return &pb.HelloReply{Message:"Hello" + req.GetName()}, nil
}

func main(){
	listen, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("start listen...")
	s:= grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil{
		log.Fatalf("failed to serve:%v", err)
	}
}
