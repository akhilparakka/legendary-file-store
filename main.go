package main

import (
	"fmt"

	"github.com/akhilparakka/legendary-file-store/p2p"
)

func OnPeer(p2p.Peer) error {
	fmt.Println("doing some logic with peer outside of tcp transport")
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportopts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%v\n", msg)
		}
	}()

	err := tr.ListenAndAccept()
	if err != nil {
		panic(err)
	}
	select {}
}
