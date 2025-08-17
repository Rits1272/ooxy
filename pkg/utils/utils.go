package utils

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func createGoPacket(buffer []byte) gopacket.Packet {
	if len(buffer) < 20 {
		fmt.Println("Invalid IPv4 packet. IPv4 packet should be of min length of 20 bytes")
		return nil
	}

	packet := gopacket.NewPacket(buffer, layers.LayerTypeIPv4, gopacket.Default)

	return packet
}

func CheckProtocol(buffer []byte) string{
	packet := createGoPacket(buffer)

	protocol := nil

	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		protocol = ip.Protocol.String()
	} else {
		fmt.Println("Not an IPv4 packet")
	}

	return protocol
}
