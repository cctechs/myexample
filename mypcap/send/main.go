package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"time"
)

var (
	device       string = "lo0"
	snapshot_len int32  = 1024
	promiscuous  bool   = false
	err          error
	timeout      time.Duration = 30 * time.Second
	handler      *pcap.Handle
	buffer       gopacket.SerializeBuffer
	options      gopacket.SerializeOptions
)

func main() {

	devices, _ := pcap.FindAllDevs()
	for _, v := range devices {
		fmt.Println(v.Name)
	}
	fmt.Println()

	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		panic(err)
	}

	rawBytes := []byte{10, 20, 30}
	fmt.Println(rawBytes)
	//err = handle.WritePacketData(rawBytes)
	//if err != nil{
	//panic(err)
	//}

	buffer = gopacket.NewSerializeBuffer()

	ipLayer := &layers.IPv4{
		SrcIP: net.IP{127, 0, 0, 1},
		DstIP: net.IP{127, 0, 0, 1},
	}

	ethernetLayer := &layers.Ethernet{
		SrcMAC: net.HardwareAddr{0xAC, 0xBC, 0x32, 0x9A, 0xBB, 0xF3},
		DstMAC: net.HardwareAddr{0xAC, 0xBC, 0x32, 0x9A, 0xBB, 0xF3},
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(61624),
		DstPort: layers.TCPPort(9998),
	}

	buffer = gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, options, ethernetLayer, ipLayer, tcpLayer, gopacket.Payload(rawBytes))
	outgoingPacket := buffer.Bytes()
	fmt.Println("outgoing:", outgoingPacket)
	err = handle.WritePacketData(outgoingPacket)
	fmt.Println(err)
}
