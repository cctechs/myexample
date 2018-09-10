package main

import (
	"fmt"
	"net"
)

func main() {
	netAddr, _ := net.ResolveIPAddr("ipv4", "127.0.0.1")
	conn, err := net.ListenIP("ip4:icmp", netAddr)
	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		fmt.Println(addr.String())
		fmt.Println(net.InterfaceByName("en0"))
	}

}
