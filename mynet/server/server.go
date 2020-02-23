package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":0")
	if err != nil{
		log.Fatal(err)
	}


	fmt.Println(listen.Addr())

	go func() {
		for{
			client, err := listen.Accept()
			if err != nil{
				panic(err)
			}

			buf := make([]byte, 1024)
			n, err := client.Read(buf)
			if err != nil{

			}
			fmt.Printf("%v recv:%s", client.RemoteAddr(), string(buf[:n]))
		}
	}()

    time.Sleep(time.Hour)
}
