package main

import "bronya.com/proxy/proxy_net"

func main() {
	server := proxy_net.NewServer("WAN proxy")
	server.Serve()
}
