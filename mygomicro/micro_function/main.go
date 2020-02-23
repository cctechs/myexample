package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	pb "myexample/mygomicro/greeter"
)

type Greeter struct {

}

func (g Greeter) Hello(ctx context.Context, req *pb.MyRequest, rsp *pb.MyResponse) error {
	//panic("implement me")
	rsp.Greeting = "hello" + req.GetName()
	fmt.Println("Hello->", rsp.GetGreeting())
	return nil
}

func main(){
	fn := micro.NewFunction(micro.Name("greeter"))
	fn.Init()
	fn.Handle(new(Greeter))
	fn.Run()
}
