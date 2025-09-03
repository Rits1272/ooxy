package utils

import (
	"fmt"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket"
)

func CheckProtocol(buffer []byte) (string, gopacket.Packet) {
	var protocol string = ""
	var packet gopacket.Packet

	if len(buffer) < 20 {
		fmt.Println("Invalid IPv4 packet. IPv4 packet should be of min length of 20 bytes")
		return protocol, packet
	}

	packet = gopacket.NewPacket(buffer, layers.LayerTypeEthernet, gopacket.Default)

	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip := ipLayer.(*layers.IPv4)
			return ip.Protocol.String(), packet
	} else if ip6Layer := packet.Layer(layers.LayerTypeIPv6); ip6Layer != nil {
			ip := ip6Layer.(*layers.IPv6)
			return ip.NextHeader.String(), packet
	}

	return protocol, packet
}

