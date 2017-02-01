package pkt

import (
	"github.com/google/gopacket/layers"
)

type UDP struct {
	udp *layers.UDP
}

func NewUDP(srcPort, dstPort uint16) *UDP {
	return &UDP{
		udp: &layers.UDP{
			SrcPort: layers.UDPPort(srcPort),
			DstPort: layers.UDPPort(dstPort),
		},
	}
}
