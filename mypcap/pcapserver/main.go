package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9998")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		fmt.Println("new connection,", conn.RemoteAddr())
		if err != nil {
			panic(err)
		}

		go func(client net.Conn) {
			for {
				buf := make([]byte, 512)
				n, err := client.Read(buf)
				if err == nil {
					fmt.Println(client.RemoteAddr(), " data:", buf[:n])
				}
			}
		}(conn)
	}
}
