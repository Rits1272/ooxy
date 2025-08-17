package main

import (
	"flag"
	"fmt"
	"ooxymoron/internal/server"
)

func main() {
	listen := flag.String("listen", "127.0.0.1:8000", "Address to listen on (e.g., 127.0.0.1:8000)")
	upstream := flag.String("upstream", "", "Upstream server address (e.g., 127.0.0.1:8080)")
	
	flag.Parse()

	if *upstream == "" {
		fmt.Println("Please provide valid upstream address")
		return
	}
	
	// start a server to listen for layer 3 packets
	opts := server.Options{
		ListenAddr:   *listen,
		UpstreamAddr: *upstream,
	}

	server.NewServer(opts).Run()
}
