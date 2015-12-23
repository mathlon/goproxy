package main

import (
	"github.com/armon/go-socks5"
)

func startSocksV5(listen string) {
	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", listen); err != nil {
		panic(err)
	}
}
