package main

import (
	"context"
	"fmt"
	gm "github.com/micro/go-micro/v2"
	pb "myexample/mygomicro/greeter"
)

func main()  {
	service := gm.NewService(gm.Name("greeter.client"))
	service.Init()

	greeter := pb.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.Background(), &pb.MyRequest{Name:"test111"})
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(rsp.Greeting)
}

