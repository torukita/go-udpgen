package pkt

import (
	"github.com/google/gopacket/layers"
	"net"
)

type IPv4 struct {
	ipv4 *layers.IPv4
}

func NewIPv4(src, dst string) *IPv4 {
	i := &IPv4{
		ipv4: &layers.IPv4{
			SrcIP:    net.ParseIP("0.0.0.0"),
			DstIP:    net.ParseIP("0.0.0.0"),
			TTL:      2,
			Protocol: layers.IPProtocolUDP,
		},
	}

	if net.ParseIP(src) != nil {
		i.ipv4.SrcIP = net.ParseIP(src)
	}
	if net.ParseIP(dst) != nil {
		i.ipv4.DstIP = net.ParseIP(dst)
	}
	return i
}
