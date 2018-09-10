package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"os"
	"time"
)

func main() {
	fmt.Println(pcap.Version())

	devices, _ := pcap.FindAllDevs()
	for _, v := range devices {
		fmt.Println("=============begin=============")
		fmt.Println(v.Name)
		for _, addr := range v.Addresses {
			fmt.Println(addr.IP, ",", addr.Netmask)
		}
		fmt.Println("=============end==============\n")
	}

	handle, _ := pcap.OpenLive("en0", 65535, false, 30*time.Second)
	defer handle.Close()

	f, _ := os.Create("test.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(1024, layers.LinkTypeEthernet)
	defer f.Close()

	//filter := "tcp and port 80"
	//handle.SetBPFFilter(filter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet.String())
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	}
}
