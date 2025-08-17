package utils

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func parsePacket(buffer []byte) gopacket.Packet {
	if len(buffer) < 20 {
		fmt.Println("Invalid IPv4 packet. IPv4 packet should be of min length of 20 bytes")
		return nil
	}

	packet := gopacket.NewPacket(buffer, layers.LayerTypeIPv4, gopacket.Default)

	return packet
}

func CheckProtocol(buffer []byte) {
	packet := parsePacket(buffer)

	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("Source: %s, Destination: %s\n", ip.SrcIP, ip.DstIP)
		fmt.Printf("Protocol: %s (%d)\n", ip.Protocol.String(), ip.Protocol)
	} else {
		fmt.Println("Not an IPv4 packet")
	}
}
