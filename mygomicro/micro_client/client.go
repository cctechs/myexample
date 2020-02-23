package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	pb "myexample/mygomicro/greeter"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	fmt.Printf("[wrapper] client request to service : %s method: %s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp, opts...)
}

func logWrap(c client.Client) client.Client{
	return &logWrapper{c}
}

func main()  {
	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.WrapClient(logWrap),
		)
	service.Init()

	greeter := pb.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.Background(), &pb.MyRequest{Name:"test111"})
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(rsp.Greeting)
}

