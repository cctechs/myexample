package client

import (
	"../pb"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"ivi.com/ucbase"
	"ivi.com/uclog"
	"ivi.com/ucnet"
)

type ClientSocket struct {
	sock *ucnet.TcpSocket
}

func NewClientSocket(sock *ucnet.TcpSocket) *ClientSocket {
	clientSock := &ClientSocket{}
	clientSock.sock = sock
	return clientSock
}

func (this *ClientSocket) OnConnected() {
	uclog.LOG_NOTICE(uclog.FACILITITY_APP, "OnConnected")
}
func (this *ClientSocket) OnClosed() {
	uclog.LOG_NOTICE(uclog.FACILITITY_APP, "OnClosed")
}

func (this *ClientSocket) OnPacket(header ucnet.PACKER_HEADER, data []byte) {
	uclog.LOG_NOTICE(uclog.FACILITITY_APP, "OnPacket, cmd=%v", header.Cmd)
	switch header.Cmd {
	case ucnet.CMD_ALIVE:
		{
			pkt := ucnet.PK_KEEP_ALIVE{}
			ucbase.UnSerializeFromBytes(data, &pkt)
			uclog.LOG_INFO(uclog.FACILITITY_APP, "CMD_ALIVE")
		}
	case uint32(pb.PK_LOGIN_CMD):
		{
			pkt := pb.PK_LOGIN{}
			nHeaderSize := ucbase.StreamSizeof(header)
			if err := proto.Unmarshal(data[nHeaderSize:], &pkt); err != nil {
				panic(err.Error())
			}
			fmt.Println(pkt)
			fmt.Println(pkt.Name)

			retMsg := pb.PK_LOGIN_R{}
			retMsg.Code = 01
			data := ucnet.AppendHeader(int32(pb.PK_LOGIN_R_CMD), &retMsg)
			this.sock.SendData(data)
		}
	}
}

func (this *ClientSocket) GetPacketSize(header ucnet.PACKER_HEADER) (uint32, error) {
	fmt.Println("GetPacketSize, cmd=%0x", header.Cmd)
	valid := false
	switch header.Cmd {
	case ucnet.CMD_ALIVE:
		{
			if ucbase.StreamSizeof(ucnet.PK_KEEP_ALIVE{}) == header.Size {
				valid = true
			}
		}
		break
		// 针对pb协议，无法精确计算出size
	case uint32(pb.PK_LOGIN_CMD):
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
	uclog.LOG_NOTICE(uclog.FACILITITY_APP, "invalid packet size, cmd:%0x", header.Cmd)
	return 0, errors.New("get packet size error")
}
