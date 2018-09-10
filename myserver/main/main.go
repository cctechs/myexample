package main

import (
	//"fmt"
	"../client"
	"fmt"
	"ivi.com/ucbase"
	"ivi.com/uclog"
	"ivi.com/ucnet"
	"net"
)

func TestPcap() {
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", 9998))
	if err != nil {
		panic(err)
	}

	go func() {
		conn, _ := ln.Accept()
		fmt.Println("accept connection.")
		go TestRecv(conn)
	}()

}

func TestRecv(conn net.Conn) {
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	if n > 0 {
		fmt.Println(n)
		fmt.Println(string(buf[:n]))
	}
}

func testListener() {
	lister := ucnet.NewTcpAcceptor()
	lister.StartListen(9998, func(sock *ucnet.TcpSocket) ucnet.TcpSocketInterface {
		clientSock := client.NewClientSocket(sock)
		client.GetClientMgr().AddSocket(clientSock)
		return clientSock
	})
}

func main() {

	ucbase.PrintSysInfo()

	uclog.Init("./log", uclog.DEBUG)

	//testListener()
	TestPcap()

	ucbase.Run()
}
