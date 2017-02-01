package pkt

import (
	"github.com/google/gopacket"
	"log"
)

func UDPPacket(eth *Ethernet, ip *IPv4, udp *UDP) ([]byte, error) {
	options := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	rawBytes := []byte{10, 20, 30}
	if err := udp.udp.SetNetworkLayerForChecksum(ip.ipv4); err != nil {
		return nil, err
	}
	buffer := gopacket.NewSerializeBuffer()
	err := gopacket.SerializeLayers(buffer, options,
		eth.eth,
		ip.ipv4,
		udp.udp,
		gopacket.Payload(rawBytes),
	)
	if err != nil {
		log.Fatal("hoge", err)
		return nil, err
	}
	out := buffer.Bytes()
	return out, nil
}
