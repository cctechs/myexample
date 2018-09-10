package main

import (
	"../pb"
	"github.com/golang/protobuf/proto"
	"ivi.com/ucbase"
	"ivi.com/uclog"
	"ivi.com/ucnet"
	"time"
)

type Server struct {
	connector ucnet.TcpConnector
	connected bool
	addr      string
}

func (this *Server) Start(addr string) {
	this.addr = addr

	// 在协程中定时检测连接状态
	checker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-checker.C:
				{
					this.CheckStatus()
				}
			}
		}
	}()
}

func (this *Server) sendKeepAlive() {
	// 发送心跳
	pkt := ucnet.PK_KEEP_ALIVE{}
	pkt.Cmd = ucnet.CMD_ALIVE
	pkt.Size = ucbase.StreamSizeof(pkt)
	data := ucbase.SerializeToBytes(&pkt)
	this.connector.SendData(data)
	uclog.LOG_INFO(uclog.FACILITITY_APP, "sendKeepAlive")
}

func (this *Server) CheckStatus() {
	if this.connected {
		//	this.sendKeepAlive()
	} else {
		// 连接
		this.connector.Connect(this.addr, this)
	}
}

func (this *Server) sendLogin() {
	pkt := pb.PK_LOGIN{}
	pkt.Name = "test001"
	data := ucnet.AppendHeader(int32(pb.PK_LOGIN_CMD), &pkt)
	this.connector.SendData(data)
	uclog.LOG_INFO(uclog.FACILITITY_APP, "sendLogin, data:%v", data)
}

func (this *Server) OnConnected() {
	this.connected = true
	uclog.LOG_INFO(uclog.FACILITITY_APP, "onconnected")

	// 连接成功后发送登录协议
	this.sendLogin()
}

func (this *Server) OnClosed() {
	uclog.LOG_INFO(uclog.FACILITITY_APP, "OnClosed")
	this.connected = false
}

// 连接关闭
func (this *Server) OnPacket(header ucnet.PACKER_HEADER, data []byte) {
	uclog.LOG_INFO(uclog.FACILITITY_APP, "OnPacket")

	switch header.Cmd {
	case uint32(pb.PK_LOGIN_R_CMD):
		{
			pkt := pb.PK_LOGIN_R{}
			nHeaderSize := ucbase.StreamSizeof(header)
			if err := proto.Unmarshal(data[nHeaderSize:], &pkt); err != nil {
				panic(err.Error())
			}
			uclog.LOG_INFO(uclog.FACILITITY_APP, "login ret:%v", pkt.Code)
		}
	}
}

func (this *Server) GetPacketSize(header ucnet.PACKER_HEADER) (uint32, error) {
	valid := false

	switch header.Cmd {
	case uint32(pb.PK_LOGIN_R_CMD):
		{
			if header.Size > ucbase.StreamSizeof(header) {
				valid = true
			}

		}
	default:
		break
	}

	if valid {
		return header.Size, nil
	}
	uclog.LOG_ERROR(uclog.FACILITITY_APP, "invalid packet, cmd=%0x", header.Cmd)
	return 0, nil
}

func main() {
	ucbase.PrintSysInfo()
	uclog.Init("./log", uclog.DEBUG)

	// start logic
	svr := &Server{}
	svr.Start("127.0.0.1:9998")

	ucbase.Run()
}
