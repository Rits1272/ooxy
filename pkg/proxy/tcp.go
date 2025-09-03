package proxy

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
	"fmt"
)

type TCPSession struct {
	localSeq uint32
	localAck uint32

	acked uint32

	socket int

	srcIP net.IP
	dstIP net.IP

	srcPort layers.TCPPort
	dstPort layers.TCPPort

	time int64

	state       string
	hostname    string
	connectSent int

	proxyWriteChan chan []byte // write to upstream proxy
}

var TCPSessions = make(map[string]*TCPSession)

func ProxyTCP(packet gopacket.Packet) {
    tcpLayer := packet.Layer(layers.LayerTypeTCP)
    if tcpLayer == nil {
        fmt.Println("Not a TCP packet")
        return
    }
    tcp := tcpLayer.(*layers.TCP)

    var srcIP, dstIP net.IP
    var ipVersion string

    if ip4Layer := packet.Layer(layers.LayerTypeIPv4); ip4Layer != nil {
        ip4 := ip4Layer.(*layers.IPv4)
        srcIP, dstIP = ip4.SrcIP, ip4.DstIP
        ipVersion = "IPv4"
    } else if ip6Layer := packet.Layer(layers.LayerTypeIPv6); ip6Layer != nil {
        ip6 := ip6Layer.(*layers.IPv6)
        srcIP, dstIP = ip6.SrcIP, ip6.DstIP
        ipVersion = "IPv6"
    } else {
        fmt.Println("Not an IPv4 or IPv6 packet")
        return
    }

    // Unique session key
    key := fmt.Sprintf("%s:%d-%s:%d", srcIP, tcp.SrcPort, dstIP, tcp.DstPort)

    session, exists := TCPSessions[key]
    if !exists {
        fmt.Printf("New TCP session (%s): %s\n", ipVersion, key)

        session = &TCPSession{
            srcIP:          srcIP,
            dstIP:          dstIP,
            srcPort:        tcp.SrcPort,
            dstPort:        tcp.DstPort,
            state:          "NEW",
            proxyWriteChan: make(chan []byte, 100),
        }
        TCPSessions[key] = session
    } else {
        fmt.Printf("Existing TCP session (%s): %s\n", ipVersion, key)
    }

    // Now you can use `session` to handle this packet
}
