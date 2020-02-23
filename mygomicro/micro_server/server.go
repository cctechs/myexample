package main

import (
	"context"
	"fmt"
	mm "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	"log"
	"myexample/mygomicro/greeter"
	pb "myexample/mygomicro/greeter"
	"time"
)
type Greeter struct {

}

func (g Greeter) Hello(ctx context.Context, req *greeter.MyRequest, rsp *greeter.MyResponse) error {
	rsp.Greeting = "hello" + req.Name
	fmt.Println("Hello->", rsp.GetGreeting())
	return nil
}

func logWrapper(fn server.HandlerFunc) server.HandlerFunc{
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request:%s", time.Now(), req.Endpoint())
		return fn(ctx, req, rsp)
	}
}

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	service := mm.NewService(
		mm.Name("greeter"),
		mm.WrapHandler(logWrapper),
		)
	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil{
		fmt.Println(err)
	}
}
