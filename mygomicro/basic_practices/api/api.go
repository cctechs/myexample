package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"log"
	pb "myexample/mygomicro/basic_practices/api/proto"
	"net/http"
	"strings"
)

type Example struct {

}

func (e *Example) Call(ctx context.Context, request *pb.Request, response *pb.Response) error {
	//log.Fatalf("example.call received a request.\n")
	fmt.Println("example call...")
	name, ok := request.Get["name"]
	if !ok || len(name.Values) ==0{
		return errors.BadRequest("go.micro.api.example", "params error")
	}

	for k, v := range request.Header{
		log.Printf("header: %v, %v \n", k, v)
	}

	response.StatusCode = http.StatusOK
	b, _ :=json.Marshal(map[string]string{
		"message":"your request" + strings.Join(name.Values, " "),
	})
	response.Body = string(b)
	return nil
}

type Foo struct {

}

func (f Foo) Bar(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Printf("foo.bar ..")
	if req.Method != "POST"{
		return errors.BadRequest("go.micro.api.example", "require post")
	}
	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) ==0{
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}

	if ct.Values[0] != "application/json"{
		return errors.BadRequest("go.micro.api.example", "expected application/json")
	}

	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)
	rsp.Body = "get a message" + string([]byte(req.Body))
	return nil
}

func main(){
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		)
	service.Init()
	pb.RegisterExampleHandler(service.Server(), new(Example))
	pb.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil{
		log.Fatal(err)
	}
}
