package proxy

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
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

func proxyTCP(buffer []byte) {

}
