package test

import (
	"testing"

	"bronya.com/proxy/proxy_net"
)

func TestProxy(t *testing.T) {
	server := proxy_net.NewServer("wanproxy")
	server.Serve()
}