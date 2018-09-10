package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-t.C:
			{
				bytes := []byte{1, 2, 3}
				fmt.Println("send:", bytes)
				conn.Write(bytes)
			}
		}
	}

}
