package main

import (
	"github.com/akhilparakka/legendary-file-store/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportopts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	err := tr.ListenAndAccept()
	if err != nil {
		panic(err)
	}
	select {}
}
