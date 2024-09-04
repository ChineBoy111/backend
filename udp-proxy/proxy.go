package main

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

var opts struct {
	Src string   `long:"src" default:":3302" description:"src ip:port to listen on"`
	Dst []string `long:"dst" default:"192.168.220.132:3302" description:"dst ip:port to forward to"`
	Buf int      `long:"buf" default:"64" description:"max buffer size"`
}

func main() {
	log.SetLevel(log.DebugLevel)
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Errorf("fatal error %v", err.Error())
		os.Exit(1)
	}
	srcAddr, err := net.ResolveUDPAddr("udp", opts.Src) // e.g. :3302
	if err != nil {
		log.Errorf("can NOT resolve srcAddr %v", opts.Src)
		return
	}
	var dstAddr []*net.UDPAddr
	for _, v := range opts.Dst {
		addr, err := net.ResolveUDPAddr("udp", v) // e.g. 192.168.220.132:3302
		if err != nil {
			log.Errorf("can NOT resolve dstAddr %v", v)
			return
		}
		dstAddr = append(dstAddr, addr)
	}
	srcConn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		log.Errorf("can NOT listen on srcAddr %v", opts.Src)
		return
	}
	defer srcConn.Close()
	var dstConn []*net.UDPConn
	for _, v := range dstAddr {
		conn, err := net.DialUDP("udp", nil, v)
		if err != nil {
			log.Errorf("can NOT forward to dstAddr %v", v)
			return
		}
		defer conn.Close()
		dstConn = append(dstConn, conn)
	}
	log.Infof("start UDP proxy, srcAddr is %v, dstAddr is %v", opts.Src, opts.Dst)
	for {
		buf := make([]byte, opts.Buf)
		nBytes, addr, err := srcConn.ReadFromUDP(buf)
		if err != nil {
			log.Warn("can NOT receive packet")
			continue
		}
		_, err = srcConn.WriteToUDP([]byte("proxy ==> PROXY respond ==> remote"), addr)
		if err != nil {
			log.Warn("can NOT respond packet")
		}
		log.WithField("addr", addr.String()).WithField("bytes", nBytes).WithField("msg", string(buf[:nBytes])).Info("packet received")
		for _, v := range dstConn {
			_, err := v.Write(buf[:nBytes])
			if err != nil {
				log.Warn("can NOT forward packet")
			}
		}
	}
}
