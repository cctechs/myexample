package main

import (
	"context"
	"fmt"
	mm "github.com/micro/go-micro/v2"
	"myexample/mygomicro/greeter"
	pb "myexample/mygomicro/greeter"
)
type Greeter struct {

}

func (g Greeter) Hello(ctx context.Context, req *greeter.MyRequest, rsp *greeter.MyResponse) error {
	rsp.Greeting = "hello" + req.Name
	return nil
}

func main(){
	service := mm.NewService(mm.Name("greeter"))
	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}
