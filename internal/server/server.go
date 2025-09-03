package server

import (
	"fmt"
	"net"
	"ooxy/pkg/proxy"
	"ooxy/pkg/utils"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Options struct {
	ListenAddr   string
	UpstreamAddr string
}

type Server struct {
	opts Options
}

func NewServer(opts Options) *Server {
	return &Server{opts: opts}
}

func (s *Server) Run() error {
	sniffPacket()
	ln, err := net.Listen("tcp", s.opts.ListenAddr)

	if err != nil {
		return fmt.Errorf("Failed to listen at address: %s due to error: %w", s.opts.ListenAddr, err)
	}

	defer ln.Close()

	fmt.Printf("Ooxy listening on %s -> %s\n", s.opts.ListenAddr, s.opts.UpstreamAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("Failed to accept connection: %w", err)
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 65535)

	for {
		readBytes, err := conn.Read(buffer)

		if err != nil {
			fmt.Printf("Error reading data: %s\n", err)
			return
		}

		handlePacket(buffer[:readBytes])
	}
}

func sniffPacket() {
	iface := "en0"

	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		fmt.Println(err)
	}
	defer handle.Close()

	// Only capture IPv4/IPv6 packets
	if err := handle.SetBPFFilter("ip or ip6"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Capturing Layer 3 packets on", iface)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		buffer := packet.Data()
		handlePacket(buffer)
	}
}

func handlePacket(buffer []byte) {
	protocol, packet := utils.CheckProtocol(buffer)
	fmt.Println("PACKET", protocol)

	switch protocol {
	case "TCP":
		proxy.ProxyTCP(packet)
	case "UDP":
		proxy.ProxyUDP(packet)
	case "DNS":
		fmt.Println("DNS protocol not supported yet")
	default:
		fmt.Println("unsupported protocol: %s", protocol)
	}
}
