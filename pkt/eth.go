package pkt

import (
	"github.com/google/gopacket/layers"
	"net"
)

type Ethernet struct {
	eth *layers.Ethernet
}

func NewEthernet(src, dst string) *Ethernet {
	s, _ := net.ParseMAC("00:00:00:00:00:00")
	d, _ := net.ParseMAC("00:00:00:00:00:00")
	e := &Ethernet{
		eth: &layers.Ethernet{
			SrcMAC:       s,
			DstMAC:       d,
			EthernetType: layers.EthernetTypeIPv4,
		},
	}

	if srcmac, err := net.ParseMAC(src); err == nil {
		e.eth.SrcMAC = srcmac
	}
	if dstmac, err := net.ParseMAC(dst); err == nil {
		e.eth.DstMAC = dstmac
	}
	return e
}
