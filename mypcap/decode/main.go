package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"time"
)

func printPacketInfo_Eth(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeEthernet)
	if layer != nil {
		ethPkt, _ := layer.(*layers.Ethernet)
		fmt.Println()
		fmt.Println(ethPkt.SrcMAC, ethPkt.DstMAC, string(ethPkt.Contents))
	}
}

func printPacketInfo_Arp(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeARP)
	if layer != nil {
		arpPkt, _ := layer.(*layers.ARP)
		fmt.Println(arpPkt)
	}
}

func printPacketInfo_Ip(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeIPv4)
	if layer != nil {
		ipPkt, _ := layer.(*layers.IPv4)
		fmt.Println(ipPkt.DstIP)
	}
}

func printPacketInfo_Tcp(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeTCP)
	if layer != nil {
		tcpPkt, _ := layer.(*layers.TCP)
		fmt.Println(tcpPkt.Options)
	}
}

func printPacketInfo_Application(packet gopacket.Packet) {
	layer := packet.ApplicationLayer()
	if layer != nil {
		fmt.Println(string(layer.Payload()))
	}
}

func printLayers(packet gopacket.Packet) {
	fmt.Println(packet.Layers())
}

func fastDecoding(packet gopacket.Packet) {
	var eth layers.Ethernet
	var ip4 layers.IPv4
	var ip6 layers.IPv6
	//var tcp layers.TCP

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &ip6)
	decode := []gopacket.LayerType{}
	err := parser.DecodeLayers(packet.Data(), &decode)
	fmt.Println(err)
	for _, layerTyepe := range decode {
		fmt.Println(layerTyepe)
	}
}

func main() {
	handle, err := pcap.OpenOffline("./mm.pcap")
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	handleSend, err := pcap.OpenLive("lo0", 1024, false, 30*time.Second)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	nCount := 0
	for pk := range packetSource.Packets() {
		//printPacketInfo_Eth(pk)
		//printPacketInfo_Arp(pk)
		//printLayers(pk)
		//printPacketInfo_Tcp(pk)
		//printPacketInfo_Application(pk)
		//fastDecoding(pk)
		err := handleSend.WritePacketData(pk.Data())
		fmt.Println(err)
		nCount++
		//printPacketInfo_Ip(pk)
	}
	fmt.Println("packet size = ", nCount)
}
